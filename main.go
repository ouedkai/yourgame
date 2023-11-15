// game.go
package yourgame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game interface represents the game contract.
type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

// GameImpl struct implements the Game interface.
type GameImpl struct {
	x, y float64
}

// Update implements the Game interface.
func (g *GameImpl) Update() error {
	// Implement update logic here
	return nil
}

// Draw implements the Game interface.
func (g *GameImpl) Draw(screen *ebiten.Image) {
	// Implement draw logic here
}

// Layout implements the Game interface.
func (g *GameImpl) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Implement layout logic here
	return 320, 240
}
