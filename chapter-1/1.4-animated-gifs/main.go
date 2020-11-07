/*
Builds a gif animation frame by frame
calculating oscillations for each image

Note: writes to os.Stdout so to view what is received by the terminal
we use the > output.gif"

Example:
$ go run main.go > output.gif
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 255, 0, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 0, 255, 255},
}

const (
	background = 0 // first color in palette
	color1 = 1 // next color in palette
	color2 = 2
	color3 = 3
)

func main() {
	lissajous(os.Stdout)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 500   // image canvas covers [-size..+size]
		nframes = 128   // number of animation frames
		delay   = 5     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			uint8(rand.Intn(len(palette) - 1) + 1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
