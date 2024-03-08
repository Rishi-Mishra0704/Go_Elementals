package models

import "fmt"

type Skill interface {
	Use(player *Player, enemy *Enemy)
}

type PhysicalSkill struct {
	Name    string
	Damage  int8
	Stamina int8
}

func (s *PhysicalSkill) Use(player *Player, enemy *Enemy) {
	fmt.Printf("Using %s (Damage: %d)\n", s.Name, s.Damage)
	if enemy != nil {
		enemy.HP -= s.Damage
	}
	if player != nil {
		player.Stamina -= s.Stamina
	}
}

type ElementalSkill struct {
	Name   string
	Mana   int8
	Damage int8
}

func (s *ElementalSkill) Use(player *Player, enemy *Enemy) {
	fmt.Printf("Using %s (Mana: %d)\n", s.Name, s.Mana)
	if player != nil {
		player.Mana -= s.Mana
	}
	if enemy != nil {
		enemy.HP -= s.Damage
	}
}

type StatSkill struct {
	Name     string
	BuffHP   int8
	BuffMana int8
	BuffStam int8
}

func (s *StatSkill) Use(player *Player, enemy *Enemy) {
	fmt.Printf("Using %s (BuffHP: %d, BuffMana: %d, BuffStam: %d)\n", s.Name, s.BuffHP, s.BuffMana, s.BuffStam)
	if player != nil {
		player.HP += s.BuffHP
		player.Mana += s.BuffMana
		player.Stamina += s.BuffStam
	}
	if enemy != nil {
		enemy.HP += s.BuffHP
		enemy.Mana += s.BuffMana
		enemy.Stamina += s.BuffStam
	}
}
