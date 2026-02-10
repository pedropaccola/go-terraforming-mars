package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// Window dimensions
	ScreenWidth  = 1280
	ScreenHeight = 800
)

// Scene represents a game scene (menu, game board, etc.)
// Each scene handles its own update logic and rendering.
type Scene interface {
	// Update handles input and game logic. Called 60 times per second.
	Update() error

	// Draw renders the scene to the screen.
	Draw(screen *ebiten.Image)

	// OnEnter is called when transitioning TO this scene.
	OnEnter()

	// OnExit is called when transitioning AWAY from this scene.
	OnExit()
}

// Engine is the main game engine that manages scenes and the game loop.
type Engine struct {
	currentScene Scene
	scenes       map[string]Scene
	shouldQuit   bool
}

// NewEngine creates a new game engine.
func NewEngine() *Engine {
	return &Engine{
		scenes: make(map[string]Scene),
	}
}

// RegisterScene adds a scene to the engine with a given name.
func (e *Engine) RegisterScene(name string, scene Scene) {
	e.scenes[name] = scene
}

// SetScene switches to a different scene by name.
func (e *Engine) SetScene(name string) {
	if scene, ok := e.scenes[name]; ok {
		if e.currentScene != nil {
			e.currentScene.OnExit()
		}
		e.currentScene = scene
		e.currentScene.OnEnter()
	}
}

// Update implements ebiten.Game interface.
// This is called 60 times per second (by default).
func (e *Engine) Update() error {
	if e.shouldQuit {
		return ebiten.Termination
	}
	if e.currentScene != nil {
		return e.currentScene.Update()
	}
	return nil
}

// Draw implements ebiten.Game interface.
// This is called every frame after Update.
func (e *Engine) Draw(screen *ebiten.Image) {
	if e.currentScene != nil {
		e.currentScene.Draw(screen)
	}
}

// Layout implements ebiten.Game interface.
// It returns the logical screen size.
func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// Run starts the game engine.
func (e *Engine) Run(title string) error {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(e)
}

// Quit stops the game engine gracefully.
func (e *Engine) Quit() {
	e.shouldQuit = true
}
