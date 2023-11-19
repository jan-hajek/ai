package imagedataextractor

import (
	"context"
	"fmt"
	"image"
	"math"

	"github.com/jan-hajek/ai/pkg/ai/imagex"
)

type ImageShapeExtractor struct {}

type StrikePoint struct {
	x int
	y int
	isStrike int
	visited bool
}

func NewImageShapeExtractor(
) *ImageShapeExtractor {
	return &ImageShapeExtractor{}
}

func (i *ImageShapeExtractor) ExtractShapes(_ context.Context, img image.Image) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf(string(width))

	strikes := make([][]StrikePoint, width)
	for i := range strikes {
		strikes[i] = make([]StrikePoint, height)
	}

	// recognize strikes
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := imagex.RgbaToPixel(img.At(x, y).RGBA())
			strikes[x][y].isStrike = isStrike(pixel)
			strikes[x][y].x = x
			strikes[x][y].y = y
			fmt.Print(strikes[x][y].isStrike)
		}
		fmt.Print("\n")
	}

	shapes := make([][]int, width)
	for i := range shapes {
		shapes[i] = make([]int, height)
	}

	for y := 0; y < height; y++ {
		fmt.Printf("%d-", y)
		for x := 0; x < width; x++ {
			point := strikes[x][y]
			if point.isStrike == 1 {
				res := analyzePoint(strikes, point, width, height)
				shapes[x][y] = res
			} else {
				shapes[x][y] = 0
			}
			fmt.Print(shapes[x][y])
		}
		fmt.Print("\n")
	}
}

func isStrike(pixel imagex.Pixel) int {
	if float64(pixel.R + pixel.G + pixel.B)/float64(255*3) > 0.4 {
		return 0
	}
	return 1
}

