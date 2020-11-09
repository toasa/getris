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
    color Color
    // Need?
    rect sdl.Rect
}

func newCell(state State, color Color) *Cell {
    return &Cell{ state: state, color: color }
}

func (c *Cell) draw(h, w int, surface *sdl.Surface) {
    rect := c.getRect(h, w)
    surface.FillRect(&rect, uint32(c.color))
}

func (c *Cell) getRect(h, w int) sdl.Rect{
    return sdl.Rect {
        int32(w * CELL_LEN),
        int32(h * CELL_LEN),
        CELL_LEN,
        CELL_LEN,
    }
}

func (c *Cell) toVoid() {
    c.state = VOID
    c.color = colorVOID
}

