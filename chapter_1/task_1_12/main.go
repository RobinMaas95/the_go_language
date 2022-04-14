// .server creates a webserver that generates a lissajous figure and returns it. The following
// parameters can be passed as http parameter to modify the figure: backgroundIndex, lineIndex,
// cycles, res, size, nframes and delay.
//
// result.gif is an example result with the following query url:
// http://localhost:8000/?cycles=10&lineIndex=1&size=200&nframes=120&delay=1

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 0, 255}}

var lissajousParameter = map[string]float64{
	"backgroundIndex": 0,
	"lineIndex":       1,
	"cycles":          5,
	"res":             0.001,
	"size":            100,
	"nframes":         64,
	"delay":           8,
}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path compnent of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	p := getParameter(r.URL.Query())
	lissajous(w, p)

}

// getParameter loops over all parameter lissajous needs and checks for each parameter, if it was
// passed in the URL. It returns a map with all necessary parameters, where the values are the
// passed value if a parameter is part of the request or the default value if not.
func getParameter(p url.Values) map[string]float64 {
	parameter := make(map[string]float64)

	for k, v := range lissajousParameter {
		if val, ok := p[k]; ok {
			f, err := strconv.ParseFloat(val[0], 4)
			if err != nil {
				panic(err)
			}
			parameter[k] = f
		} else {
			parameter[k] = v
		}
	}
	return parameter
}

func lissajous(out io.Writer, p map[string]float64) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: int(p["nframes"])}
	phase := 0.0

	for i := 0; i < int(p["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*int(p["size"])+1, 2*int(p["size"])+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < (p["cycles"])*2*math.Pi; t += p["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(p["size"])+int(x*p["size"]+0.5), int(p["size"])+int(y*p["size"]+0.5), uint8(p["lineIndex"]))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(p["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
