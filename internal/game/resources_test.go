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

	t.Run("Unmarshal text", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    Resource
			expectedErr bool
		}{
			{"animals", "animals", ResourceAnimals, false},
			{"cards", "cards", ResourceCards, false},
			{"energy uppercase", "ENERGY", ResourceEnergy, false},
			{"heat mixed case", "HeAt", ResourceHeat, false},
			{"megacredits", "megacredits", ResourceMegacredits, false},
			{"invalid resource", "gold", 0, true},
		}

		for _, tc := range tests {
			var r Resource
			err := r.UnmarshalText([]byte(tc.input))
			if (err != nil) != tc.expectedErr {
				t.Errorf("Resource.UnmarshalText() expectedErr = %v, gotErr %v", tc.expectedErr, err)
				continue
			}
			if !tc.expectedErr && r != tc.expected {
				t.Errorf("Resource.UnmarshalText() expected = %v, got %v", r, tc.expected)
			}
		}
	})
}
