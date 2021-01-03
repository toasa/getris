package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/veandco/go-sdl2/sdl"
)

const (
    WINDOW_NAME = "getris"
    MAGNIFICATION = 1.5
    WINDOW_WIDTH = 200 * MAGNIFICATION
    WINDOW_HEIGHT = 400 * MAGNIFICATION
    CELL_LEN = 20 * MAGNIFICATION
    FIELD_WIDTH = 10
    FIELD_HEIGHT = 20
    MINO_NUM = 7
    DROP_INTERVAL_MS = 1500
)

type Color uint32
const (
    colorWindow Color = 0x00000000
    colorVOID Color = 0x00DDDDDD
    colorDropPred Color = 0x00EEEEEE
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
    // To implement automatically descent of tetri-mino
    // in every unit time, we use ticker.
    ticker := time.NewTicker(time.Millisecond * DROP_INTERVAL_MS)
    defer ticker.Stop()

    running := true
    for running {
        select {
        case <-ticker.C:
            w.field.attempt(MoveDrop)
            w.field.attemptDescent()

            if w.field.curMino == nil {
                m := getInitPosMino(getRandomForm())
                if w.field.isGameOver(m) {
                    fmt.Println("Game Over!")
                    running = false
                }
                w.field.addMino(m)
            }

            w.update()
        default:
            for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
                switch t := e.(type) {
                case *sdl.QuitEvent:
                    fmt.Println("Quit")
                    running = false
                    break
                case *sdl.KeyboardEvent:
                    if t.GetType() == sdl.KEYDOWN {
                        switch getKeycode(t) {
                        case KEY_LEFT:
                            w.field.attempt(MoveLeft)
                        case KEY_RIGHT:
                            w.field.attempt(MoveRight)
                        case KEY_DROP:
                            w.field.attempt(MoveDrop)
                        case KEY_HARD_DROP:
                            w.field.attempt(MoveHardDrop)
                        case KEY_ROT_LEFT:
                            w.field.attempt(RotLeft)
                        case KEY_ROT_RIGHT:
                            w.field.attempt(RotRight)
                        }
                    }

                    w.field.attemptDescent()

                    if w.field.curMino == nil {
                        m := getInitPosMino(getRandomForm())
                        if w.field.isGameOver(m) {
                            fmt.Println("Game Over!")
                            running = false
                        }
                        w.field.addMino(m)
                    }

                    w.update()
                }
            }
        }
    }
}

func start(w *Window) {
    newm := getInitPosMino(getRandomForm())
    w.field.addMino(newm)
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
