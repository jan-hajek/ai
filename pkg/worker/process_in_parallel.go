package worker

import (
	"context"
	"runtime"
	"sync"

	"github.com/jan-hajek/ai/pkg/ai/mathx"
)

type config struct {
	workersCount int
}

func WithWorkersCount(workersCount int) Option {
	return func(cfg *config) {
		cfg.workersCount = workersCount
	}
}

func WithWorkersCountLimitedByCpu(workersCount int) Option {
	return func(cfg *config) {
		cfg.workersCount = mathx.Min(workersCount, runtime.NumCPU())
	}
}

type Option = func(cfg *config)

func ProcessInParallel[inputType any, outputType any](
	ctx context.Context,
	inputArray []inputType,
	f func(context.Context, inputType) (outputType, error),
	auxOptions ...Option,
) (output []outputType, _ error) {
	cfg := config{
		workersCount: 3,
	}
	for _, opt := range auxOptions {
		opt(&cfg)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	inputCh := make(chan inputType)
	outputCh := make(chan outputType)

	var outputErr error
	once := sync.Once{}

	wg := sync.WaitGroup{}
	for i := 0; i < cfg.workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case item, ok := <-inputCh:
					if !ok {
						return
					}
					out, err := f(ctx, item)
					if err != nil {
						once.Do(func() {
							outputErr = err
						})
						cancel()
					}
					outputCh <- out
				}
			}
		}()
	}

	go func() {
		// input channel is close when:
		// - all items are processed
		// - error occurred, and context was canceled
		for index := range inputArray {
			select {
			case <-ctx.Done():
				break
			case inputCh <- inputArray[index]:
			}
		}
		close(inputCh)

		// inputCh is closed, so all workers will be stopped
		// wait for all workers to finish
		wg.Wait()

		// when all workers are finished, we can be sure that outputCh is empty
		// and we can close it
		close(outputCh)
	}()

	for item := range outputCh {
		output = append(output, item)
	}

	if outputErr != nil {
		return nil, outputErr
	}

	return output, nil
}
