package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

const (
    WINDOW_NAME = "tetris"
    WINDOW_WIDTH = 800
    WINDOW_HEIGHT = 600
    CELL_LEN = 20
    FIELD_WIDTH = 10
    FIELD_HEIGHT = 20
)

type Color uint32
const (
    colorWindow Color = 0x00000000
    colorVOID Color = 0x00BBBBBB
	colorI Color = 0x00E5282E
	colorO Color = 0x00274696
	colorS Color = 0x00EF7E18
	colorZ Color = 0x002CB099
	colorJ Color = 0x00F8D517
	colorL Color = 0x00DF2384
	colorT Color = 0x005CAD2C
)

type Window struct {
    window *sdl.Window
    field *Field
}

type Form uint8
const (
	I Form = iota
	O
	S
	Z
	J
	L
	T
)

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

    field, err := newField(window)
    if err != nil {
        return nil, err
    }

    w := &Window{
        window: window,
        field: field,
    }

    return w, nil
}

func (w *Window) initialize() {
    w.field.draw()
	w.update()
}

func (w *Window) update() {
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

    w.field.addCell(2, 3, colorI)
    w.update()

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
