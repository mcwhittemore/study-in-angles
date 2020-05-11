package main

import (
  "image"
  "os"
  "image/color"
  "github.com/mcwhittemore/pixicog-go"
)

func main() {
  cog, err := pixicog.ImageListFromVideoFileName(os.Args[1])
  if err != nil {
    panic(1)
  }

  job := pixicog.NewJob(cog.Rotate(90))
  job = job.Process(func(source, state pixicog.ImageList) pixicog.ImageList {
    width := source.Width()
    height := source.Height()

    img := image.NewRGBA(source.Bounds())
    state = append(state, img)

    for x := 0; x < width; x++ {
      for y := 0; y < height; y++ {
        colors := source.GetDiminished(x,y,16)
        c := mostCommon(colors)
        img.Set(x, y, c)
      }
    }

    return state
  })

  job.GetState().SavePNG(os.Args[2])
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

