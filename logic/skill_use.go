package logic

import (
	"Go_Elemental/models"
	"fmt"
)

func UsePhysicalSkill(s *models.PhysicalSkill, c *models.Character) {
	fmt.Printf("Using %s (Stamina: %d)\n", s.Name, s.Stamina)
	c.Stamina -= s.Stamina
}

func UseElementalSkill(s *models.ElementalSkill, c *models.Character) {
	fmt.Printf("Using %s (Mana: %d)\n", s.Name, s.Mana)
	c.Mana -= s.Mana
}

func UseStatSkill(s *models.StatSkill, player *models.Player, enemy *models.Enemy) {
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
