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

func newMino(f Form) *Mino {
    return &Mino{ form: f, rot: rot0 }
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
    m := newMino(f)

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

