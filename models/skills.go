package models

type Skill interface {
	Use(player *Player, enemy *Enemy)
}

type PhysicalSkill struct {
	Name    string
	Stamina int
}

type ElementalSkill struct {
	Name string
	Mana int
}

type StatSkill struct {
	Name     string
	BuffHP   int
	BuffMana int
	BuffStam int
}
