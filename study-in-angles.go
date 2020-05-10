package main

import (
  "os"
  "log"
  "image"
  "image/png"
  "image/color"
  "github.com/mcwhittemore/pixicog-go"
)

func main() {
  cog, err := pixicog.PixicogFromVideoFileName(os.Args[1])
  if err != nil {
    panic(1)
  }

  job := pixicog.NewJob(cog.Rotate(90))
  job = job.Process(func(source, state pixicog.Pixicog) pixicog.Pixicog {
    n := len(source)
    width := source.Width()
    height := source.Height()

    img := image.NewRGBA(source.Bounds())
    state = append(state, img)

    colors := make([]color.Color, n)
    for x := 0; x < width; x++ {
      for y := 0; y < height; y++ {
        for i := 0; i < n; i++ {
          colors[i] = source.GetDiminished(i, x, y, 16)
        }
        c := mostCommon(colors)
        img.Set(x, y, c)
      }
    }

    return state
  })

  Save(job.GetState(), os.Args[2])
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

func Save(img image.Image, filename string) {
  f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
