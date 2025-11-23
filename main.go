package main 
import (
	"log"
	"fmt" 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/AnshulDekate/flappyBird/objects"
	"time"
)

func main () {
	fmt.Println("start of side projects with flappy bird")
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("flappy bird")
	game := &objects.Game{} 
	go func () {
		for {
			fmt.Println("in the loop")
			time.Sleep(time.Second)
			game.Pos = 0 
		}
	}()
	err := ebiten.RunGame(game) 
	
	if err != nil {
		log.Fatal(err)
	}
}
