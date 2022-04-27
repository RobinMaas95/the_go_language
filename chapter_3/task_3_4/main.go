// Surface computes an SVG rendering of a 3D surface function.
// Task:
// Follwoing the approach of the Lissajous example in Section 1.7, construct approach
// a webserver taht computes surfaces and writes SVG data to the client. The server
// must set the content-type header like this ('w.Header().Set("Content-Type", "image/svg+xml")')
// Allow the client to specify values like height width and color as HTTP request parameters.

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	cells   = 100         // numnber of grid cells
	xyrange = 30.0        // axis ranges -xyrange ... + xyrange
	angle   = math.Pi / 6 // angle of x, y, y axes (=30º)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var plotParameter = map[string]string{
	"width":  "600",
	"height": "320",
	"type":   "org",
	"color":  "grey",
}

type mathFunction func(x, y float64) (float64, bool)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	p := getParameter(r.URL.Query())
	svg(w, p)
}

// getParameter loops over all parameter the svg needs and checks for each parameter, if it was
// passed in the URL. It returns a map with all necessary parameters, where the values are the
// passed value if a parameter is part of the request or the default value if not.
func getParameter(p url.Values) map[string]string {
	parameter := make(map[string]string)

	for k, v := range plotParameter {
		if val, ok := p[k]; ok {
			f := val[0]
			parameter[k] = f
		} else {
			parameter[k] = v
		}
	}
	return parameter
}

func svg(w io.Writer, parameter map[string]string) {
	var mfunc mathFunction
	switch parameter["type"] {
	case "eggbox":
		mfunc = eggbox
	case "saddle":
		mfunc = saddle
	default:
		mfunc = org
	}

	width, _ := strconv.ParseFloat(parameter["width"], 4)
	height, _ := strconv.ParseFloat(parameter["height"], 4)

	zmin, zmax := minmax(mfunc)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: grey; stroke-width: 0.7' "+"width='%d' height='%d'>", 320, 600)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aOk := corner(i+1, j, mfunc, width, height)
			bx, by, bOk := corner(i, j, mfunc, width, height)
			cx, cy, cOk := corner(i, j+1, mfunc, width, height)
			dx, dy, dOk := corner(i+1, j+1, mfunc, width, height)

			// If no error occures, print polygon
			if aOk && bOk && cOk && dOk {
				fmt.Fprintf(w, "<polygon points='%f,%g %g,%g %g,%g %g,%g' style='stroke:%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color(i, j, zmin, zmax, mfunc))

			}
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, targetMathFunction mathFunction, width, height float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z, ok := targetMathFunction(x, y)
	if !ok {
		return 0.0, 0.0, false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx, sy)

	xyscale := width / 2 / xyrange // pixels per x or y unit
	zscale := height * 0.4         // pixels per z unit
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func minmax(f mathFunction) (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z, _ := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64, f mathFunction) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z, _ := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}
func org(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	surfaceHeight := math.Sin(r) / r

	if math.IsNaN(surfaceHeight) {
		return 0.0, false
	}
	return surfaceHeight, true

}

func eggbox(x, y float64) (float64, bool) {
	res := -0.1 * (-math.Cos(x) + -math.Cos(y))

	if math.IsNaN(res) {
		return 0.0, false
	}
	return res, true
}

func saddle(x, y float64) (float64, bool) {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	res := (y*y/a2 - x*x/b2)

	if math.IsNaN(res) {
		return 0.0, false
	}
	return res, true
}
