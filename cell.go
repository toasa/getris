package main

type State uint8
const (
    VOID State = iota
    Filled
    Falling
)

type Cell struct {
    state State
    color Color
}

func newCell(state State, color Color) *Cell {
    return &Cell{ state: state, color: color }
}

func (c *Cell) modify(s State, col Color) {
    c.state = s
    c.color = col
}

func (c *Cell) toVoid() {
    c.state = VOID
    c.color = colorVOID
}

