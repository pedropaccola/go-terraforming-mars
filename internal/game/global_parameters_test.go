package game

import "testing"

func TestGlobalParameter(t *testing.T) {
	t.Run("GlobalParameter String", func(t *testing.T) {
		tests := []struct {
			param    GlobalParameter
			expected string
		}{
			{GlobalParameterTemperature, "temperature"},
			{GlobalParameterOxygen, "oxygen"},
			{GlobalParameterOcean, "ocean"},
		}

		for _, tc := range tests {
			if tc.param.String() != tc.expected {
				t.Errorf("GlobalParameter.String(): expected %q, got %q", tc.expected, tc.param.String())
			}
		}
	})
}

func TestGlobalParameterTrack(t *testing.T) {
	t.Run("NewGlobalParameters initializes correct tracks", func(t *testing.T) {
		tracks := NewGlobalParameters()

		expected := []struct {
			param GlobalParameter
			min   int
			max   int
			step  int
		}{
			{GlobalParameterTemperature, -30, 8, 2},
			{GlobalParameterOxygen, 0, 14, 1},
			{GlobalParameterOcean, 0, 9, 1},
		}

		if len(tracks) != len(expected) {
			t.Errorf("Expected %d tracks, got %d", len(expected), len(tracks))
		}

		for _, e := range expected {
			track, ok := tracks[e.param]
			if !ok {
				t.Errorf("Missing track for %v", e.param)
				continue
			}
			if track.Min != e.min || track.Max != e.max || track.Step != e.step {
				t.Errorf("%v: expected min/max/step %d/%d/%d, got %d/%d/%d",
					e.param, e.min, e.max, e.step, track.Min, track.Max, track.Step)
			}
			if track.Current != e.min {
				t.Errorf("%v: expected current at min (%d), got %d", e.param, e.min, track.Current)
			}
		}
	})

	t.Run("Advance increases current value", func(t *testing.T) {
		track := &GlobalParameterTrack{Current: 0, Max: 10, Min: 0, Step: 1}
		success, actualSteps := track.Advance(2)

		if !success {
			t.Error("Expected Advance to succeed")
		}
		if actualSteps != 2 {
			t.Errorf("Expected 2 actual steps, got %d", actualSteps)
		}
		if track.Current != 2 {
			t.Errorf("Expected current 2, got %d", track.Current)
		}
	})

	t.Run("Advance caps at max", func(t *testing.T) {
		track := &GlobalParameterTrack{Current: 8, Max: 10, Min: 0, Step: 1}
		success, actualSteps := track.Advance(5)

		if !success {
			t.Error("Expected Advance to succeed as it wasn't maxed")
		}
		if actualSteps != 2 {
			t.Errorf("Expected 2 actual steps (reaching max), got %d", actualSteps)
		}
		if track.Current != 10 {
			t.Errorf("Expected current 10, got %d", track.Current)
		}
	})

	t.Run("Advance returns false if already at max", func(t *testing.T) {
		track := &GlobalParameterTrack{Current: 10, Max: 10, Min: 0, Step: 1}
		success, actualSteps := track.Advance(1)

		if success {
			t.Error("Expected Advance to fail when already at max")
		}
		if actualSteps != 0 {
			t.Errorf("Expected 0 actual steps, got %d", actualSteps)
		}
	})

	t.Run("IsMaxed returns true at or above max", func(t *testing.T) {
		track := &GlobalParameterTrack{Current: 9, Max: 10}
		if track.IsMaxed() {
			t.Error("Expected 9 to not be maxed (max 10)")
		}

		track.Current = 10
		if !track.IsMaxed() {
			t.Error("Expected 10 to be maxed (max 10)")
		}

		track.Current = 11
		if !track.IsMaxed() {
			t.Error("Expected 11 to be maxed (max 10)")
		}
	})

	t.Run("String format", func(t *testing.T) {
		track := &GlobalParameterTrack{Current: 5, Max: 10}
		expected := "5/10"
		if track.String() != expected {
			t.Errorf("Expected %q, got %q", expected, track.String())
		}
	})
}
