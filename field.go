package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

type Field struct {
    surface *sdl.Surface
    board [FIELD_HEIGHT][FIELD_WIDTH]*Cell
    curMino *Mino
}

func newField(window *sdl.Window) (*Field, error) {
    var board [FIELD_HEIGHT][FIELD_WIDTH]*Cell
    for h := 0; h < FIELD_HEIGHT; h++ {
        for w := 0; w < FIELD_WIDTH; w++ {
            board[h][w] = newCell(VOID, colorVOID)
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

func (f *Field) draw() {
    for h := 0; h < FIELD_HEIGHT; h++ {
        for w := 0; w < FIELD_WIDTH; w++ {
            cell := f.board[h][w]
            cell.draw(h, w, f.surface)
        }
    }
}

func (f *Field) setCell(h, w int, c *Cell) {
    f.board[h][w] = c
}

func (f *Field) getCell(h, w int) *Cell {
    return f.board[h][w]
}

func (f *Field) setMino(m *Mino) {
    f.curMino = m
    for _, coord := range m.coords {
        h := coord[0]
        w := coord[1]

        // Cells that extend beyond the top of the field are valid,
        // however do not draw.
        if h < 0 {
            continue
        }
        cell := newCell(Falling, m.color())
        f.setCell(h, w, cell)
    }
}

func (f *Field) addMino(form Form) {
    m := getInitPosMino(form)
    f.setMino(m)
    f.draw()
}
