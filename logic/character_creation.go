package logic

import (
	"Go_Elemental/models"
	"math/rand"
	"time"
)

func init() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)
}

func randomElementalSkill() models.Skill {
	names := []string{"Fireball", "Water Blast", "Thunder Strike", "Earthquake", "Tornado"}
	name := names[rand.Intn(len(names))]
	switch name {
	case "Fireball":
		return &models.ElementalSkill{Name: name, Mana: 20, Damage: 20}
	case "Water Blast":
		return &models.ElementalSkill{Name: name, Mana: 25, Damage: 25}
	case "Thunder Strike":
		return &models.ElementalSkill{Name: name, Mana: 30, Damage: 35}
	case "Earthquake":
		return &models.ElementalSkill{Name: name, Mana: 35, Damage: 38}
	case "Tornado":
		return &models.ElementalSkill{Name: name, Mana: 40, Damage: 30}
	}
	return nil
}

func randomPhysicalSkill() models.Skill {
	names := []string{"Slash", "Thrust", "Punch", "Kick", "Tackle"}
	name := names[rand.Intn(len(names))]
	switch name {
	case "Slash":
		return &models.PhysicalSkill{Name: name, Damage: 15, Stamina: 10}
	case "Thrust":
		return &models.PhysicalSkill{Name: name, Damage: 18, Stamina: 12}
	case "Punch":
		return &models.PhysicalSkill{Name: name, Damage: 12, Stamina: 8}
	case "Kick":
		return &models.PhysicalSkill{Name: name, Damage: 20, Stamina: 15}
	case "Tackle":
		return &models.PhysicalSkill{Name: name, Damage: 10, Stamina: 5}
	}
	return nil
}

func randomStatSkill() models.Skill {
	names := []string{"Boost", "Nerf", "Regenerate", "Exhaust", "Reflect"}
	name := names[rand.Intn(len(names))]
	switch name {
	case "Boost":
		return &models.StatSkill{Name: name, BuffHP: 10, BuffMana: 15, BuffStam: 15}
	case "Nerf":
		return &models.StatSkill{Name: name, BuffHP: -5, BuffMana: -5, BuffStam: -5}
	case "Regenerate":
		return &models.StatSkill{Name: name, BuffHP: 25, BuffMana: 0, BuffStam: 0}
	case "Exhaust":
		return &models.StatSkill{Name: name, BuffHP: -10, BuffMana: -10, BuffStam: -10}
	}
	return nil
}

// CreatePlayerCharacter creates a player character with randomized skill distribution and values.
func CreatePlayerCharacter() *models.Player {
	player := &models.Player{
		Character: &models.Character{HP: 50, Mana: 100, Stamina: 100},
		Skills:    make([]models.Skill, 0),
	}
	player.Skills = append(player.Skills, randomElementalSkill(), randomElementalSkill(), randomPhysicalSkill(), randomStatSkill())
	return player
}

// CreateEnemyCharacter creates an enemy character with randomized skill distribution and values.
func CreateEnemyCharacter() *models.Enemy {
	enemy := &models.Enemy{
		Character: &models.Character{HP: 50, Mana: 100, Stamina: 100},
		Skills:    make([]models.Skill, 0),
	}
	enemy.Skills = append(enemy.Skills, randomElementalSkill())
	enemy.Skills = append(enemy.Skills, randomPhysicalSkill(), randomStatSkill())
	return enemy
}
