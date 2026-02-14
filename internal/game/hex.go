package game

import "fmt"

// Represents a single hex cell on the board using axial coordinates.
// Q, R, S are cube coordinates (S is redundant but stored for convenience).
// Resources from: https://www.redblobgames.com/grids/hexagons/
type Hex struct {
	Q int
	R int
	S int
}

// HexDirections contains the six neighbor direction vectors for axial hex coordinates.
var HexDirections = []Hex{
	{Q: 1, R: -1, S: 0},
	{Q: 1, R: 0, S: -1},
	{Q: 0, R: -1, S: 1},
	{Q: 0, R: 1, S: -1},
	{Q: -1, R: 0, S: 1},
	{Q: -1, R: 1, S: 0},
}

// Creates a new Hex with the given q and r coordinates, computing s as -q-r.
func NewHex(q, r int) Hex {
	return Hex{Q: q, R: r, S: -q - r}
}

// Calculates the hex distance from this hex to another.
func (h Hex) DistanceFrom(other Hex) int {
	vector := h.subtract(other)
	return (max(vector.Q, -vector.Q) + max(vector.R, -vector.R) + max(vector.S, -vector.S)) / 2
}

// Checks if this hex is equal to another hex.
func (h Hex) Equals(other Hex) bool {
	return h.Q == other.Q && h.R == other.R && h.S == other.S
}

// Returns the neighboring hex in the specified direction (0-5).
func (h Hex) Neighbor(direction int) Hex {
	return h.add(h.direction(direction))
}

// fmt.Stringer for Hex.
func (h Hex) String() string {
	return fmt.Sprintf("%d,%d,%d", h.Q, h.R, h.S)
}

// Adds another hex to this one, returning a new hex.
func (h Hex) add(other Hex) Hex {
	return Hex{
		Q: h.Q + other.Q,
		R: h.R + other.R,
		S: h.S + other.S,
	}
}

// Returns the hex direction vector for the given direction index (0-5).
func (h Hex) direction(direction int) Hex {
	if direction < 0 || direction > 5 {
		panic("Direction must be between 0 and 5")
	}
	return HexDirections[direction]
}

// Subtracts another hex from this one, returning a new hex.
func (h Hex) subtract(other Hex) Hex {
	return Hex{
		Q: h.Q - other.Q,
		R: h.R - other.R,
		S: h.S - other.S,
	}
}

// Represents the placement rules of hexes on the board.
// A hex placement rule determines exclusive placements for certain tiles,
// such as cities, oceans and volcanic reserved areas.
type PlacementRule int

const (
	PlacementRuleNone PlacementRule = iota
	PlacementRuleOcean
	PlacementRuleVolcanic
	PlacementRuleCity
)

var placementRuleNames = map[PlacementRule]string{
	PlacementRuleNone:     "None",
	PlacementRuleOcean:    "Ocean",
	PlacementRuleVolcanic: "Volcanic",
	PlacementRuleCity:     "City",
}

// fmt.Stringer for PlacementRule.
func (pr PlacementRule) String() string {
	return placementRuleNames[pr]
}

type HexMetadata struct {
	Description      string
	PlacementBonuses []Resource
	PlacementRule    PlacementRule
}

// fmt.Stringer for HexMetadata.
func (hm HexMetadata) String() string {
	return fmt.Sprintf("Description: %q, PlacementBonuses: %v, PlacementRule: %q",
		hm.Description, hm.PlacementBonuses, hm.PlacementRule)
}
