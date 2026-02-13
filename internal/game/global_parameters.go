package game

import "fmt"

// GlobalParameter identifies the primary terraforming scales of Mars.
type GlobalParameter int

const (
	GlobalParameterTemperature GlobalParameter = iota
	GlobalParameterOxygen
	GlobalParameterOcean
)

var globalParameterNames = map[GlobalParameter]string{
	GlobalParameterTemperature: "temperature",
	GlobalParameterOxygen:      "oxygen",
	GlobalParameterOcean:       "ocean",
}

// fmt.Stringer for GlobalParameter.
func (gp GlobalParameter) String() string {
	return globalParameterNames[gp]
}

// GlobalParameterTrack manages the current state and boundaries of a terraforming scale.
type GlobalParameterTrack struct {
	Current, Max, Min, Step int
}

// Advance attempts to increase the parameter by the specified number of steps.
// It returns (success, actualSteps).
// Success is false only if the track was already at Max.
// actualSteps counts how many increments actually occurred (useful for awarding TR).
func (gpt *GlobalParameterTrack) Advance(steps int) (bool, int) {
	if gpt.IsMaxed() {
		return false, 0
	}

	for i := range steps {
		gpt.Current += gpt.Step

		if gpt.IsMaxed() {
			return true, i + 1
		}
	}

	return true, steps
}

// IsMaxed returns true if the global parameter has reached or exceeded its maximum value.
func (gpt *GlobalParameterTrack) IsMaxed() bool {
	return gpt.Current >= gpt.Max
}

// fmt.Stringer for GlobalParameterTrack.
func (gpt *GlobalParameterTrack) String() string {
	return fmt.Sprintf("%d/%d", gpt.Current, gpt.Max)
}

// NewGlobalParameters initializes a standard set of terraforming tracks for Mars.
func NewGlobalParameters() map[GlobalParameter]*GlobalParameterTrack {
	return map[GlobalParameter]*GlobalParameterTrack{
		GlobalParameterTemperature: {Current: -30, Min: -30, Max: 8, Step: 2},
		GlobalParameterOxygen:      {Current: 0, Min: 0, Max: 14, Step: 1},
		GlobalParameterOcean:       {Current: 0, Min: 0, Max: 9, Step: 1},
	}
}
