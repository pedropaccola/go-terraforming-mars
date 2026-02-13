package game

import (
	"testing"
)

func countResources(resources []Resource, target Resource) int {
	count := 0
	for _, r := range resources {
		if r == target {
			count++
		}
	}
	return count
}

func TestBoardName(t *testing.T) {
	t.Run("String returns correct names", func(t *testing.T) {
		tests := []struct {
			name     BoardName
			expected string
		}{
			{BoardNameTharsis, "Tharsis"},
		}
		for _, tc := range tests {
			if tc.name.String() != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, tc.name.String())
			}
		}
	})
}

func TestBoard(t *testing.T) {
	t.Run("NewBoard loads Tharsis board", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)
		if board.Name != BoardNameTharsis {
			t.Errorf("Expected name %v, got %v", BoardNameTharsis, board.Name)
		}
	})

	t.Run("NewBoard initializes full hex grid", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)
		// A radius-4 hex grid should have 61 hexes
		expectedHexCount := 61
		if len(board.Hexes) != expectedHexCount {
			t.Errorf("Expected %d hexes, got %d", expectedHexCount, len(board.Hexes))
		}
	})

	t.Run("NewBoard loads special hexes correctly", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)

		// Check Noctis City (q=-2, r=0)
		noctis := board.Hexes[NewHex(-2, 0)]
		if noctis.PlacementRule != PlacementRuleCity {
			t.Errorf("Noctis City should have PlacementRuleCity, got %v", noctis.PlacementRule)
		}
		if noctis.Description != "Noctis City" {
			t.Errorf("Expected description 'Noctis City', got %q", noctis.Description)
		}

		// Check Tharsis Tholus (q=0, r=-3)
		tholus := board.Hexes[NewHex(0, -3)]
		if tholus.PlacementRule != PlacementRuleVolcanic {
			t.Errorf("Tharsis Tholus should have PlacementRuleVolcanic, got %v", tholus.PlacementRule)
		}

		// Check an ocean hex (q=1, r=-4)
		ocean := board.Hexes[NewHex(1, -4)]
		if ocean.PlacementRule != PlacementRuleOcean {
			t.Errorf("Expected PlacementRuleOcean, got %v", ocean.PlacementRule)
		}
	})

	t.Run("NewBoard loads placement bonuses", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)

		// Check hex at (0, -4) has steel: 2
		hex := board.Hexes[NewHex(0, -4)]
		if count := countResources(hex.PlacementBonuses, ResourceSteel); count != 2 {
			t.Errorf("Expected 2 steel bonus, got %d", count)
		}

		// Check hex at (1, -1) has plants: 2
		hex2 := board.Hexes[NewHex(1, -1)]
		if count := countResources(hex2.PlacementBonuses, ResourcePlants); count != 2 {
			t.Errorf("Expected 2 plants bonus, got %d", count)
		}
	})

	t.Run("Center hex exists", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)
		center := NewHex(0, 0)
		if _, exists := board.Hexes[center]; !exists {
			t.Error("Center hex (0, 0) should exist")
		}
	})

	t.Run("Edge hexes exist", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)
		// Test some edge hexes at radius 4
		edgeHexes := []Hex{
			NewHex(4, 0),
			NewHex(-4, 0),
			NewHex(0, 4),
			NewHex(0, -4),
			NewHex(4, -4),
			NewHex(-4, 4),
		}

		for _, h := range edgeHexes {
			if _, exists := board.Hexes[h]; !exists {
				t.Errorf("Edge hex %v should exist", h)
			}
		}
	})

	t.Run("Out of bounds hexes do not exist", func(t *testing.T) {
		board := NewBoard(BoardNameTharsis)
		// These should NOT exist (outside radius 4)
		invalidHexes := []Hex{
			NewHex(5, 0),
			NewHex(0, 5),
			NewHex(-5, 0),
			NewHex(3, 3), // q+r > MapRadius
		}

		for _, h := range invalidHexes {
			if _, exists := board.Hexes[h]; exists {
				t.Errorf("Hex %v should not exist (out of bounds)", h)
			}
		}
	})

	t.Run("initEmptyHexes creates correct number of hexes", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}
		board.initEmptyHexes()

		// Formula for hex count in a radius-r grid: 3*r^2 + 3*r + 1
		expectedCount := 3*MapRadius*MapRadius + 3*MapRadius + 1
		if len(board.Hexes) != expectedCount {
			t.Errorf("Expected %d hexes, got %d", expectedCount, len(board.Hexes))
		}
	})

	t.Run("initEmptyHexes does not overwrite existing hexes", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}

		// Pre-populate a hex with custom data
		specialHex := NewHex(0, 0)
		board.Hexes[specialHex] = HexMetadata{
			Description:   "Special",
			PlacementRule: PlacementRuleOcean,
		}

		board.initEmptyHexes()

		// Verify the special hex was not overwritten
		if board.Hexes[specialHex].Description != "Special" {
			t.Error("initEmptyHexes should not overwrite existing hexes")
		}
		if board.Hexes[specialHex].PlacementRule != PlacementRuleOcean {
			t.Error("initEmptyHexes should not overwrite existing hex placement rules")
		}
	})

	t.Run("initEmptyHexes new hexes have zero-value metadata", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}
		board.initEmptyHexes()

		// Check a hex that wasn't pre-defined
		hex := board.Hexes[NewHex(1, 1)]
		if hex.PlacementRule != PlacementRuleNone {
			t.Error("New hexes should have PlacementRuleNone")
		}
		if hex.Description != "" {
			t.Error("New hexes should have empty description")
		}
	})
}
