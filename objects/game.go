package objects 

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Game struct {
//	input *Input 
	boardImage *ebiten.Image  
}
func (game *Game) Update() error {
	return nil
}
func (game *Game) Draw(screen *ebiten.Image) {
	if game.boardImage == nil {
		game.boardImage = ebiten.NewImage(50, 50)
		game.boardImage.Fill(color.RGBA{0x33, 0x66, 0xCC, 0xFF})
	}
	screen.Fill(color.RGBA{0xfa, 0xf8, 0xef, 0xff})

	sw, sh := screen.Size()
	ow, oh := game.boardImage.Size()

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.7)
	op.GeoM.Translate(float64((sw-ow)/2), float64((sh-oh)/2))
	screen.DrawImage(game.boardImage, op) 
}
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight 
}

