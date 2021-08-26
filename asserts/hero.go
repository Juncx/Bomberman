package asserts

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpiteBoy struct {
	LRect   [3]image.Rectangle
	DRect   [3]image.Rectangle
	RRect   [3]image.Rectangle
	URect   [3]image.Rectangle
	AllRect [4][3]image.Rectangle

	Img    *ebiten.Image
	CoordX int
	CoordY int

	BoyX int
	BoyY int

	h int
	w int

	framevalid int
}

func (s *SpiteBoy) initRect() {
	itemX := 16
	itemY := 16
	// offsetX := 1
	// offsetY := 0

	for i := 0; i != 3; i++ {
		s.LRect[i] = image.Rect(i*itemX, 0, (i+1)*itemX, itemY)
	}
	for i := 0; i != 3; i++ {
		s.RRect[i] = image.Rect(i*itemX, itemY, (i+1)*itemX, 2*itemY)
	}
	for i := 0; i != 3; i++ {
		s.URect[i] = image.Rect((i+3)*itemX, 0, (i+1+3)*itemX, itemY)
	}
	for i := 0; i != 3; i++ {
		s.DRect[i] = image.Rect((i+3)*itemX, itemY, (i+1+3)*itemX, 2*itemY)
	}

	s.AllRect[0] = s.LRect
	s.AllRect[1] = s.RRect
	s.AllRect[2] = s.URect
	s.AllRect[3] = s.DRect
}

func (s *SpiteBoy) Init() {
	s.LoadImg()
	s.InitCoord()
}

func (s *SpiteBoy) InitCoord() {
	s.w = 32
	s.h = 32
}

func (s *SpiteBoy) LoadImg() {
	f := `D:\chenjun\ebiten\Bomberman\spite\Miscellaneous_General_Sprites.png`
	imgCont, _ := ioutil.ReadFile(f)
	img, _, err := image.Decode(bytes.NewReader(imgCont))
	if err != nil {
		log.Fatal(err)
	}
	s.initRect()
	s.Img = ebiten.NewImageFromImage(img)
}

func (s *SpiteBoy) ImageCurr(x, y int) *ebiten.Image {
	return ebiten.NewImageFromImage(s.Img.SubImage(s.AllRect[x][y]))
}

func (s *SpiteBoy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	//op.GeoM.Rotate(float64(g.vy16) / 96.0 * math.Pi / 6)
	op.GeoM.Translate(float64(s.w)/2.0, float64(s.h)/2.0)
	screen.DrawImage(s.ImageCurr(s.BoyX, s.BoyY), op)
	// screen.DrawImage(s.Img, nil)
}
