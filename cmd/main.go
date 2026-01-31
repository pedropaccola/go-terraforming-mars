package main

import (
	"fmt"
	"log"

	"github.com/pedropaccola/go-terraforming-mars/internal/game"
)

func main() {
	path := "internal/game/boards/tharsis/board.yaml"

	board, err := game.NewBoard(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Board: %s\n", board.Name)
	fmt.Printf("Description: %s\n\n", board.Description)
	for i, h := range board.HexCoordinates {
		fmt.Printf("- %q: %q\n", i, h)
	}
	for i, h := range board.HexAreas {
		fmt.Printf("- %q: %q\n", i, h)
	}
	fmt.Printf("Hexes: %d\n", len(board.HexCoordinates))
	fmt.Printf("Reserved Areas: %d\n", len(board.HexAreas))
}
