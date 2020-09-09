package main

import "fmt"

// Stats ...
type Stats struct {
	Values map[string]Stat
}

// String ...
func (s *Stats) String() string {
	str := ""
	for _, stat := range s.Values {
		str += fmt.Sprintf("%s", &stat)
		str += " "
	}
	return str
}

// Strength ...
func (s *Stats) Strength() int {
	return s.Values["strength"].Value
}

// Intelligence ...
func (s *Stats) Intelligence() int {
	return s.Values["intelligence"].Value
}

// Wisdom ...
func (s *Stats) Wisdom() int {
	return s.Values["wisdom"].Value
}

// Dexterity ...
func (s *Stats) Dexterity() int {
	return s.Values["dexterity"].Value
}

// Constitution ...
func (s *Stats) Constitution() int {
	return s.Values["constitution"].Value
}

// Charisma ...
func (s *Stats) Charisma() int {
	return s.Values["charisma"].Value
}

// NewStats ...
func NewStats() *Stats {
	s := Stats{}
	s.Values = make(map[string]Stat, 6)

	s.Values["strength"] = *NewStat()
	s.Values["intelligence"] = *NewStat()
	s.Values["wisdom"] = *NewStat()
	s.Values["dexterity"] = *NewStat()
	s.Values["constitution"] = *NewStat()
	s.Values["charisma"] = *NewStat()

	return &s
}
