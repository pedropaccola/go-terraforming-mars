package scenes

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/pedropaccola/go-terraforming-mars/internal/engine"
	"github.com/pedropaccola/go-terraforming-mars/internal/game"
)

const (
	// Hex rendering parameters
	HexSize     = 36.0 // Distance from center to corner
	HexSpacingX = 1.1  // Horizontal spacing multiplier
	HexSpacingY = 1.1  // Vertical spacing multiplier

	PlacementBonusSize    = 10.0
	PlacementBonusSpacing = 2.0
	PlacementBonusOffset  = 20.0
)

// GameScene displays the main game board
type GameScene struct {
	engine      *engine.Engine
	board       *game.Board
	hoveredHex  *game.Hex
	boardOffset struct{ x, y float64 } // Camera offset for panning
}

// NewGameScene creates a new game scene with the given board
func NewGameScene(eng *engine.Engine, board *game.Board) *GameScene {
	return &GameScene{
		engine: eng,
		board:  board,
		boardOffset: struct{ x, y float64 }{
			x: float64(engine.ScreenWidth) / 2,
			y: float64(engine.ScreenHeight) / 2,
		},
	}
}

// Initialize game state when entering
func (g *GameScene) OnEnter() {
}

// Cleanup when leaving
func (g *GameScene) OnExit() {
}

func (g *GameScene) Update() error {
	mx, my := ebiten.CursorPosition()

	// Convert screen position to hex coordinates
	g.hoveredHex = g.screenToHex(float64(mx), float64(my))

	// ESC to return to menu
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.engine.SetScene("menu")
	}

	// Simple panning with arrow keys
	panSpeed := 5.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.boardOffset.x += panSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.boardOffset.x -= panSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.boardOffset.y += panSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.boardOffset.y -= panSpeed
	}

	return nil
}

// Called every frame when the scene is active
func (g *GameScene) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.drawBoard(screen)
	g.drawUI(screen)
}

// Draws the background of the game scene
func (g *GameScene) drawBackground(screen *ebiten.Image) {
	screen.Fill(engine.ColorBoardBg)
}

// Draws the board of the game scene
func (g *GameScene) drawBoard(screen *ebiten.Image) {
	for coord, metadata := range g.board.Hexes {
		g.drawHex(screen, coord, metadata)
	}
}

// Draws a single hex on the screen
func (g *GameScene) drawHex(screen *ebiten.Image, coord game.Hex, metadata game.HexMetadata) {
	// Convert axial coordinates to screen position
	// Using pointy-top hex orientation
	x, y := g.hexToScreen(coord)

	// Determine hex color based on type
	fillColor := g.getHexColor(metadata.PlacementRule)

	// Check if this hex is hovered
	isHovered := g.hoveredHex != nil && g.hoveredHex.Equals(coord)
	if isHovered {
		// Brighten the color when hovered
		fillColor = brightenColor(fillColor, 1.3)
	}

	// Draw the hexagon
	g.drawHexagon(screen, float32(x), float32(y), HexSize, fillColor)

	// Draw bonus indicators (small dots for resources)
	g.drawBonusIndicators(screen, x, y, metadata.PlacementBonuses)
}

// Converts axial hex coordinates to screen pixel coordinates
// Using pointy-top orientation https://www.redblobgames.com/grids/hexagons/#hex-to-pixel
func (g *GameScene) hexToScreen(coord game.Hex) (float64, float64) {
	// Pointy-top hex layout
	// x = size * (sqrt(3) * q + sqrt(3)/2 * r)
	// y = size * (				      3./2 * r)
	sqrt3 := math.Sqrt(3)
	x := HexSize * HexSpacingX * (sqrt3*float64(coord.Q) + sqrt3/2*float64(coord.R))
	y := HexSize * HexSpacingY * (3.0 / 2.0 * float64(coord.R))

	// Apply board offset (centering)
	x += g.boardOffset.x
	y += g.boardOffset.y

	return x, y
}

