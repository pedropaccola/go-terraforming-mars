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
			{ResourceAnimals, "animals"},
			{ResourceCards, "cards"},
			{ResourceEnergy, "energy"},
			{ResourceHeat, "heat"},
			{ResourceMegacredits, "megacredits"},
			{ResourceMicrobes, "microbes"},
			{ResourcePlants, "plants"},
			{ResourceSteel, "steel"},
			{ResourceTitanium, "titanium"},
		}

		for _, tc := range tests {
			if tc.resource.String() != tc.expected {
				t.Errorf("Resource.String(): expected %v, got %v", tc.expected, tc.resource.String())
			}
		}
	})
}
