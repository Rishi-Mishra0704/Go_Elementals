package models

type Skill interface {
	Use(player *Player, enemy *Enemy)
}

type PhysicalSkill struct {
	Name    string
	Damage  int8
	Stamina int8
}

type ElementalSkill struct {
	Name string
	Mana int8
}

type StatSkill struct {
	Name     string
	BuffHP   int8
	BuffMana int8
	BuffStam int8
}
