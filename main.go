package main

import (
	"gr8/sys"
	"gr8/util"
	"image/color"
	"os"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	DISPLAY_WIDTH  int = 1024
	DISPLAY_HEIGHT int = 512
)

func now() int64 {
	return time.Now().UnixMilli()
}

func run() {
	var (
		system  sys.Chip8    = sys.NewChip8()
		texture pixel.Sprite = util.BlankSprite()
	)

	// Load our ROM file
	system.LoadROMFromFile(os.Args[1])

	// Get our timing delay
	delay, _ := strconv.Atoi(os.Args[2])

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "gr8",
		Bounds: pixel.R(0, 0, float64(DISPLAY_WIDTH), float64(DISPLAY_HEIGHT)),
		VSync:  false,
	})

	if err != nil {
		panic(err)
	}

	win.Clear(color.Black)

	var cycleStarted, currentTime, lastRender int64 = now(), now(), now()

	for !win.Closed() {
		cycleStarted = now()
		util.SpriteFromVideo(&system.Video, &texture)

		currentTime = now()

		if currentTime-cycleStarted < int64(delay) {
			time.Sleep(time.Millisecond * time.Duration(int64(delay)-(currentTime-cycleStarted)))
		}

		system.Cycle()

		if currentTime-lastRender < 13 {
			continue
		}

		texture.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))

		win.Update()
		lastRender = now()
	}
}

func main() {
	pixelgl.Run(run)
}
