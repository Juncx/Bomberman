package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	Img    *ebiten.Image
	CoordX int
	CoordY int
}

func (b *Background) LoadImg() {
	f := `D:\chenjun\ebiten\Bomberman\spite\Backgrounds_Playfield.png`
	imgCont, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgCont))
	if err != nil {
		log.Fatal(err)
	}
	b.Img = ebiten.NewImageFromImage(img)
}

func (b *Background) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.CoordX), 0)
	screen.DrawImage(b.Img, op)
}

func (b *Background) InitCoord() {
	b.CoordX = 0
}

func (b *Background) Init() {
	b.LoadImg()
	b.InitCoord()
}
