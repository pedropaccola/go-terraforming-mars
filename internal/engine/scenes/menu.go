package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/pedropaccola/go-terraforming-mars/internal/engine"
)

const (
	// Menu item constants
	menuItemWidth  = 280
	menuItemHeight = 48
	menuItemGap    = 16
)

// MenuItem represents a menu button
type MenuItem struct {
	Label   string
	Action  func()
	bounds  struct{ x, y, w, h float32 }
	hovered bool
}

// MenuScene displays the main menu
type MenuScene struct {
	engine *engine.Engine
	items  []*MenuItem
}

// NewMenuScene creates a new main menu scene
func NewMenuScene(eng *engine.Engine) *MenuScene {
	ms := &MenuScene{
		engine: eng,
		items: []*MenuItem{
			{Label: "New Game", Action: func() { eng.SetScene("game") }},
			{Label: "Load Game", Action: func() { /* TODO */ }},
			{Label: "Settings", Action: func() { /* TODO */ }},
			{Label: "Quit", Action: func() { eng.Quit() }},
		},
	}

	return ms
}

// Called when entering the menu scene
func (ms *MenuScene) OnEnter() {
}

// Called when leaving the menu scene
func (ms *MenuScene) OnExit() {
}

// Called every frame when the scene is active
func (ms *MenuScene) Update() error {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ms.engine.Quit()
	}

	for _, item := range ms.items {
		b := item.bounds
		item.hovered = float32(mx) >= b.x && float32(mx) <= b.x+b.w &&
			float32(my) >= b.y && float32(my) <= b.y+b.h

		if item.hovered && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if item.Action != nil {
				item.Action()
			}
		}
	}

	return nil
}

// Called every frame when the scene is active
func (ms *MenuScene) Draw(screen *ebiten.Image) {
	ms.drawBackground(screen)
	ms.drawTitle(screen)
	ms.drawMenuItems(screen)
	ms.drawFooter(screen)
}

// Draws the background of the menu scene
func (ms *MenuScene) drawBackground(screen *ebiten.Image) {
	screen.Fill(engine.ColorBackground)

	sw := float32(engine.ScreenWidth)
	sh := float32(engine.ScreenHeight)

	// Overlay at top
	vector.FillRect(screen, 0, 0, sw, sh/4, color.RGBA{R: 100, G: 50, B: 30, A: 128}, false)

	// Planet silhouette in background
	vector.FillCircle(screen, sw*0.8, sh*0.9, 200, color.RGBA{R: 40, G: 25, B: 20, A: 255}, false)
	vector.FillCircle(screen, sw*0.8, sh*0.9, 190, color.RGBA{R: 60, G: 35, B: 25, A: 255}, false)
}

// Draws the title of the menu scene
func (ms *MenuScene) drawTitle(screen *ebiten.Image) {
	title := "TERRAFORMING MARS"
	subtitle := "A Go Study Project"

	sw := engine.ScreenWidth
	sh := engine.ScreenHeight

	// Position: approximate centering (approximating 6px per char)
	titleX := (sw / 2) - (len(title) * 3)
	titleY := sh / 8
	subtitleX := (sw / 2) - (len(subtitle) * 3)
	subtitleY := titleY + 28

	// debug text (simple but works)
	ebitenutil.DebugPrintAt(screen, title, titleX, titleY)
	ebitenutil.DebugPrintAt(screen, subtitle, subtitleX, subtitleY)
}

// Draws the menu items
func (ms *MenuScene) drawMenuItems(screen *ebiten.Image) {
	centerX := float32(engine.ScreenWidth) / 2
	startY := float32(200) + menuItemHeight

	for i, item := range ms.items {
		x := centerX - menuItemWidth/2
		y := startY + float32(i)*(menuItemHeight+menuItemGap)

		// Store bounds for hit detection
		item.bounds.x = x
		item.bounds.y = y
		item.bounds.w = menuItemWidth
		item.bounds.h = menuItemHeight

		// Draw button background
		bgColor := engine.ColorButtonBg
		if item.hovered {
			bgColor = engine.ColorButtonHover
		}

		// Draw button rectangle
		vector.FillRect(screen, x, y, menuItemWidth, menuItemHeight, bgColor, false)

		// Draw border
		if item.hovered {
			vector.StrokeRect(screen, x, y, menuItemWidth, menuItemHeight, 2, engine.ColorAccent, false)
		}

		// Draw label (centered in button using debug text)
		labelX := int(centerX) - len(item.Label)*3
		labelY := int(y) + int(menuItemHeight)/2 - 6
		ebitenutil.DebugPrintAt(screen, item.Label, labelX, labelY)
	}
}

func (ms *MenuScene) drawFooter(screen *ebiten.Image) {
	footer := "Press ESC to quit | Made with Ebitengine"
	footerX := engine.ScreenWidth/2 - len(footer)*3
	footerY := engine.ScreenHeight - 50
	ebitenutil.DebugPrintAt(screen, footer, footerX, footerY)
}
