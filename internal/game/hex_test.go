package game

import "testing"

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
