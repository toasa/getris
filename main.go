package main

import (
    "math/rand"
    "time"
    "github.com/veandco/go-sdl2/sdl"
)

const (
    WINDOW_NAME = "tetris"
    WINDOW_WIDTH = 800
    WINDOW_HEIGHT = 600
    CELL_LEN = 20
    FIELD_WIDTH = 10
    FIELD_HEIGHT = 20
    MINO_NUM = 7
)

type Color uint32
const (
    colorWindow Color = 0x00000000
    colorVOID Color = 0x00BBBBBB
    colorI Color = 0x007BD7F9
    colorO Color = 0x00F7D320
    colorS Color = 0x0015A81F
    colorZ Color = 0x00D1252B
    colorJ Color = 0x003122B5
    colorL Color = 0x00E56820
    colorT Color = 0x00673CAD
)

type Window struct {
    window *sdl.Window
    field *Field
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
    // To randomize the selection of tetri-mino.
    rand.Seed(time.Now().UnixNano())
}

func (w *Window) update() {
    w.window.UpdateSurface()
}

func getRandomForm() Form {
    return Form(rand.Intn(MINO_NUM))
}

func (w *Window) run() {
	running := true
    for running {
        for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
            switch t := e.(type) {
            case *sdl.QuitEvent:
                println("Quit")
                running = false
                break
            case *sdl.KeyboardEvent:
                if t.GetType() == sdl.KEYDOWN {
                    switch getKeycode(t) {
                    case KEY_LEFT:
                        w.field.attempt(MoveLeft)
                    case KEY_RIGHT:
                        w.field.attempt(MoveRight)
                    case KEY_DOWN:
                        w.field.attempt(MoveDown)
                        if w.field.curMino == nil {
                            w.field.addMino(getRandomForm())
                            w.update()
                        }
                    }
                }
                w.update()
            }
        }
    }
}

func start(w *Window) {
    w.field.addMino(getRandomForm())
    w.update()
    w.run()
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
    start(w)
}
