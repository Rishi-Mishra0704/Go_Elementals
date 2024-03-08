package logic

import (
	"Go_Elemental/models"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Microsecond)
}

func PlayerTurn(player *models.Player, enemy *models.Enemy) {
	// Display player's stats
	fmt.Printf("Player's HP: %d, Mana: %d, Stamina: %d\n", player.Character.HP, player.Character.Mana, player.Character.Stamina)
	// Display available skills
	fmt.Println("Player's available skills:")
	for i, skill := range player.Skills {
		switch s := skill.(type) {
		case *models.ElementalSkill:
			fmt.Printf("%d. %s\n", i+1, s.Name)
		case *models.PhysicalSkill:
			fmt.Printf("%d. %s\n", i+1, s.Name)
		case *models.StatSkill:
			if s.Name != "Regenerate" || player.Character.HP < 50 {
				fmt.Printf("%d. %s\n", i+1, s.Name)
			}
		}
	}
	// Let the player choose a skill
	var choice int
	fmt.Print("Choose a skill: ")
	fmt.Scanln(&choice)
	// Use the selected skill on the enemy
	if player.Character.Mana >= 0 && player.Character.Stamina >= 0 {
		player.Skills[choice-1].Use(player, enemy)
	} else {
		fmt.Println("Not enough mana or stamina to use this skill.")
	}
	// Check if the enemy's health is zero after the attack
	if enemy.Character.HP <= 0 {
		fmt.Println("Enemy defeated! You win!")
		return
	}
}

// Enemy's turn
func EnemyTurn(player *models.Player, enemy *models.Enemy) {
	// Select a random skill for the enemy
	enemySkill := enemy.Skills[rand.Intn(len(enemy.Skills))]
	// Check if the enemy should use the skill
	if enemy.Character.HP < 50 || enemySkill.(*models.StatSkill).Name != "Regenerate" {
		// Use the selected skill on the enemy
		if enemy.Character.Mana >= 0 && enemy.Character.Stamina >= 0 {
			enemySkill.Use(nil, enemy) // Using on self, passing nil for player
		} else {
			fmt.Println("Enemy does not have enough mana or stamina to use this skill.")
		}
	}
	// Check if the player's health is zero after the attack
	if player.Character.HP <= 0 {
		fmt.Println("Player defeated! You lose!")
		return
	}
}
