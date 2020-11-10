package main

import (
    "sort"
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

func (f *Field) attemptDescent() {
    lines := f.getCompleteHorizontalLines()
    if len(lines) > 0 {
        f.eraseLines(lines)
        f.dropRemains(lines)
    }
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

func (f *Field) getCompleteHorizontalLines() []int {
    lines := []int{}
    for h := 0; h < FIELD_HEIGHT; h++ {
        all_filled := true
        for w := 0; w < FIELD_WIDTH; w++ {
            if f.getCell(h, w).state != Filled {
                all_filled = false
            }
        }
        if all_filled {
            lines = append(lines, h)
        }
    }
    return lines
}

func (f *Field) eraseLine(h int) {
    for w := 0; w < FIELD_WIDTH; w++ {
        cell := f.getCell(h, w)
        cell.toVoid()
    }
}

func (f *Field) eraseLines(lines []int) {
    for _, l := range lines {
        f.eraseLine(l)
    }
}

func (f *Field) dropRemains(erasedLines []int) {
    f.eraseLines(erasedLines)

    if len(erasedLines) == 0 {
        return
    }

    erasedMap := func(eLines []int) []bool {
        eMap := make([]bool, FIELD_HEIGHT, FIELD_HEIGHT)
        for _, h := range eLines {
            eMap[h] = true
        }
        return eMap
    }(erasedLines)

    nFixedLine := func(eMap []bool) int {
        n := 0
        for h := FIELD_HEIGHT-1; h >= 0; h-- {
            if eMap[h] {
                break
            }
            n++
        }
        return n
    }(erasedMap)

    // Do nothing.
    if nFixedLine >= FIELD_HEIGHT {
        return
    }

    copyLine := func(dstL, srcL int) {
        for w := 0; w < FIELD_WIDTH; w++ {
            f.board[dstL][w] = f.board[srcL][w]
        }
    }

    copySrcL := func(nFixed int, eMap []bool) int {
        for h := FIELD_HEIGHT-(1+nFixed); h >= 0; h-- {
            if !eMap[h] {
                return h
            }
        }
        return FIELD_HEIGHT
    }(nFixedLine, erasedMap)

    // No drop and erase lines only.
    if copySrcL >= FIELD_HEIGHT {
        return
    }

    sort.Sort(sort.IntSlice(erasedLines))
    copyDstL := erasedLines[len(erasedLines)-1]

    if copySrcL >= copyDstL {
        panic("Fail to drop line: copying")
    }

    // Assume that copySrcL is greater than copyDstL.
    for copySrcL >= 0 && copyDstL >= 0 {
        if copySrcL < 0{
            f.eraseLine(copyDstL)
        }
        copyLine(copyDstL, copySrcL)
        copyDstL--; copySrcL--
        for copySrcL >= 0 && erasedMap[copySrcL] {
            copySrcL--
        }

        for copySrcL < 0 && copyDstL >= 0 {
            f.eraseLine(copyDstL)
            copyDstL--
        }
    }

    f.draw()
}
