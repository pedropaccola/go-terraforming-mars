package game

import "fmt"

// ResourceType represents the different types of resources in the game.
type ResourceType int

const (
	ResourceAnimals ResourceType = iota
	ResourceCards
	ResourceEnergy
	ResourceHeat
	ResourceMegaCredits
	ResourceMicrobes
	ResourcePlants
	ResourceSteel
	ResourceTitanium
)

// Collection of resources for resource management.
type ResourceSet struct {
	Animals     int `yaml:"animals"`
	Cards       int `yaml:"cards"`
	Energy      int `yaml:"energy"`
	Heat        int `yaml:"heat"`
	MegaCredits int `yaml:"megacredits"`
	Microbes    int `yaml:"microbes"`
	Plants      int `yaml:"plants"`
	Steel       int `yaml:"steel"`
	Titanium    int `yaml:"titanium"`
}

// Adds the specified amount to the given resource type, returning a new ResourceSet.
// It does not modify the original set. Panics if rtype is unknown.
func (rs ResourceSet) AddResource(rtype ResourceType, amount int) ResourceSet {
	result := rs
	switch rtype {
	case ResourceAnimals:
		result.Animals += amount
	case ResourceCards:
		result.Cards += amount
	case ResourceEnergy:
		result.Energy += amount
	case ResourceHeat:
		result.Heat += amount
	case ResourceMegaCredits:
		result.MegaCredits += amount
	case ResourceMicrobes:
		result.Microbes += amount
	case ResourcePlants:
		result.Plants += amount
	case ResourceSteel:
		result.Steel += amount
	case ResourceTitanium:
		result.Titanium += amount
	default:
		panic("unknown resource type")
	}

	return result
}

// Combines this ResourceSet with another, returning a new ResourceSet with summed quantities.
// It does not modify the original set.
func (rs ResourceSet) AddSet(other ResourceSet) ResourceSet {
	result := ResourceSet{
		Animals:     rs.Animals + other.Animals,
		Cards:       rs.Cards + other.Cards,
		Energy:      rs.Energy + other.Energy,
		Heat:        rs.Heat + other.Heat,
		MegaCredits: rs.MegaCredits + other.MegaCredits,
		Microbes:    rs.Microbes + other.Microbes,
		Plants:      rs.Plants + other.Plants,
		Steel:       rs.Steel + other.Steel,
		Titanium:    rs.Titanium + other.Titanium,
	}
	return result
}

// Get returns the quantity of the specified resource type.
// Returns 0 if rtype is unknown.
func (rs ResourceSet) Get(rtype ResourceType) int {
	switch rtype {
	case ResourceAnimals:
		return rs.Animals
	case ResourceCards:
		return rs.Cards
	case ResourceEnergy:
		return rs.Energy
	case ResourceHeat:
		return rs.Heat
	case ResourceMegaCredits:
		return rs.MegaCredits
	case ResourceMicrobes:
		return rs.Microbes
	case ResourcePlants:
		return rs.Plants
	case ResourceSteel:
		return rs.Steel
	case ResourceTitanium:
		return rs.Titanium
	default:
		return 0
	}
}

func (rs ResourceSet) String() string {
	return fmt.Sprintf("{Animals: %d, Cards: %d, Energy: %d, Heat: %d, MegaCredits: %d, Microbes: %d, Plants: %d, Steel: %d, Titanium: %d}",
		rs.Animals, rs.Cards, rs.Energy, rs.Heat, rs.MegaCredits, rs.Microbes, rs.Plants, rs.Steel, rs.Titanium)
}
