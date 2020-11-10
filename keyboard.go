package main

import (
    "github.com/veandco/go-sdl2/sdl"
)

type Keycode sdl.Keycode

const (
    KEY_LEFT = sdl.K_LEFT
    KEY_RIGHT = sdl.K_RIGHT
    KEY_DOWN = sdl.K_DOWN
    KEY_ROT_LEFT = sdl.K_z
    KEY_ROT_RIGHT = sdl.K_x
)

func getKeycode(e *sdl.KeyboardEvent) Keycode {
    return Keycode(e.Keysym.Sym)
}
