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

type State uint8
const (
    VOID State = iota
    Filled
    Falling
)

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

type Cell struct {
    state State
    form Form
    rect sdl.Rect
}

type Field struct {
    surface *sdl.Surface
    board [FIELD_HEIGHT][FIELD_WIDTH]*Cell
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

func newField(window *sdl.Window) (*Field, error) {
    var board [FIELD_HEIGHT][FIELD_WIDTH]*Cell
    for h := 0; h < FIELD_HEIGHT; h++ {
        for w := 0; w < FIELD_WIDTH; w++ {
            board[h][w] = newCell(VOID)
        }
    }

	surface, err := window.GetSurface()
    if err != nil {
        return nil, err
    }

    f := &Field{
        board: board,
        surface: surface,
    }

    return f, nil
}

func newCell(state State) *Cell {
    return &Cell{ state: state }
}

func (w *Window) initialize() {
    w.field.draw()
	w.update()
}

func (w *Window) update() {
    w.window.UpdateSurface()
}

func (f *Field) draw() {
    for h := 0; h < FIELD_HEIGHT; h++ {
        for w := 0; w < FIELD_WIDTH; w++ {
            cell := f.board[h][w]
            cell.draw(h, w, f.surface)
        }
    }
}

func (c *Cell) draw(h, w int, surface *sdl.Surface) {
    rect := c.getRect(h, w)
    surface.FillRect(&rect, uint32(c.color()))
}

func (c *Cell) color() Color {
    if c.state == VOID {
        return colorVOID
    }

    switch c.form {
    case I:
        return colorI
    case O:
        return colorO
    case S:
        return colorS
    case Z:
        return colorZ
    case J:
        return colorJ
    case L:
        return colorL
    case T:
        return colorT
    }

    panic("cell color cannot determined")
}

func (c *Cell) getRect(h, w int) sdl.Rect{
    return sdl.Rect {
        int32(w * CELL_LEN),
        int32(h * CELL_LEN),
        CELL_LEN,
        CELL_LEN,
    }
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

    c := newCell(Falling)
    c.form = I
    c.draw(2, 3, w.field.surface)
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
