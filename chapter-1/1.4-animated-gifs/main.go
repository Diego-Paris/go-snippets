/*
Builds a gif animation frame by frame
calculating oscillations for each image
then adding the image to the gif struct

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
	color.RGBA{0, 255, 0, 255}, // Green
	color.RGBA{255, 0, 0, 255}, // Red
	color.RGBA{0, 0, 255, 255}, // Blue
}

// Used to identify each color in the palette
const (
	background = 0 // first color in palette
	color1     = 1 // next color in palette
	color2     = 2
	color3     = 3
)

func main() {
	// Writes to terminal
	//? You can save a generated image by streaming output
	//? of your program to a file e.g. app > image.gif
	lissajous(os.Stdout)
}

func init() {
	// Assures random values each run
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

	anim := gif.GIF{LoopCount: nframes} // gif struct, used to build animation

	phase := 0.0 // phase difference

	// iterate through each frame
	for i := 0; i < nframes; i++ {

		// create a rectangle, e.g. if size=100 then
		// dimensions are 201x201
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		//Paletted image with the given width, height of the rect and palette
		img := image.NewPaletted(rect, palette)

		// randomly chooses a color from the palette for frame
		frameColor := uint8(rand.Intn(len(palette)-1) + 1)

		//Calculate each pixel of the oscillation in this current frame
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				frameColor)
		}

		// adds 0.1 to phase
		phase += 0.1

		// append delay and image to the animation for this frame
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// encode gif to writer passing the reference to the struct
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
