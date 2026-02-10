package game

import (
	"fmt"
	"strings"
)

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

var resourceValues = map[string]Resource{
	"animals":     ResourceAnimals,
	"cards":       ResourceCards,
	"energy":      ResourceEnergy,
	"heat":        ResourceHeat,
	"megacredits": ResourceMegacredits,
	"microbes":    ResourceMicrobes,
	"plants":      ResourcePlants,
	"steel":       ResourceSteel,
	"titanium":    ResourceTitanium,
}

// fmt.Stringer for Resource.
func (r Resource) String() string {
	return resourceNames[r]
}

// encoding.TextUnmarshaler for Resource
func (r *Resource) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	val, ok := resourceValues[str]
	if !ok {
		return fmt.Errorf("unknown resource: %q", str)
	}
	*r = val
	return nil
}
