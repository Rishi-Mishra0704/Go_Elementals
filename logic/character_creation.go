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
	names := []string{"Fireball", "Water Blast", "Thunder Strike", "Earthquake", "Tornado", "Ice Shard", "Lightning Bolt", "Magma Surge"}
	name := names[rand.Intn(len(names))]

	var mana, damage int8
	switch name {
	case "Fireball":
		mana = 10
		damage = 14
	case "Water Blast":
		mana = 12
		damage = 13
	case "Thunder Strike":
		mana = 15
		damage = 28
	case "Earthquake":
		mana = 35
		damage = 30
	case "Tornado":
		mana = 25
		damage = 35
	case "Ice Shard":
		mana = 22
		damage = 18
	case "Lightning Bolt":
		mana = 28
		damage = 33
	case "Magma Surge":
		mana = 38
		damage = 40
	}

	return &models.ElementalSkill{Name: name, Mana: mana, Damage: damage}
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

// CreatePlayerCharacter creates a player character with randomized skill distribution and values.
func CreatePlayerCharacter() *models.Player {
	player := &models.Player{
		Character: &models.Character{HP: 120, Mana: 100, Stamina: 100},
		Skills:    make([]models.Skill, 0),
	}
	player.Skills = append(player.Skills, randomElementalSkill(), randomElementalSkill(), randomPhysicalSkill())
	return player
}

// CreateEnemyCharacter creates an enemy character with randomized skill distribution and values.
func CreateEnemyCharacter() *models.Enemy {
	enemy := &models.Enemy{
		Character: &models.Character{HP: 120, Mana: 100, Stamina: 100},
		Skills:    make([]models.Skill, 0),
	}
	enemy.Skills = append(enemy.Skills, randomElementalSkill())
	enemy.Skills = append(enemy.Skills, randomPhysicalSkill())
	return enemy
}
