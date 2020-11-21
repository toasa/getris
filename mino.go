package main

type Coord [2]int

type Mino struct {
    form Form
    rot Rot
    coords [4]Coord

    // To simplify implementation of rotation.
    pivot Coord
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
    MoveDrop
    MoveHardDrop
    RotLeft
    RotRight
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

func (c Coord) isExceedBottom() bool {
    return c.getHeight() >= FIELD_HEIGHT
}

func (c Coord) isExceedSide() bool {
    w := c.getWidth()
    return w < 0 || w >= FIELD_WIDTH
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

func (c Coord) up() Coord {
    h := c.getHeight() - 1
    w := c.getWidth()
    return Coord{ h, w }
}

func (c Coord) down() Coord {
    h := c.getHeight() + 1
    w := c.getWidth()
    return Coord{ h, w }
}

// Get cell coordinates determined by a form, a pivot and
// state of rotation.
func getComposedCoords(f Form, pivot Coord, rot Rot) [4]Coord {
    ph := pivot.getHeight()
    pw := pivot.getWidth()

    var coords [4]Coord
    switch f {
    case I:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph, pw-1}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph, pw+2},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw},
                Coord{ph+1, pw}, Coord{ph+2, pw},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw-2}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph, pw+1},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-2, pw}, Coord{ph-1, pw},
                Coord{ph, pw}, Coord{ph+1, pw},
            }
        }
    case O:
        // A pivot of O form is fixed at upper left cell.
        coords = [4]Coord{
            Coord{ph, pw}, Coord{ph, pw+1},
            Coord{ph+1, pw}, Coord{ph+1, pw+1},
        }
    case S:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph-1, pw+1},
                Coord{ph, pw-1}, Coord{ph, pw},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw+1},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw}, Coord{ph, pw+1},
                Coord{ph+1, pw-1}, Coord{ph+1, pw},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-1, pw-1}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph+1, pw},
            }
        }
    case Z:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph-1, pw-1}, Coord{ph-1, pw},
                Coord{ph, pw}, Coord{ph, pw+1},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw+1}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw-1}, Coord{ph, pw},
                Coord{ph+1, pw}, Coord{ph+1, pw+1},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph+1, pw-1},
            }
        }
    case J:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph-1, pw-1}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph, pw+1},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph-1, pw+1},
                Coord{ph, pw}, Coord{ph+1, pw},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw-1}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw+1},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw},
                Coord{ph+1, pw-1}, Coord{ph+1, pw},
            }
        }
    case L:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph-1, pw+1}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph, pw+1},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw},
                Coord{ph+1, pw}, Coord{ph+1, pw+1},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw-1}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw-1},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-1, pw-1}, Coord{ph-1, pw},
                Coord{ph, pw}, Coord{ph+1, pw},
            }
        }
    case T:
        switch rot {
        case rot0:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph, pw+1},
            }
        case rot90:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw},
            }
        case rot180:
            coords = [4]Coord{
                Coord{ph, pw-1}, Coord{ph, pw},
                Coord{ph, pw+1}, Coord{ph+1, pw},
            }
        case rot270:
            coords = [4]Coord{
                Coord{ph-1, pw}, Coord{ph, pw-1},
                Coord{ph, pw}, Coord{ph+1, pw},
            }
        }
    }
    return coords
}

func (m *Mino) left() {
    for i, c := range m.coords {
        m.coords[i] = c.left()
    }
    m.pivot = m.pivot.left()
}

func (m *Mino) right() {
    for i, c := range m.coords {
        m.coords[i] = c.right()
    }
    m.pivot = m.pivot.right()
}

func (m *Mino) drop() {
    for i, c := range m.coords {
        m.coords[i] = c.down()
    }
    m.pivot = m.pivot.down()
}

func (m Mino) getRotatedPivotForIForm(rot Move) Coord {
    if m.form != I {
        panic("getRotatedPivotForIForm can available only for I form")
    }

    var new_p Coord

    switch rot {
    case RotLeft:
        switch m.rot {
            case rot0:
                new_p = m.pivot.down()
            case rot90:
                new_p = m.pivot.left()
            case rot180:
                new_p = m.pivot.up()
            case rot270:
                new_p = m.pivot.right()
        }
    case RotRight:
        switch m.rot {
            case rot0:
                new_p = m.pivot.right()
            case rot90:
                new_p = m.pivot.down()
            case rot180:
                new_p = m.pivot.left()
            case rot270:
                new_p = m.pivot.up()
        }
    default:
        panic("getRotatedPivotForIForm can handle only rotaion")
    }
    return new_p
}

func (m *Mino) rotLeft() {
    turnLeft := func(r Rot) Rot {
        switch r {
        case rot0:
            return rot270
        case rot90:
            return rot0
        case rot180:
            return rot90
        default:
            return rot180
        }
    }
    new_rot := turnLeft(m.rot)

    switch m.form {
    case I:
        m.pivot = (*m).getRotatedPivotForIForm(RotLeft)
        m.coords = getComposedCoords(m.form, m.pivot, new_rot)
    case O:
        return
    case S, Z, J, L, T:
        m.coords = getComposedCoords(m.form, m.pivot, new_rot)
    }

    m.rot = new_rot
}

func (m *Mino) rotRight() {
    turnRight := func(r Rot) Rot {
        switch r {
        case rot0:
            return rot90
        case rot90:
            return rot180
        case rot180:
            return rot270
        default:
            return rot0
        }
    }
    new_rot := turnRight(m.rot)

    switch m.form {
    case I:
        m.pivot = (*m).getRotatedPivotForIForm(RotRight)
        m.coords = getComposedCoords(m.form, m.pivot, new_rot)
    case O:
        return
    case S, Z, J, L, T:
        m.coords = getComposedCoords(m.form, m.pivot, new_rot)
    }

    m.rot = new_rot
}

func newMino(f Form, r Rot) *Mino {
    return &Mino{ form: f, rot: r }
}

func (m *Mino) copy() *Mino {
    new_m := new(Mino)
    new_m.form = m.form
    new_m.rot = m.rot
    new_m.coords = m.coords
    new_m.pivot = m.pivot
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
    // Every forms that are in initial potision have
    // (0, 4) pivot.
    m.pivot = Coord{0, 4}
    m.coords = getComposedCoords(f, m.pivot, rot0)
    return m
}

func (m *Mino) move(move Move) *Mino {
    new_m := m.copy()
    switch move {
    case MoveLeft:
        new_m.left()
    case MoveRight:
        new_m.right()
    case MoveDrop:
        new_m.drop()
    case MoveHardDrop:
        new_m.drop()
    case RotLeft:
        new_m.rotLeft()
    case RotRight:
        new_m.rotRight()
    }
    return new_m
}
