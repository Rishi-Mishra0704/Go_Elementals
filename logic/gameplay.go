package logic

import (
	"Go_Elemental/models"
	"fmt"
	"math/rand"
)

func PlayerTurn(player *models.Player, enemy *models.Enemy) {
	// Display player's stats
	fmt.Printf("Player's HP: %d, Mana: %d, Stamina: %d\n", player.HP, player.Mana, player.Stamina)
	// Display available skills
	fmt.Println("Player's available skills:")
	for i, skill := range player.Skills {
		switch s := skill.(type) {
		case *models.ElementalSkill:
			fmt.Printf("%d. %s\n", i+1, s.Name)
		case *models.PhysicalSkill:
			fmt.Printf("%d. %s\n", i+1, s.Name)
		case *models.StatSkill:
			fmt.Printf("%d. %s\n", i+1, s.Name)
		}
	}
	// Let the player choose a skill
	var choice int
	fmt.Print("Choose a skill: ")
	fmt.Scanln(&choice)
	// Use the selected skill on the enemy
	player.Skills[choice-1].Use(player, enemy)
	// Check if the enemy's health is zero after the attack
	if enemy.HP <= 0 {
		fmt.Println("Enemy defeated! You win!")
		return
	}
}

// Enemy's turn
func EnemyTurn(player *models.Player, enemy *models.Enemy) {
	// Select a random skill for the enemy
	fmt.Printf("Enemy's HP: %d, Mana: %d, Stamina: %d\n", enemy.HP, enemy.Mana, enemy.Stamina)
	enemySkill := enemy.Skills[rand.Intn(len(enemy.Skills))]
	// Use the selected skill on the player
	enemySkill.Use(player, nil)
	// Check if the player's health is zero after the attack
	if player.HP <= 0 {
		fmt.Println("Player defeated! You lose!")
		return
	}
}
