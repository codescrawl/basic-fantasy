package main

import (
	"fmt"
	"math/rand"
)

// Stat ...
type Stat struct {
	Value int
}

// NewStat ...
func NewStat() *Stat {
	s := Stat{}
	s.Value = roll3d6()
	return &s
}

// Modifier ...
func (s *Stat) Modifier() int {
	switch s.Value {
	case 3:
		return -3
	case 4, 5:
		return -2
	case 6, 7, 8:
		return -1
	case 9, 10, 11, 12:
		return 0
	case 13, 14, 15:
		return 1
	case 16, 17:
		return 2
	case 18:
		return 3
	}
	return int(^uint(0) >> 1)
}

func (s *Stat) String() string {
	return fmt.Sprintf("%d (%+d)", s.Value, s.Modifier())
}

func roll3d6() int {
	return rand.Intn(6) + rand.Intn(6) + rand.Intn(6) + 3
}
