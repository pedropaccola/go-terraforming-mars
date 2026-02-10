package engine

import (
	"image/color"

	"github.com/pedropaccola/go-terraforming-mars/internal/game"
)

var (
	// Background colors
	ColorBackground = color.RGBA{R: 20, G: 15, B: 25, A: 255} // Deep space dark
	ColorMenuBg     = color.RGBA{R: 30, G: 25, B: 40, A: 255} // Slightly lighter
	ColorBoardBg    = color.RGBA{R: 40, G: 30, B: 20, A: 255} // Mars surface tone
	ColorBorder     = color.RGBA{R: 60, G: 50, B: 40, A: 255} // Border

	// Hex colors by type
	ColorHexDefault  = color.RGBA{R: 180, G: 120, B: 80, A: 255}  // Mars terrain
	ColorHexOcean    = color.RGBA{R: 40, G: 100, B: 180, A: 255}  // Ocean blue
	ColorHexVolcanic = color.RGBA{R: 140, G: 60, B: 40, A: 255}   // Volcanic red-brown
	ColorHexCity     = color.RGBA{R: 100, G: 100, B: 110, A: 255} // City gray

	// Resource colors
	ColorResourceAnimals     = color.RGBA{R: 35, G: 135, B: 100, A: 255} // Green for animals
	ColorResourceCards       = color.RGBA{R: 75, G: 55, B: 25, A: 255}   // Black for cards
	ColorResourceEnergy      = color.RGBA{R: 140, G: 50, B: 130, A: 255} // Purple for energy
	ColorResourceHeat        = color.RGBA{R: 220, G: 90, B: 50, A: 255}  // Red for heat
	ColorResourceMegacredits = color.RGBA{R: 255, G: 220, B: 20, A: 255} // Yellow for megacredits
	ColorResourceMicrobes    = color.RGBA{R: 200, G: 215, B: 85, A: 255} // Green for microbes
	ColorResourcePlants      = color.RGBA{R: 50, G: 140, B: 50, A: 255}  // Green for plants
	ColorResourceSteel       = color.RGBA{R: 139, G: 90, B: 43, A: 255}  // Brown for steel
	ColorResourceTitanium    = color.RGBA{R: 95, G: 95, B: 95, A: 255}   // Silver for titanium

	// UI colors
	ColorText        = color.RGBA{R: 240, G: 235, B: 220, A: 255} // Off-white
	ColorTextDim     = color.RGBA{R: 150, G: 145, B: 130, A: 255} // Dimmed text
	ColorAccent      = color.RGBA{R: 220, G: 120, B: 50, A: 255}  // Orange accent
	ColorButtonBg    = color.RGBA{R: 60, G: 50, B: 70, A: 255}    // Button background
	ColorButtonHover = color.RGBA{R: 80, G: 70, B: 100, A: 255}   // Button hover
)

var ResourceColors = map[game.Resource]color.Color{
	game.ResourceAnimals:     ColorResourceAnimals,
	game.ResourceCards:       ColorResourceCards,
	game.ResourceEnergy:      ColorResourceEnergy,
	game.ResourceHeat:        ColorResourceHeat,
	game.ResourceMegacredits: ColorResourceMegacredits,
	game.ResourceMicrobes:    ColorResourceMicrobes,
	game.ResourcePlants:      ColorResourcePlants,
	game.ResourceSteel:       ColorResourceSteel,
	game.ResourceTitanium:    ColorResourceTitanium,
}
