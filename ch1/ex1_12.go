// Server that takes URL paremeters and draw a Lissajous gif
// Parameters:
// cycles int [default value: 5]
// res float [default value: 0.001]
// size int [default value: 100]
// nframes int [default value: 64]
// delay int [default value: 8]

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"

	"fmt"
	"log"
	"net/http"	// https://pkg.go.dev/net/http#Request , https://pkg.go.dev/net/url#URL
	"net/url"	// https://pkg.go.dev/net/url#ParseQuery
	"strconv"	// https://pkg.go.dev/strconv
)

type lissajous_params struct {
	cycles, size, nframes, delay int
	res float64
}

var palette = []color.Color{
	color.RGBA{0x3F, 0x3F, 0x3F, 0xFF}, 
	color.RGBA{0x33, 0x00, 0x00, 0xFF},
	color.RGBA{0x77, 0x00, 0x00, 0xFF},
	color.RGBA{0xBB, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x33, 0x00, 0xFF},
	color.RGBA{0x00, 0x77, 0x00, 0xFF},
	color.RGBA{0x00, 0xBB, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0x33, 0xFF},
	color.RGBA{0x00, 0x00, 0x77, 0xFF},
	color.RGBA{0x00, 0x00, 0xBB, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.RGBA{0x33, 0x00, 0x00, 0xFF},
	color.RGBA{0x77, 0x33, 0x00, 0xFF},
	color.RGBA{0xBB, 0x77, 0x33, 0xFF},
	color.RGBA{0xFF, 0xBB, 0x77, 0xFF},
	color.RGBA{0x00, 0xFF, 0xBB, 0xFF},
	color.RGBA{0x33, 0x00, 0xFF, 0xFF},
	color.RGBA{0x77, 0x33, 0x00, 0xFF},
	color.RGBA{0xBB, 0x77, 0x33, 0xFF}}

// this constant is not used (here from the first example with Lissajous figures)
const (
	backgroundIndex = 0	// first color in palette
)

func lissajousV2(out io.Writer, params lissajous_params) {
	palette_len := uint8(len(palette))
	var pixel_color_palette_index uint8 = 0
	freq := rand.Float64() * 3.0	// relative frequency of y oscillator
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0

	for i := 0; i < params.nframes; i++ {
		pixel_color_palette_index = (uint8(i) % (palette_len - 1) + 1)
		rect := image.Rect(0, 0, 2 * params.size + 1, 2 * params.size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(params.cycles) * 2 * math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(params.size + int(x * float64(params.size) + 0.5), params.size + int(y * float64(params.size) + 0.5), pixel_color_palette_index)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)	// NOTE: ignoring encoding errors
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler to send Lissajous gif
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous_p, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse the query.\n")
		return
	}

	// set default values for Lissajous figure
	liss_p := lissajous_params{cycles: 0,
								size: 100,
								nframes: 64,
								delay: 8,
								res: 0.001}
	
	if lissajous_p.Get("cycles") != "" {
		if c, err := strconv.ParseUint(lissajous_p.Get("cycles"), 10, 64); err != nil {
			fmt.Fprintf(w, "Failed to convert \"cycles\"\n")
			return
		} else {
			liss_p.cycles = int(c)
		}
	}

	if lissajous_p.Get("size") != "" {
		if s, err := strconv.ParseUint(lissajous_p.Get("size"), 10, 64); err != nil {
			fmt.Fprintf(w, "Failed to convert \"size\"\n")
			return
		} else {
			liss_p.size = int(s)
		}
	}

	if lissajous_p.Get("nframes") != "" {	
		if f, err := strconv.ParseUint(lissajous_p.Get("nframes"), 10, 64); err != nil {
			fmt.Fprintf(w, "Failed to convert \"nframes\"\n")
			return
		} else {
			liss_p.nframes = int(f)
		}
	}

	if lissajous_p.Get("delay") != "" {	
		if d, err := strconv.ParseUint(lissajous_p.Get("delay"), 10, 64); err != nil {
			fmt.Fprintf(w, "Failed to convert \"delay\"\n")
			return
		} else {
			liss_p.delay = int(d)
		}
	}

	if lissajous_p.Get("res") != "" {		
		if r, err := strconv.ParseFloat(lissajous_p.Get("res"), 64); err != nil {
			fmt.Fprintf(w, "Failed to convert \"res\"\n")
			return
		} else {
			liss_p.res = float64(r)
			if r <= 0.0 {
				fmt.Fprint(w, "Resolution must be greter than 0.0\n")
				return
			}
		}
	}

	lissajousV2(w, liss_p)
}