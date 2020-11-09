package main

type Coord [2]int

type Mino struct {
    form Form
    rot Rot
    coords [4]Coord
}

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

type Rot uint8
const (
    rot0 Rot = iota
    rot90
    rot180
    rot270
)

type Move uint8
const (
    MoveLeft Move = iota
    MoveRight
    MoveDown
)

func (c Coord) getHeight() int {
    return c[0]
}

func (c Coord) getWidth() int {
    return c[1]
}

func (c Coord) isExceedTop() bool {
    return c.getHeight() < 0
}

func (c Coord) left() Coord {
    h := c.getHeight()
    w := c.getWidth() - 1
    return Coord{ h, w }
}

func (c Coord) right() Coord {
    h := c.getHeight()
    w := c.getWidth() + 1
    return Coord{ h, w }
}

func (c Coord) down() Coord {
    h := c.getHeight() + 1
    w := c.getWidth()
    return Coord{ h, w }
}

func (m *Mino) left() {
    for i, c := range m.coords {
        m.coords[i] = c.left()
    }
}

func (m *Mino) right() {
    for i, c := range m.coords {
        m.coords[i] = c.right()
    }
}

func (m *Mino) down() {
    for i, c := range m.coords {
        m.coords[i] = c.down()
    }
}

func newMino(f Form, r Rot) *Mino {
    return &Mino{ form: f, rot: r }
}

func (m *Mino) copy() *Mino {
    new_m := new(Mino)
    new_m.form = m.form
    new_m.rot = m.rot
    new_m.coords = m.coords
    return new_m
}

func (m *Mino) color() Color {
    switch m.form {
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

    panic("Tetrimino has unknown color")
}

func getInitPosMino(f Form) *Mino {
    m := newMino(f, rot0)

    var coords [4]Coord
    switch m.form {
    case I:
        coords = [4]Coord{
            Coord{0, 3}, Coord{0, 4},
            Coord{0, 5}, Coord{0, 6},
        }
    case O:
        coords = [4]Coord{
            Coord{-1, 4}, Coord{-1, 5},
            Coord{0, 4}, Coord{0, 5},
        }
    case S:
        coords = [4]Coord{
            Coord{-1, 4}, Coord{-1, 5},
            Coord{0, 3}, Coord{0, 4},
        }
    case Z:
        coords = [4]Coord{
            Coord{-1, 3}, Coord{-1, 4},
            Coord{0, 4}, Coord{0, 5},
        }
    case J:
        coords = [4]Coord{
            Coord{-1, 3}, Coord{0, 3},
            Coord{0, 4}, Coord{0, 5},
        }
    case L:
        coords = [4]Coord{
            Coord{0, 3}, Coord{0, 4},
            Coord{0, 5}, Coord{-1, 5},
        }
    case T:
        coords = [4]Coord{
            Coord{-1, 4}, Coord{0, 3},
            Coord{0, 4}, Coord{0, 5},
        }
    }
    m.coords = coords
    return m
}

func (m *Mino) move(move Move) *Mino {
    new_m := m.copy()
    switch move {
    case MoveLeft:
        new_m.left()
    case MoveRight:
        new_m.right()
    case MoveDown:
        new_m.down()
    }
    return new_m
}