func analyzePoint(strikes [][]StrikePoint, point StrikePoint, width int, height int) int {
	boxSize := 3
	xx := point.x
	yy := point.y

	xLeft := int(math.Max(float64(xx-boxSize), 0))
	xRight := int(math.Min(float64(xx+boxSize), float64(width-1)))
	yTop := int(math.Max(float64(yy-boxSize), 0))
	yBottom := int(math.Min(float64(yy+boxSize), float64(height-1)))

	//fmt.Printf(" Box is %d %d %d %d ", xLeft, xRight, yTop, yBottom)

	fx := make([]float64, 0)
	fy := make([]float64, 0)
	fxdL := make([]float64, 0)
	fydL := make([]float64, 0)
	fxdR := make([]float64, 0)
	fydR := make([]float64, 0)

	edgeCrossesCount := 0
	crossTop := 0
	crossBottom := 0
	crossLeft := 0
	crossRight := 0

	for i := 0; i <= xRight-xLeft; i++ {
		x := xLeft + i
		colPoints := getXPointsGroup(strikes, x, yBottom, yTop)
		if len(colPoints) > 0 {
			_, ay := computeAverageCoords(colPoints)
			//fmt.Printf("Average in x = %d is %.2f and has %d points \n", x, ay, len(colPoints))
			fx = append(fx, ay)
			if x <= xx {
				if len(fx) > 1 {
					fxdL = append(fxdL, ay-fx[len(fx)-2])
				}
			}
			if x > xx {
				if len(fx) > 1 {
					fxdR = append(fxdR, ay-fx[len(fx)-2])
				}
			}
		}
		// find edge crosses
		if strikes[x][yTop].isStrike == 1 && crossTop == 0 {
			crossTop = 1
			edgeCrossesCount++
		} else if strikes[x][yTop].isStrike == 0 && crossTop == 1 {
			crossTop = 0
		}
		if strikes[x][yBottom].isStrike == 1 && crossBottom == 0 {
			crossBottom = 1
			edgeCrossesCount++
		} else if strikes[x][yTop].isStrike == 0 && crossBottom == 1 {
			crossBottom = 0
		}
	}

	for i := 0; i <= yBottom-yTop; i++ {
		y := yTop + i
		colPoints := getYPointsGroup(strikes, y, xLeft, xRight)
		if len(colPoints) > 0 {
			ax, _ := computeAverageCoords(colPoints)
			//fmt.Printf("Average in y = %d is %.2f \n", y, ax)
			fy = append(fy, ax)
			if y <= yy {
				if len(fy) > 1 {
					fydL = append(fydL, ax-fy[len(fy)-2])
				}
			}
			if y > yy {
				if len(fy) > 1 {
					fydR = append(fydR, ax-fy[len(fy)-2])
				}
			}
		}
		// find edge crosses
		if strikes[xLeft][y].isStrike == 1 && crossLeft == 0 {
			crossLeft = 1
			edgeCrossesCount++
		} else if strikes[xLeft][y].isStrike == 0 && crossLeft == 1 {
			crossLeft = 0
		}
		if strikes[xRight][y].isStrike == 1 && crossRight == 0 {
			crossRight = 1
			edgeCrossesCount++
		} else if strikes[xRight][y].isStrike == 0 && crossRight == 1 {
			crossRight = 0
		}
	}
	if strikes[xRight][yTop].isStrike == 1 {
		edgeCrossesCount--
	}
	if strikes[xRight][yBottom].isStrike == 1 {
		edgeCrossesCount--
	}
	if strikes[xLeft][yTop].isStrike == 1 {
		edgeCrossesCount--
	}
	if strikes[xLeft][yBottom].isStrike == 1 {
		edgeCrossesCount--
	}

	result := 0

	//fmt.Printf("Len(fx): %d \n", len(fx))
	//fmt.Printf("Len(fy): %d \n", len(fy))
	if edgeCrossesCount == 4 {
		result = 4
	} else if len(fx) >= len(fy) {
		//direction in x axis
		if len(fxdL) > 0 && len(fxdR) > 0 {
			fxdLAvg := avg(fxdL)
			fxdRAvg := avg(fxdR)
			//fmt.Printf("fxdLAvg: %.2f \n", fxdLAvg)
			//fmt.Printf("fxdRAvg: %.2f \n", fxdRAvg)
			if math.Abs(fxdLAvg - fxdRAvg) < 0.2 {
				//fmt.Print("it is a line in x direction \n")
				result = 1
			} else if (fxdLAvg < 0 && fxdRAvg >= 0) || (fxdLAvg >= 0 && fxdRAvg < 0) {
				//fmt.Print("it is a spike in x direction \n")
				result = 3
			} else {
				//fmt.Printf("it is a curve, dif is %.2f \n", fxdLAvg - fxdRAvg)
				if point.y == 20 {
					fmt.Printf("Point: x=%d, y=%d, ", point.x, point.y)
					fmt.Printf("fxdl: %.2f, fxdr: %.2f, fxdLAvg: %.2f, fxdRAvg: %.2f \n", fxdL, fxdR, fxdLAvg, fxdRAvg)
					fmt.Printf("x: %d y: %d = %d \n", xx, yTop, strikes[xx][yTop].isStrike)
					fmt.Printf("x: %d y: %d = %d \n", xx, yBottom, strikes[xx][yBottom].isStrike)
					fmt.Printf("x: %d y: %d = %d \n", xLeft, yy, strikes[xLeft][yy].isStrike)
					fmt.Printf("x: %d y: %d = %d \n", xRight, yy, strikes[xRight][yy].isStrike)

					fmt.Printf("fx: %d fy: %d \n", len(fx), len(fy))
					fmt.Printf("Edge crosses count: %d \n", edgeCrossesCount)
				}
				result = 2
			}
		} else  {
			//fmt.Printf("Point: x=%d, y=%d", point.x, point.y)
			//fmt.Printf("fxdl: %.2f, fxdr: %.2f \n", fxdL, fxdR)
			result = 5
		}
	} else if len(fy) > len(fx) {
		//direction in y axis
		if len(fydL) > 0 && len(fydR) > 0 {
			fydLAvg := avg(fydL)
			fydRAvg := avg(fydR)
			if math.Abs(fydLAvg - fydRAvg) < 0.2 {
				//fmt.Print("it is a line in y direction \n")
				result = 1
			} else if (fydLAvg < 0 && fydRAvg >= 0) || (fydLAvg >= 0 && fydRAvg < 0) {
				//fmt.Print("it is a spike in y direction \n")
				result = 3
			} else {
				//fmt.Printf("it is a curve, dif is %.2f \n", fydLAvg - fydRAvg)
				result = 2
			}
		} else  {
			//fmt.Printf("Point: x=%d, y=%d", point.x, point.y)
			//fmt.Printf("fydl: %.2f, fydr: %.2f, len fydl: %d, len fydR: %d \n", fydL, fydR, len(fydL), len(fydR))
			result = 5
		}
	}
	// 1-line
	// 2-curve
	// 3-spike
	// 4-cross
	// 5-edge

	return result
}

func computeAverageCoords(points []StrikePoint) (float64, float64) {
	sumX := 0
	sumY := 0
	for _, point := range points {
		sumX = sumX + point.x
		sumY = sumY + point.y
	}

	if len(points) > 0 {
		return float64(sumX)/float64(len(points)), float64(sumY)/float64(len(points))
	}
	return 0, 0
}

func getXPointsGroup(strikes [][]StrikePoint, xx int, yBottom int, yTop int) []StrikePoint {
	colPoints := make([]StrikePoint,0)

	for i := 0; i <= yBottom-yTop; i++ {
		y := yTop + i
		p := strikes[xx][y]
		if p.isStrike == 1 {
			colPoints = append(colPoints, p)
		}
	}

	return colPoints
}

func getYPointsGroup(strikes [][]StrikePoint, yy int, xLeft int, xRight int) []StrikePoint {
	rowPoints := make([]StrikePoint,0)

	for i := 0; i <= xRight-xLeft; i++ {
		x := xLeft + i
		p := strikes[x][yy]
		if p.isStrike == 1 {
			rowPoints = append(rowPoints, p)
		}
	}

	return rowPoints
}

func avg(arr []float64) float64 {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum = sum+arr[i]
	}
	return sum/float64(len(arr))
}

