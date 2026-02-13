package game

// The radius of the hexagonal game board.
const MapRadius = 4

// The name of a board being played.
type BoardName int

const (
	BoardNameTharsis BoardName = iota
)

var boardNames = map[BoardName]string{
	BoardNameTharsis: "Tharsis",
}

// fmt.Stringer for BoardName.
func (bn BoardName) String() string {
	return boardNames[bn]
}

// Main Board struct, containing all information about the board's hexes.
type Board struct {
	Name        BoardName
	Description string
	Hexes       map[Hex]HexMetadata
}

// Loads the hexes and creates a new game board.
func NewBoard(boardName BoardName) *Board {
	board := Board{
		Hexes: make(map[Hex]HexMetadata),
	}

	switch boardName {
	default:
		board.initTharsis()
	}

	board.initEmptyHexes()

	return &board
}

// Adds a hex to the board if it doesn't already exist.
func (b *Board) addHex(q, r int, meta HexMetadata) {
	hex := NewHex(q, r)
	if _, exists := b.Hexes[hex]; exists {
		return
	}
	b.Hexes[hex] = meta
}

// Initializes the board with empty hexes.
func (b *Board) initEmptyHexes() {
	for q := -MapRadius; q <= MapRadius; q++ {
		r1 := max(-MapRadius, -q-MapRadius)
		r2 := min(MapRadius, -q+MapRadius)
		for r := r1; r <= r2; r++ {
			b.addHex(q, r, HexMetadata{})
		}
	}
}

// Initialize the board with Tharsis-specific hexes
func (b *Board) initTharsis() {
	b.Name = BoardNameTharsis
	b.Description = `The game board has an accurate map of the Tharsis
		region of Mars, including Valles Marineris and 3 of the 4 great
		volcanos. Only the region around Olympus Mons is missing.
		The areas reserved for Ocean tiles are low in elevation so water
		will naturally flow there. The plant bonuses around the equator
		simulate that the higher average temperature will make it easier
		for life to thrive there. Mountain ridges have steel and titanium
		bonuses, while other interesting sites may have a card draw
		bonus, like the Viking site where the first man-made lander
		touched down.`
	b.addHex(0, -4, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel, ResourceSteel},
	})
	b.addHex(1, -4, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel, ResourceSteel},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(3, -4, HexMetadata{
		PlacementBonuses: []Resource{ResourceCards},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(4, -4, HexMetadata{
		PlacementRule: PlacementRuleOcean,
	})
	b.addHex(0, -3, HexMetadata{
		Description:      "Tharsis Tholus",
		PlacementBonuses: []Resource{ResourceSteel},
		PlacementRule:    PlacementRuleVolcanic,
	})
	b.addHex(4, -3, HexMetadata{
		PlacementBonuses: []Resource{ResourceCards, ResourceCards},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(-2, -2, HexMetadata{
		Description:      "Ascraeus Mons",
		PlacementBonuses: []Resource{ResourceCards},
		PlacementRule:    PlacementRuleVolcanic,
	})
	b.addHex(4, -2, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel},
	})
	b.addHex(-3, -1, HexMetadata{
		Description:      "Pavonis Mons",
		PlacementBonuses: []Resource{ResourcePlants, ResourceTitanium},
		PlacementRule:    PlacementRuleVolcanic,
	})
	b.addHex(-2, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(-1, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(0, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(1, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(2, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(3, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(4, -1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(-4, 0, HexMetadata{
		Description:      "Arsia Mons",
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleVolcanic,
	})
	b.addHex(-3, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(-2, 0, HexMetadata{
		Description:      "Noctis City",
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleCity,
	})
	b.addHex(-1, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(0, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(1, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(2, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(3, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(4, 0, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(-4, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(-3, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants, ResourcePlants},
	})
	b.addHex(-2, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(-1, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(0, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(1, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(2, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(3, 1, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
		PlacementRule:    PlacementRuleOcean,
	})
	b.addHex(1, 2, HexMetadata{
		PlacementBonuses: []Resource{ResourcePlants},
	})
	b.addHex(-4, 3, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel, ResourceSteel},
	})
	b.addHex(-2, 3, HexMetadata{
		PlacementBonuses: []Resource{ResourceCards},
	})
	b.addHex(-1, 3, HexMetadata{
		PlacementBonuses: []Resource{ResourceCards},
	})
	b.addHex(1, 3, HexMetadata{
		PlacementBonuses: []Resource{ResourceTitanium},
	})
	b.addHex(-4, 4, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel},
	})
	b.addHex(-3, 4, HexMetadata{
		PlacementBonuses: []Resource{ResourceSteel, ResourceSteel},
	})
	b.addHex(0, 4, HexMetadata{
		PlacementBonuses: []Resource{ResourceTitanium, ResourceTitanium},
		PlacementRule:    PlacementRuleOcean,
	})
}
