package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

type Field struct {
    surface *sdl.Surface
    board [FIELD_HEIGHT][FIELD_WIDTH]*Cell
    // TODO: curMinoを実装次第、置き換える
    curCell *Cell
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

// TODO: addMinoに置き換える
func (f *Field) addCell(h, w int, color Color) {
    c := newCell(Falling, color)
    f.curCell = c
    f.board[h][w] = c
    c.draw(h, w, f.surface)
}
