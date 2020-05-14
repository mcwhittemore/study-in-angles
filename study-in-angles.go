package main

import (
  "image"
  "os"
  "image/color"
  "github.com/mcwhittemore/pixicog"
)

func ProcessLoad(src, state pixicog.ImageList) (pixicog.ImageList, pixicog.ImageList) {
  src, err := pixicog.ImageListFromVideoFileName(os.Args[1])
  if err != nil {
    panic(err)
  }

  state = pixicog.ImageList{}
  src = src.Rotate(90)

  return src, state
}

func ProcessFirstLayer(src, state pixicog.ImageList) (pixicog.ImageList, pixicog.ImageList) {
  width := src.Width()
  height := src.Height()

  img := image.NewRGBA(src.Bounds())

  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      colors := src.GetDiminished(x,y,16)
      c := mostCommon(colors)
      img.Set(x, y, c)
    }
  }

  state = append(state, img)
  return src, state
}

func ProcessSave(src, state pixicog.ImageList) (pixicog.ImageList, pixicog.ImageList) {
  state.SavePNG(os.Args[2])
  return src, state
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

