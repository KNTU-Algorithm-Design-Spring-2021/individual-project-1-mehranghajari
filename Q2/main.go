package main

import (
	"github.com/Arafatk/glot"
)

func main() {
	dimensions := 2
	// The dimensions supported by the plot
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)

	pointGroupName := "Simple Lines"
	style := "lines"
	points := [][]float64{{0, 0, 1, 1}, {0, 1, 1, 0}}
	plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle("Example Plot")
	// Optional: Setting the title of the plot
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	// Optional: Setting label for X and Y axis
	plot.SetXrange(-2, 18)
	plot.SetYrange(-2, 18)
	// Optional: Setting axis ranges
	plot.SavePlot("2.png")
}
