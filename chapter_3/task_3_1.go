// Surface computes an SVG rendering of a 3D surface function.
// Task:
// If the function f returns a non-finite float64 value, the SVG file will contain invalid <polygon> elements
// (although many svg renderers handle this gracefully). Modify the program to skip invalid polygons.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 420            // canvas size
	cells         = 100                 // numnber of grid cells
	xyrange       = 30.0                // axis ranges -xyrange ... + xyrange
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: whithe; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aOk := corner(i+1, j)
			bx, by, bOk := corner(i, j)
			cx, cy, cOk := corner(i, j+1)
			dx, dy, dOk := corner(i+1, j+1)

			// If no error occures, print polygon
			if aOk && bOk && cOk && dOk {
				fmt.Printf("<polygon points='%f,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z, ok := f(x, y)
	if !ok {
		return 0.0, 0.0, false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	surfaceHeight := math.Sin(r) / r
	if math.IsNaN(surfaceHeight) {
		return 0.0, false
	}
	return surfaceHeight, true
}
