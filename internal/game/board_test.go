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

func TestHex(t *testing.T) {
	t.Run("NewHex computes S coordinate", func(t *testing.T) {
		tests := []struct {
			q, r, expected int
		}{
			{0, 0, 0},
			{1, 2, -3},
			{-1, -2, 3},
			{4, -4, 0},
		}
		for _, tc := range tests {
			h := NewHex(tc.q, tc.r)
			if h.S != tc.expected {
				t.Errorf("NewHex(%d, %d): expected S=%d, got S=%d", tc.q, tc.r, tc.expected, h.S)
			}
		}
	})

	t.Run("DistanceFrom calculates hex distance", func(t *testing.T) {
		tests := []struct {
			h1, h2   Hex
			expected int
		}{
			{NewHex(0, 0), NewHex(0, 0), 0},   // Same hex
			{NewHex(0, 0), NewHex(1, 0), 1},   // Adjacent
			{NewHex(0, 0), NewHex(2, 0), 2},   // Two steps
			{NewHex(1, 2), NewHex(4, 5), 6},   // Diagonal
			{NewHex(-2, -2), NewHex(2, 2), 8}, // Opposite corners
		}
		for _, tc := range tests {
			dist := tc.h1.DistanceFrom(tc.h2)
			if dist != tc.expected {
				t.Errorf("DistanceFrom(%v, %v): expected %d, got %d", tc.h1, tc.h2, tc.expected, dist)
			}
		}
	})

	t.Run("DistanceFrom is symmetric", func(t *testing.T) {
		h1, h2 := NewHex(1, 2), NewHex(4, -1)
		if h1.DistanceFrom(h2) != h2.DistanceFrom(h1) {
			t.Error("Distance should be symmetric")
		}
	})

	t.Run("Equals compares coordinates", func(t *testing.T) {
		h1 := NewHex(1, 2)
		h2 := NewHex(1, 2)
		h3 := NewHex(1, 3)

		if !h1.Equals(h2) {
			t.Error("Expected h1 to equal h2")
		}
		if h1.Equals(h3) {
			t.Error("Expected h1 to not equal h3")
		}
	})

	t.Run("Neighbor returns correct adjacent hex", func(t *testing.T) {
		center := NewHex(0, 0)
		expectedNeighbors := []Hex{
			{Q: 1, R: -1, S: 0}, // Direction 0
			{Q: 1, R: 0, S: -1}, // Direction 1
			{Q: 0, R: -1, S: 1}, // Direction 2
			{Q: 0, R: 1, S: -1}, // Direction 3
			{Q: -1, R: 0, S: 1}, // Direction 4
			{Q: -1, R: 1, S: 0}, // Direction 5
		}

		for dir, expected := range expectedNeighbors {
			got := center.Neighbor(dir)
			if !got.Equals(expected) {
				t.Errorf("Neighbor(%d): expected %v, got %v", dir, expected, got)
			}
		}
	})

	t.Run("Neighbor panics on invalid direction", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for direction < 0")
			}
		}()
		NewHex(0, 0).Neighbor(-1)
	})

	t.Run("Neighbor panics on direction > 5", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic for direction > 5")
			}
		}()
		NewHex(0, 0).Neighbor(6)
	})

	t.Run("All neighbors are distance 1", func(t *testing.T) {
		center := NewHex(3, -2)
		for dir := 0; dir < 6; dir++ {
			neighbor := center.Neighbor(dir)
			if center.DistanceFrom(neighbor) != 1 {
				t.Errorf("Neighbor in direction %d should be distance 1", dir)
			}
		}
	})

	t.Run("String format", func(t *testing.T) {
		h := NewHex(1, 2)
		expected := "{1, 2, -3}"
		if h.String() != expected {
			t.Errorf("Expected %q, got %q", expected, h.String())
		}
	})
}

