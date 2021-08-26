package asserts

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	MIN_H int = 32 + 16
	MAX_H int = 32*10 + 16
	MIN_W int = 32 + 16
	MAX_W int = 560
	// MAX_W int = 592
)

func (g *SpiteBoy) MoveFount() {
	if g.framevalid%30 != 3 {
		g.framevalid++
		return
	}
	g.framevalid = 0

	g.BoyX = 3
	if g.BoyY == 0 {
		g.BoyY = 2
	} else {
		g.BoyY = 0
	}

	if g.h < MIN_H {
		return
	}
	g.h -= 16
}

func (g *SpiteBoy) MoveDown() {
	if g.framevalid%30 != 3 {
		g.framevalid++
		return
	}
	g.framevalid = 1

	g.BoyX = 2
	if g.BoyY == 0 {
		g.BoyY = 2
	} else {
		g.BoyY = 0
	}

	if g.h > MAX_H {
		return
	}
	// fmt.Println("-----------", g.h)
	g.h += 16
}

func (g *SpiteBoy) MoveLeft() {
	if g.framevalid%30 != 3 {
		g.framevalid++
		return
	}
	g.framevalid = 0

	g.BoyX = 0
	if g.BoyY == 0 {
		g.BoyY = 2
	} else {
		g.BoyY = 0
	}
	if g.w < MIN_W {
		return
	}

	g.w -= 16
}

func (g *SpiteBoy) MoveRight() {
	if g.framevalid%30 != 3 {
		g.framevalid++
		return
	}
	g.framevalid = 0

	g.BoyX = 1
	if g.BoyY == 0 {
		g.BoyY = 2
	} else {
		g.BoyY = 0
	}

	if g.w > MAX_W {
		return
	}
	g.w += 16
}

func isAnyKeyJustPressed() bool {
	for _, k := range inpututil.PressedKeys() {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}

func (g *SpiteBoy) MoveOn() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.MoveDown()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.MoveFount()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.MoveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.MoveRight()
	}
}
