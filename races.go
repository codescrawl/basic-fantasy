package main

// Race ...
type Race interface {
	meets(s *Stats) bool
	name() string
}

// Dwarf ...
type Dwarf struct {
}

func (d *Dwarf) meets(s *Stats) bool {
	return s.Constitution() >= 9 && s.Charisma() <= 17
}

func (d *Dwarf) name() string {
	return "Dwarf"
}

// Elf ...
type Elf struct {
}

func (e *Elf) meets(s *Stats) bool {
	return s.Intelligence() >= 9 && s.Constitution() <= 17
}

// Name ...
func (e *Elf) name() string {
	return "Elf"
}

// Halfling ...
type Halfling struct {
}

func (h *Halfling) meets(s *Stats) bool {
	return s.Dexterity() >= 9 && s.Strength() <= 17
}

// Name ...
func (h *Halfling) name() string {
	return "Halfling"
}

// Human ...
type Human struct {
}

func (h *Human) meets(s *Stats) bool {
	return true
}

func (h *Human) name() string {
	return "Human"
}

var races = []Race{
	&Dwarf{},
	&Elf{},
	&Halfling{},
	&Human{},
}
