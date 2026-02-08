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
type HexCoordinates struct {
	Q int `yaml:"q"`
	R int `yaml:"r"`
	S int `yaml:"s"`
}

// HexDirections contains the six neighbor direction vectors for axial hex coordinates.
var HexDirections = []HexCoordinates{
	{Q: 1, R: -1, S: 0},
	{Q: 1, R: 0, S: -1},
	{Q: 0, R: -1, S: 1},
	{Q: 0, R: 1, S: -1},
	{Q: -1, R: 0, S: 1},
	{Q: -1, R: 1, S: 0},
}

// Creates a new Hex with the given q and r coordinates, computing s as -q-r.
func NewHexCoordinates(q, r int) HexCoordinates {
	return HexCoordinates{Q: q, R: r, S: -q - r}
}

// Calculates the hex distance between this hex and another.
func (h HexCoordinates) Distance(other HexCoordinates) int {
	vector := h.subtract(other)
	return (max(vector.Q, -vector.Q) + max(vector.R, -vector.R) + max(vector.S, -vector.S)) / 2
}

// Checks if this hex is equal to another hex.
func (h HexCoordinates) Equals(other HexCoordinates) bool {
	return h.Q == other.Q && h.R == other.R && h.S == other.S
}

// Returns the neighboring hex in the specified direction (0-5).
func (h HexCoordinates) Neighbor(direction int) HexCoordinates {
	return h.add(h.direction(direction))
}

func (h HexCoordinates) String() string {
	return fmt.Sprintf("{Q: %d, R: %d, S: %d}", h.Q, h.R, h.S)
}

// Adds another hex to this one, returning a new hex.
func (h HexCoordinates) add(other HexCoordinates) HexCoordinates {
	return HexCoordinates{
		Q: h.Q + other.Q,
		R: h.R + other.R,
		S: h.S + other.S,
	}
}

// Returns the hex direction vector for the given direction index (0-5).
func (h HexCoordinates) direction(direction int) HexCoordinates {
	if direction < 0 || direction > 5 {
		panic("Direction must be between 0 and 5")
	}
	return HexDirections[direction]
}

// Subtracts another hex from this one, returning a new hex.
func (h HexCoordinates) subtract(other HexCoordinates) HexCoordinates {
	return HexCoordinates{
		Q: h.Q - other.Q,
		R: h.R - other.R,
		S: h.S - other.S,
	}
}

// Represents the "general area" the hex is on.
// This is used to distinguish between special reserved areas,
// such as "Ganymede Colony" and "Phobos Space Haven" and the main area
// (Tharsis - which is coordinated and adjacency rules apply).
type HexArea int

const (
	HexAreaTharsis HexArea = iota
	HexAreaGanymedeColony
	HexAreaPhobosSpaceHaven
)

var hexAreaNames = map[HexArea]string{
	HexAreaTharsis:          "tharsis",
	HexAreaGanymedeColony:   "ganymede_colony",
	HexAreaPhobosSpaceHaven: "phobos_space_haven",
}

var hexAreaValues = map[string]HexArea{
	"tharsis":            HexAreaTharsis,
	"ganymede_colony":    HexAreaGanymedeColony,
	"phobos_space_haven": HexAreaPhobosSpaceHaven,
}

var hexAreaTileable = map[HexArea]bool{
	HexAreaTharsis:          true,
	HexAreaGanymedeColony:   false,
	HexAreaPhobosSpaceHaven: false,
}

func (ha HexArea) String() string {
	return hexAreaNames[ha]
}

// IsTileable returns whether this area uses a coordinate grid with adjacency rules.
func (ha HexArea) IsTileable() bool {
	return hexAreaTileable[ha]
}

// UnmarshalText implements encoding.TextUnmarshaler for YAML decoding.
func (ha *HexArea) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	val, ok := hexAreaValues[str]
	if !ok {
		return fmt.Errorf("unknown hex area: %q", str)
	}
	*ha = val
	return nil
}

// Represents the reservations of hexes on the board.
// A hex reservation determines exclusive placements for certain tiles,
// such as cities, oceans and volcanic reserved areas.
type HexReservation int

const (
	HexReservationDefault HexReservation = iota
	HexReservationOcean
	HexReservationVolcanic
	HexReservationCity
)

var hexReservationNames = map[HexReservation]string{
	HexReservationDefault:  "default",
	HexReservationOcean:    "ocean",
	HexReservationVolcanic: "volcanic",
	HexReservationCity:     "city",
}

var hexReservationValues = map[string]HexReservation{
	"":         HexReservationDefault,
	"default":  HexReservationDefault,
	"ocean":    HexReservationOcean,
	"volcanic": HexReservationVolcanic,
	"city":     HexReservationCity,
}

func (hr HexReservation) String() string {
	return hexReservationNames[hr]
}

// UnmarshalText implements encoding.TextUnmarshaler for YAML decoding.
func (hr *HexReservation) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	val, ok := hexReservationValues[str]
	if !ok {
		return fmt.Errorf("unknown hex reservation: %q", str)
	}
	*hr = val
	return nil
}

type HexMetadata struct {
	Area             HexArea        `yaml:"area"`
	Description      string         `yaml:"description"`
	PlacementBonuses ResourceSet    `yaml:"placement_bonuses"`
	Reservation      HexReservation `yaml:"reservation"`
}

func (hm HexMetadata) String() string {
	return fmt.Sprintf("{Area: %q, Reservation: %q, Description: %q, PlacementBonuses: %v}",
		hm.Area, hm.Reservation, hm.Description, hm.PlacementBonuses)
}

type Board struct {
	Name           string
	Description    string
	MainArea       HexArea                        // Primary board area (e.g., Tharsis)
	HexCoordinates map[HexCoordinates]HexMetadata // Map hexes (Tileable)
	HexAreas       map[HexArea]HexMetadata        // Special areas (Non-tileable)
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

	// Map Board name to MainArea (panics if unknown)
	val, ok := hexAreaValues[strings.ToLower(b.Name)]
	if !ok {
		panic(fmt.Sprintf("unknown board name: %q", b.Name))
	}
	b.MainArea = val

	b.HexCoordinates = make(map[HexCoordinates]HexMetadata)
	b.HexAreas = make(map[HexArea]HexMetadata)

	for _, entry := range temp.Hexes {
		if entry.Area.IsTileable() {
			coord := NewHexCoordinates(entry.Q, entry.R)
			b.HexCoordinates[coord] = entry.HexMetadata
		} else {
			b.HexAreas[entry.Area] = entry.HexMetadata
		}
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
			coord := NewHexCoordinates(q, r)
			if _, exists := board.HexCoordinates[coord]; !exists {
				board.HexCoordinates[coord] = HexMetadata{
					Area:        board.MainArea,
					Reservation: HexReservationDefault,
				}
			}
		}
	}
}
