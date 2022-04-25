// Surface computes an SVG rendering of a 3D surface function.
// Task:
// Experiment with visualizations of other functions from the math package. Can you produce
// an egg boxc, moguls or a saddle?
package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type mathFunction func(x, y float64) (float64, bool)
type Writer interface {
	Write(p []byte) (n int, err error)
}

const (
	width, height = 600, 420            // canvas size
	cells         = 100                 // numnber of grid cells
	xyrange       = 30.0                // axis ranges -xyrange ... + xyrange
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y, y axes (=30º)
	usage         = "Usage: main.go filename org|saddle|eggbox"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	// Make sure that the target file exists
	f, err := os.OpenFile(strings.Replace(os.Args[1], ".svg", "", -1)+".svg", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	svg(f, org)
}

func svg(w io.Writer, targetMathFunction mathFunction) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: whithe; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aOk := corner(i+1, j, targetMathFunction)
			bx, by, bOk := corner(i, j, targetMathFunction)
			cx, cy, cOk := corner(i, j+1, targetMathFunction)
			dx, dy, dOk := corner(i+1, j+1, targetMathFunction)

			// If no error occures, print polygon
			if aOk && bOk && cOk && dOk {
				fmt.Fprintf(w, "<polygon points='%f,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, targetMathFunction mathFunction) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z, ok := targetMathFunction(x, y)
	if !ok {
		return 0.0, 0.0, false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func org(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	surfaceHeight := math.Sin(r) / r
	if math.IsNaN(surfaceHeight) {
		return 0.0, false
	}
	return surfaceHeight, true

}
