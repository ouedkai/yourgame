package yourgame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game struct represents your game and implements the ebiten.Game interface.
type Game struct {
	x, y float64
}

// Update is called every frame (60 times per second) to update the game state.
func (g *Game) Update() error {
	// Move the square to the right when the screen is touched or the left mouse button is pressed.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.x += 10
	}

	return nil
}

// Draw is called every frame (60 times per second) to draw the game.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw a square at the current position (g.x, g.y).
	ebitenutil.DrawRect(screen, g.x, g.y, 50, 50, color.White)
}

// Layout is called whenever the window size is changed (but we don't use it in this example).
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240 // Fixed screen size in this example
}
