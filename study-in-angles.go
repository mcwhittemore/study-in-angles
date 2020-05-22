package main

import (
	"encoding/base64"
	"fmt"
	"github.com/mcwhittemore/pixicog"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func (r *Runner) ProcessLoad() {
	src, err := pixicog.ImageListFromVideoFileName(os.Args[1])
	if err != nil {
		panic(err)
	}
	src = src.Rotate(90)
	r.state["src"] = src
}

func (r *Runner) ProcessFirstLayer() {
	src := r.state["src"]

	width := src.Width()
	height := src.Height()

	img := image.NewRGBA(src.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			colors := src.GetDiminished(x, y, 16)
			c := mostCommon(colors)
			img.Set(x, y, c)
		}
	}

	state := pixicog.ImageList{}
	state = append(state, img)

	r.state["src"] = state
}

func (r *Runner) ProcessSave() {
	r.state["src"].SavePNG(os.Args[2])
}

func mostCommon(colors []color.Color) color.Color {
	common := colors[0]
	commonCount := 0

	n := len(colors)

	for i := 0; i < n; i++ {
		t := colors[i]
		tc := 0
		for j := i + 1; j < n; j++ {
			if t == colors[j] {
				tc++
			}
		}
		if tc > commonCount {
			commonCount = tc
			common = t
		}
	}

	return common
}
