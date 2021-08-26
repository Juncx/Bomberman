package main

import (
	"bomberman/asserts"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	spiteboy  asserts.SpiteBoy
	backgroud Background
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	//g.spiteboy.MoveFount()
	// g.MoveLeft()
	//g.MoveRight()

	g.spiteboy.MoveOn()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	// rect := image.Rect(0, 0, 40, 80)
	// s := ebiten.NewImageFromImage(g.hero.SubImage(rect))
	g.backgroud.Draw(screen)
	g.spiteboy.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
	// return 640, 480
}

func main() {
	game := &Game{}
	game.spiteboy.Init()
	game.backgroud.Init()
	// Specify the window size as you like. Here, a doubled size is specified.
	// ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Your game's title")
	// ebiten.SetMaxTPS(5)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
