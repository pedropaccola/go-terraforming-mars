package game

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// The size of the hexagonal game board
const MapRadius = 4

// Represents a single hex cell on the board using axial coordinates.
// Q, R, S are cube coordinates (S is redundant but stored for convenience).
// Resources from: https://www.redblobgames.com/grids/hexagons/
type Hex struct {
	Q int `yaml:"q"`
	R int `yaml:"r"`
	S int `yaml:"s"`
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

// Stringer for Hex.
func (h Hex) String() string {
	return fmt.Sprintf("{%d, %d, %d}", h.Q, h.R, h.S)
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
	PlacementRuleNone:     "none",
	PlacementRuleOcean:    "ocean",
	PlacementRuleVolcanic: "volcanic",
	PlacementRuleCity:     "city",
}

var placementRuleValues = map[string]PlacementRule{
	"":         PlacementRuleNone,
	"none":     PlacementRuleNone,
	"ocean":    PlacementRuleOcean,
	"volcanic": PlacementRuleVolcanic,
	"city":     PlacementRuleCity,
}

// Stringer for PlacementRule.
func (pr PlacementRule) String() string {
	return placementRuleNames[pr]
}

// UnmarshalText implements encoding.TextUnmarshaler for YAML decoding.
func (pr *PlacementRule) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	val, ok := placementRuleValues[str]
	if !ok {
		return fmt.Errorf("unknown placement rule: %q", str)
	}
	*pr = val
	return nil
}

type HexMetadata struct {
	Description      string        `yaml:"description"`
	PlacementBonuses ResourceSet   `yaml:"placement_bonuses"`
	PlacementRule    PlacementRule `yaml:"placement_rule"`
}

func (hm HexMetadata) String() string {
	return fmt.Sprintf("{Description: %q, PlacementBonuses: %v, PlacementRule: %q}",
		hm.Description, hm.PlacementBonuses, hm.PlacementRule)
}

type Board struct {
	Name        string
	Description string
	Hexes       map[Hex]HexMetadata
}

// Custom unmarshaler for Board.
func (b *Board) UnmarshalYAML(unmarshal func(any) error) error {
	var temp struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Hexes       []struct {
			Q           int `yaml:"q"`
			R           int `yaml:"r"`
			HexMetadata `yaml:",inline"`
		} `yaml:"hexes"`
	}

	if err := unmarshal(&temp); err != nil {
		return err
	}

	b.Name = temp.Name
	b.Description = temp.Description
	b.Hexes = make(map[Hex]HexMetadata)

	for _, entry := range temp.Hexes {
		coord := NewHex(entry.Q, entry.R)
		b.Hexes[coord] = entry.HexMetadata
	}

	return nil
}

// Creates a new board by loading from a YAML file.
// It initializes all hexes in the grid and overlays special hexes from the YAML.
func NewBoard(path string) (*Board, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read board file: %w", err)
	}

	var board Board
	if err := yaml.Unmarshal(data, &board); err != nil {
		return nil, fmt.Errorf("failed to unmarshal board: %w", err)
	}

	// Initialize full hex grid (fill in missing hexes with empty metadata)
	initializeHexGrid(&board)

	return &board, nil
}

// Fills in missing hexes of the grid with empty metadata.
func initializeHexGrid(board *Board) {
	for q := -MapRadius; q <= MapRadius; q++ {
		r1 := max(-MapRadius, -q-MapRadius)
		r2 := min(MapRadius, -q+MapRadius)
		for r := r1; r <= r2; r++ {
			coord := NewHex(q, r)
			if _, exists := board.Hexes[coord]; !exists {
				board.Hexes[coord] = HexMetadata{}
			}
		}
	}
}
