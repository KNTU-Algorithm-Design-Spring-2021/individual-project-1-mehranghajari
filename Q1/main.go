package main

import (
	"fmt"
	"log"
	"math"
)

type Point struct {
	x float64
	y float64
}
type XYMinMax struct {
	minX float64
	maxX float64
	minY float64
	maxY float64
}

func (mm *XYMinMax) String() {
	fmt.Printf("Min X: %f\nMin Y: %f\nMax X: %f\nMax Y: %f\n", mm.minX, mm.minY, mm.maxX, mm.maxY)
}

func calculateMinMax(points []Point, low int, high int) (minMax XYMinMax) {
	var minMaxL, minMaxR XYMinMax
	// If array contains one element
	if low == high {
		minMax.maxX = points[high].x
		minMax.minX = points[high].x
		minMax.maxY = points[high].y
		minMax.minY = points[high].y
		return minMax

	}
	// if array contains two element.
	if high == low+1 {
		minMax.minX, minMax.maxX = compare(points[low].x, points[high].x)
		minMax.minY, minMax.maxY = compare(points[low].y, points[high].y)
		return minMax

	}
	mid := (high + low) / 2
	// Divide arrays in two left and right subsets.
	minMaxL = calculateMinMax(points, low, mid)
	minMaxR = calculateMinMax(points, mid+1, high)

	// Compare to find out max and min of x and y in left and right subsets.
	minMax.minX = math.Min(minMaxL.minX, minMaxR.minX)
	minMax.minY = math.Min(minMaxL.minY, minMaxR.minY)
	minMax.maxX = math.Max(minMaxL.maxX, minMaxR.maxX)
	minMax.maxY = math.Max(minMaxL.maxY, minMaxR.maxY)

	return minMax

}

func compare(x, y float64) (min, max float64) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func main() {
	var numberOfPoints int
	fmt.Print("Please Enter the Number of Points: ")
	_, err := fmt.Scanf("%d", &numberOfPoints)
	if err != nil {
		log.Fatalln("Bad Input...")
	}
	points := make([]Point, numberOfPoints)
	var tempX, tempY float64
	for i := 0; i < numberOfPoints; i++ {
		fmt.Printf("Point %d: ", i+1)
		_, err = fmt.Scanf("%f %f", &tempX, &tempY)
		if err != nil {
			log.Fatalln("Fatal error...")
		}
		points[i] = Point{x: tempX, y: tempY}
	}
	result := calculateMinMax(points, 0, len(points)-1)
	result.String()
}
