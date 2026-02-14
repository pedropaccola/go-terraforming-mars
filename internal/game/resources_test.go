package game

import (
	"testing"
)

func TestResource(t *testing.T) {
	t.Run("Resource String", func(t *testing.T) {
		tests := []struct {
			resource Resource
			expected string
		}{
			{ResourceAnimals, "Animals"},
			{ResourceCards, "Cards"},
			{ResourceEnergy, "Energy"},
			{ResourceHeat, "Heat"},
			{ResourceMegacredits, "Megacredits"},
			{ResourceMicrobes, "Microbes"},
			{ResourcePlants, "Plants"},
			{ResourceSteel, "Steel"},
			{ResourceTitanium, "Titanium"},
		}

		for _, tc := range tests {
			if tc.resource.String() != tc.expected {
				t.Errorf("Resource.String(): expected %v, got %v", tc.expected, tc.resource.String())
			}
		}
	})
}
