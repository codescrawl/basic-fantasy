package main

import (
	"fmt"
)

// Character ...
type Character struct {
	Stats *Stats
}

// NewCharacter ...
func NewCharacter() *Character {
	ch := Character{}
	ch.Stats = NewStats()

	return &ch
}

func (c *Character) String() string {
	return fmt.Sprintf("%s\n%v\n", c.Stats, validRaces(c))
}

func validRaces(c *Character) []string {
	validRaces := make([]string, 4)

	for _, race := range races {
		if race.meets(c.Stats) {
			validRaces = append(validRaces, race.name())
		}
	}

	return validRaces
}
