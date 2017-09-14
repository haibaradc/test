package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.White, color.RGBA{0, 255, 0, 1}}

const (
	whiteIndex = 0
	blackIndex = 1
	cumIndex   = 2
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ { //64帧
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette) //调色板
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+5), size+int(y*size+5), cumIndex) //画线
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) //延迟
		anim.Image = append(anim.Image, img)   //一帧图片
	}
	gif.EncodeAll(out, &anim) //输出
}

func main() {
	lissajous(os.Stdout)
}