func TestPlacementRule(t *testing.T) {
	t.Run("String returns correct names", func(t *testing.T) {
		tests := []struct {
			rule     PlacementRule
			expected string
		}{
			{PlacementRuleNone, "none"},
			{PlacementRuleOcean, "ocean"},
			{PlacementRuleVolcanic, "volcanic"},
			{PlacementRuleCity, "city"},
		}
		for _, tc := range tests {
			if tc.rule.String() != tc.expected {
				t.Errorf("PlacementRule(%d).String(): expected %q, got %q",
					tc.rule, tc.expected, tc.rule.String())
			}
		}
	})

	t.Run("UnmarshalText parses valid values", func(t *testing.T) {
		tests := []struct {
			input    string
			expected PlacementRule
		}{
			{"", PlacementRuleNone},
			{"none", PlacementRuleNone},
			{"ocean", PlacementRuleOcean},
			{"volcanic", PlacementRuleVolcanic},
			{"city", PlacementRuleCity},
			{"OCEAN", PlacementRuleOcean}, // Case insensitive
			{"Volcanic", PlacementRuleVolcanic},
		}
		for _, tc := range tests {
			var pr PlacementRule
			err := pr.UnmarshalText([]byte(tc.input))
			if err != nil {
				t.Errorf("UnmarshalText(%q): unexpected error: %v", tc.input, err)
			}
			if pr != tc.expected {
				t.Errorf("UnmarshalText(%q): expected %v, got %v", tc.input, tc.expected, pr)
			}
		}
	})

	t.Run("UnmarshalText rejects invalid values", func(t *testing.T) {
		var pr PlacementRule
		err := pr.UnmarshalText([]byte("invalid"))
		if err == nil {
			t.Error("Expected error for invalid placement rule")
		}
	})
}

func TestHexMetadata(t *testing.T) {
	t.Run("String format", func(t *testing.T) {
		hm := HexMetadata{
			Description:      "Test Hex",
			PlacementBonuses: []Resource{ResourceSteel, ResourceSteel, ResourcePlants},
			PlacementRule:    PlacementRuleVolcanic,
		}

		// %v for the slice will use the String() method of Resource
		expected := `{Description: "Test Hex", PlacementBonuses: [steel steel plants], PlacementRule: "volcanic"}`
		if hm.String() != expected {
			t.Errorf("Expected %s, got %s", expected, hm.String())
		}
	})

	t.Run("Zero value has sensible defaults", func(t *testing.T) {
		var hm HexMetadata
		if hm.PlacementRule != PlacementRuleNone {
			t.Error("Zero value should have PlacementRuleNone")
		}
		if hm.Description != "" {
			t.Error("Zero value should have empty description")
		}
	})
}

func TestBoard(t *testing.T) {
	// Note: These tests require the YAML file to exist
	// Run from project root: go test ./internal/game/...

	t.Run("NewBoard loads Tharsis board", func(t *testing.T) {
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

		if board.Name != "tharsis" {
			t.Errorf("Expected name 'tharsis', got %q", board.Name)
		}
	})

	t.Run("NewBoard initializes full hex grid", func(t *testing.T) {
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

		// A radius-4 hex grid should have 61 hexes
		// Formula: 3*r^2 + 3*r + 1 = 3*16 + 12 + 1 = 61
		expectedHexCount := 61
		if len(board.Hexes) != expectedHexCount {
			t.Errorf("Expected %d hexes, got %d", expectedHexCount, len(board.Hexes))
		}
	})

	t.Run("NewBoard loads special hexes correctly", func(t *testing.T) {
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

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
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

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

	t.Run("NewBoard returns error for missing file", func(t *testing.T) {
		_, err := NewBoard("nonexistent/path.yaml")
		if err == nil {
			t.Error("Expected error for missing file")
		}
	})

	t.Run("Center hex exists", func(t *testing.T) {
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

		center := NewHex(0, 0)
		if _, exists := board.Hexes[center]; !exists {
			t.Error("Center hex (0, 0) should exist")
		}
	})

	t.Run("Edge hexes exist", func(t *testing.T) {
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

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
		board, err := NewBoard("boards/tharsis/board.yaml")
		if err != nil {
			t.Fatalf("Failed to load board: %v", err)
		}

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
}

func TestInitializeHexGrid(t *testing.T) {
	t.Run("Creates correct number of hexes", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}
		initializeHexGrid(board)

		// Formula for hex count in a radius-r grid: 3*r^2 + 3*r + 1
		expectedCount := 3*MapRadius*MapRadius + 3*MapRadius + 1
		if len(board.Hexes) != expectedCount {
			t.Errorf("Expected %d hexes, got %d", expectedCount, len(board.Hexes))
		}
	})

	t.Run("Does not overwrite existing hexes", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}

		// Pre-populate a hex with custom data
		specialHex := NewHex(0, 0)
		board.Hexes[specialHex] = HexMetadata{
			Description:   "Special",
			PlacementRule: PlacementRuleOcean,
		}

		initializeHexGrid(board)

		// Verify the special hex was not overwritten
		if board.Hexes[specialHex].Description != "Special" {
			t.Error("initializeHexGrid should not overwrite existing hexes")
		}
		if board.Hexes[specialHex].PlacementRule != PlacementRuleOcean {
			t.Error("initializeHexGrid should not overwrite existing hex placement rules")
		}
	})

	t.Run("New hexes have zero-value metadata", func(t *testing.T) {
		board := &Board{Hexes: make(map[Hex]HexMetadata)}
		initializeHexGrid(board)

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
