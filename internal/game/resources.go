package game

// Resource represents the different resources in the game.
type Resource int

const (
	ResourceAnimals Resource = iota
	ResourceCards
	ResourceEnergy
	ResourceHeat
	ResourceMegacredits
	ResourceMicrobes
	ResourcePlants
	ResourceSteel
	ResourceTitanium
)

var resourceNames = map[Resource]string{
	ResourceAnimals:     "Animals",
	ResourceCards:       "Cards",
	ResourceEnergy:      "Energy",
	ResourceHeat:        "Heat",
	ResourceMegacredits: "Megacredits",
	ResourceMicrobes:    "Microbes",
	ResourcePlants:      "Plants",
	ResourceSteel:       "Steel",
	ResourceTitanium:    "Titanium",
}

// fmt.Stringer for Resource.
func (r Resource) String() string {
	return resourceNames[r]
}
