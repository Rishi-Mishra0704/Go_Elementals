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
