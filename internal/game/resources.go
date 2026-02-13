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
	ResourceAnimals:     "animals",
	ResourceCards:       "cards",
	ResourceEnergy:      "energy",
	ResourceHeat:        "heat",
	ResourceMegacredits: "megacredits",
	ResourceMicrobes:    "microbes",
	ResourcePlants:      "plants",
	ResourceSteel:       "steel",
	ResourceTitanium:    "titanium",
}

// fmt.Stringer for Resource.
func (r Resource) String() string {
	return resourceNames[r]
}
