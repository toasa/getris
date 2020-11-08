package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

const (
    WINDOW_NAME = "tetris"
    WINDOW_WIDTH = 800
    WINDOW_HEIGHT = 600
    CELL_LEN = 20
    FIELD_WIDTH = CELL_LEN * 10
    FIELD_HEIGHT = CELL_LEN * 20
)

const (
    colorField uint32 = 0x00BBBBBB
	colorI uint32 = 0x00E5282E
	colorO uint32 = 0x00274696
	colorS uint32 = 0x00EF7E18
	colorZ uint32 = 0x002CB099
	colorJ uint32 = 0x00F8D517
	colorL uint32 = 0x00DF2384
	colorT uint32 = 0x005CAD2C
)

type Window struct {
    window *sdl.Window
    surface *sdl.Surface
}

func newWindow() (*Window, error) {
	window, err := sdl.CreateWindow(
        WINDOW_NAME,
        sdl.WINDOWPOS_UNDEFINED,
        sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
        WINDOW_HEIGHT,
        sdl.WINDOW_SHOWN,
    )
    if err != nil {
        return nil, err
    }

	surface, err := window.GetSurface()
    if err != nil {
        return nil, err
    }

    w := &Window{
        window: window,
        surface: surface,
    }

    return w, nil
}

func (w *Window) initialize() {
	field := sdl.Rect{0, 0, FIELD_WIDTH, FIELD_HEIGHT}
	w.surface.FillRect(&field, colorField)
	w.window.UpdateSurface()
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

    w, err := newWindow()
    if err != nil {
        panic(err)
    }

    w.initialize()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