// Converts screen pixel coordinates to axial hex coordinates
// https://www.redblobgames.com/grids/hexagons/#pixel-to-hex
func (g *GameScene) screenToHex(screenX, screenY float64) *game.Hex {
	// Remove offset
	x := screenX - g.boardOffset.x
	y := screenY - g.boardOffset.y

	// Reverse the hex-to-screen calculation
	sqrt3 := math.Sqrt(3)
	q := (sqrt3/3*x - 1.0/3*y) / (HexSize * HexSpacingX)
	r := (2.0 / 3 * y) / (HexSize * HexSpacingY)

	// Round to nearest hex (cube coordinate rounding)
	return g.roundHex(q, r)
}

// Rounds fractional hex coordinates to the nearest hex
// https://www.redblobgames.com/grids/hexagons/#rounding
func (g *GameScene) roundHex(q, r float64) *game.Hex {
	s := -q - r

	rq := math.Round(q)
	rr := math.Round(r)
	rs := math.Round(s)

	qDiff := math.Abs(rq - q)
	rDiff := math.Abs(rr - r)
	sDiff := math.Abs(rs - s)

	if qDiff > rDiff && qDiff > sDiff {
		rq = -rr - rs
	} else if rDiff > sDiff {
		rr = -rq - rs
	}

	coord := game.NewHex(int(rq), int(rr))

	// Return nil if the hex does not exist on the board
	if _, exists := g.board.Hexes[coord]; !exists {
		return nil
	}
	return &coord
}

// Returns the color for a given hex reservation
// TODO: Tiles placed on the hex should override the hex color
func (g *GameScene) getHexColor(hexReservation game.PlacementRule) color.RGBA {
	switch hexReservation {
	case game.PlacementRuleOcean:
		return engine.ColorHexOcean
	case game.PlacementRuleVolcanic:
		return engine.ColorHexVolcanic
	case game.PlacementRuleCity:
		return engine.ColorHexCity
	default:
		return engine.ColorHexDefault
	}
}

// Draws a hexagon at the given position (hexagon center)
// https://www.redblobgames.com/grids/hexagons/#angles
func (g *GameScene) drawHexagon(screen *ebiten.Image, cx, cy, size float32, fill color.RGBA) {
	// Calculate the 6 vertices of a pointy-top hexagon
	var path vector.Path

	for i := 0; i < 6; i++ {
		angle := math.Pi / 180 * float64(60*i-30)
		vx := cx + size*float32(math.Cos(angle))
		vy := cy + size*float32(math.Sin(angle))

		if i == 0 {
			path.MoveTo(vx, vy)
		} else {
			path.LineTo(vx, vy)
		}
	}
	path.Close()

	// Fill the hexagon
	var colorScale ebiten.ColorScale
	colorScale.Scale(float32(fill.R)/255, float32(fill.G)/255, float32(fill.B)/255, float32(fill.A)/255)

	vector.FillPath(screen, &path, nil, &vector.DrawPathOptions{
		ColorScale: colorScale,
	})
}

// Draws bonus indicators for a hex
func (g *GameScene) drawBonusIndicators(screen *ebiten.Image, cx, cy float64, bonuses []game.Resource) {
	if len(bonuses) == 0 {
		return
	}

	for i, res := range bonuses {
		clr, ok := engine.ResourceColors[res]
		if !ok {
			continue
		}

		ix := float32(cx) - PlacementBonusOffset + float32(i)*(PlacementBonusSize+PlacementBonusSpacing)
		iy := float32(cy) - PlacementBonusOffset
		vector.FillRect(screen, ix, iy, PlacementBonusSize, PlacementBonusSize, clr, false)
		vector.StrokeRect(screen, ix, iy, PlacementBonusSize, PlacementBonusSize, 1, engine.ColorBorder, false)
	}
}

func (g *GameScene) drawUI(screen *ebiten.Image) {
	// TODO: Draw UI
}

// brightenColor multiplies the RGB values by a factor
// Used to highlight hovered hexes
func brightenColor(c color.RGBA, factor float64) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(255, float64(c.R)*factor)),
		G: uint8(math.Min(255, float64(c.G)*factor)),
		B: uint8(math.Min(255, float64(c.B)*factor)),
		A: c.A,
	}
}
