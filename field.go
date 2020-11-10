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

// TODO: Judge game over
func (f *Field) addMino(form Form) {
    m := getInitPosMino(form)
    f.setMino(m)
    f.draw()
}

// attempt attempts to a specific move for current tetri-mino.
func (f *Field) attempt(move Move) {
    new_m := f.curMino.move(move)

    // current tetri-mino reaches to bottom or
    // already filled cells.
    if move == MoveDown && f.atBottom(*new_m) {
        f.toFix(*(f.curMino))
        f.draw()
        f.curMino = nil
        return
    }

    if !f.legalMove(*new_m) {
        return
    }
    f.blank(*(f.curMino))
    f.setMino(new_m)
    f.draw()
}

func (f *Field) legalMove(m Mino) bool {
    coords := m.coords
    for _, coord := range coords {
        if coord.isExceedTop() {
            continue
        }
        h := coord.getHeight()
        w := coord.getWidth()

        // Exceed field
        if h >= FIELD_HEIGHT || w < 0 || w >= FIELD_WIDTH {
            return false
        }

        cell := f.getCell(h, w)
        if cell.state == Filled {
            return false
        }
    }
    return true
}

func (f *Field) blank(m Mino) {
    coords := m.coords
    for _, coord := range coords {
        if coord.isExceedTop() {
            continue
        }

        h := coord.getHeight()
        w := coord.getWidth()
        cell := f.getCell(h, w)
        cell.toVoid()
    }
}

func (f *Field) atBottom(m Mino) bool {
    coords := m.coords
    for _, coord := range coords {
        if coord.isExceedTop() {
            continue
        }

        h := coord.getHeight()
        w := coord.getWidth()
        if h >= FIELD_HEIGHT {
            return true
        }

        cell := f.getCell(h, w)
        if cell.state == Filled {
            return true
        }
    }

    return false
}

func (f *Field) toFix(m Mino) {
    coords := m.coords
    for _, coord := range coords {
        if coord.isExceedTop() {
            continue
        }

        h := coord.getHeight()
        w := coord.getWidth()
        cell := f.getCell(h, w)
        cell.state = Filled
    }
}
