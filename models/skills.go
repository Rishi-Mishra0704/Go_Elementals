package models

import "fmt"

type Skill interface {
	Use(caster, target *Character)
}

type PhysicalSkill struct {
	Name    string
	Damage  int8
	Stamina int8
}

func (s *PhysicalSkill) Use(caster, target *Character) {
	fmt.Printf("Using %s (Damage: %d)\n", s.Name, s.Damage)
	if target != nil {
		target.HP -= s.Damage
	}
	if caster != nil {
		caster.Stamina -= s.Stamina
	}
}

type ElementalSkill struct {
	Name   string
	Mana   int8
	Damage int8
}

func (s *ElementalSkill) Use(caster, target *Character) {
	fmt.Printf("Using %s (Mana: %d)\n", s.Name, s.Mana)
	if caster != nil && caster.Mana >= s.Mana {
		caster.Mana -= s.Mana
		if target != nil {
			target.HP -= s.Damage
		}
	} else {
		fmt.Println("Not enough mana to use this skill!")
	}
}

type StatSkill struct {
	Name     string
	BuffHP   int8
	BuffMana int8
	BuffStam int8
}

func (s *StatSkill) Use(caster, target *Character) {
	fmt.Printf("Using %s (BuffHP: %d, BuffMana: %d, BuffStamina: %d)\n", s.Name, s.BuffHP, s.BuffMana, s.BuffStam)
	if target != nil {
		target.HP += s.BuffHP
		target.Mana += s.BuffMana
		target.Stamina += s.BuffStam
	}
	if caster != nil {
		caster.Mana -= s.BuffMana
		caster.Stamina -= s.BuffStam
	}
}
