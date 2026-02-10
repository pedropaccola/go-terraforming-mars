package main

import (
	"log"

	"github.com/pedropaccola/go-terraforming-mars/internal/engine"
	"github.com/pedropaccola/go-terraforming-mars/internal/engine/scenes"
	"github.com/pedropaccola/go-terraforming-mars/internal/game"
)

func main() {
	// Load the Tharsis board
	board, err := game.NewBoard("internal/game/boards/tharsis/board.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Create the game engine
	eng := engine.NewEngine()

	// Create and register scenes
	gameScene := scenes.NewGameScene(eng, board)
	menuScene := scenes.NewMenuScene(eng)

	eng.RegisterScene("menu", menuScene)
	eng.RegisterScene("game", gameScene)

	// Start at the menu
	eng.SetScene("menu")

	// Run the game!
	if err := eng.Run("Terraforming Mars"); err != nil {
		log.Fatal(err)
	}
}
