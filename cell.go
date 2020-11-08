package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

type State uint8
const (
    VOID State = iota
    Filled
    Falling
)

type Cell struct {
    state State
    form Form
    rect sdl.Rect
}

func newCell(state State) *Cell {
    return &Cell{ state: state }
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

