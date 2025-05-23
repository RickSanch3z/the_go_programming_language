// Server4 is a minimal servert to send random Lissajous figures to the web client

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"

	"log"
	"net/http"
)

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

const (
	backgroundIndex = 0	// first color in palette
)

func lissajous(out io.Writer) {
	const (
		cycles = 5		// number of complete x oscillator revolutions
		res = 0.001		// angular resolution
		size = 100		// image canvas covers [-size...+size]
		nframes = 64	// number of animation frames
		delay = 8		// delay between frames in 10ms utils
	)

	palette_len := uint8(len(palette))
	var pixel_color_palette_index uint8 = 0
	freq := rand.Float64() * 3.0	// relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		pixel_color_palette_index = (uint8(i) % (palette_len - 1) + 1)
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), pixel_color_palette_index)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)	// NOTE: ignoring encoding errors
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler to send Lissajour gif
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}