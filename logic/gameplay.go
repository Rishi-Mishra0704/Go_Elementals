package logic

import (
	"Go_Elemental/models"
	"fmt"
	"math/rand"
)

func PlayerTurn(player *models.Player, enemy *models.Enemy) {
	// Display player's stats
	fmt.Printf("Player's HP: %d, Mana: %d, Stamina: %d\n", player.Character.HP, player.Character.Mana, player.Character.Stamina)
	// Display available skills
	fmt.Println("Player's available skills:")
	for i, skill := range player.Skills {
		switch s := skill.(type) {
		case *models.ElementalSkill:
			fmt.Printf("%d. %s (Mana: %d, Damage: %d)\n", i+1, s.Name, s.Mana, s.Damage)
		case *models.PhysicalSkill:
			fmt.Printf("%d. %s (Damage: %d, Stamina: %d)\n", i+1, s.Name, s.Damage, s.Stamina)
		}
	}
	// Let the player choose a skill
	var choice int
	fmt.Print("Choose a skill: ")
	fmt.Scanln(&choice)
	// Use the selected skill on the enemy
	selectedSkill := player.Skills[choice-1]

	selectedSkill.Use(player.Character, enemy.Character)

	fmt.Println("________________________________________________________________________________")
}

// Enemy's turn
func EnemyTurn(player *models.Player, enemy *models.Enemy) {
	if player.Character.HP <= 0 {

		fmt.Println("Player is defeated! You lose!")
		return
	}
	if enemy.Character.HP <= 0 {
		fmt.Println("Enemy is defeated! You win!")
		return
	}

	// Select a random skill for the enemy
	fmt.Printf("Enemy's HP: %d, Mana: %d, Stamina: %d\n", enemy.Character.HP, enemy.Character.Mana, enemy.Character.Stamina)
	enemySkill := enemy.Skills[rand.Intn(len(enemy.Skills))]
	// Use the selected skill on the player
	selectedSkill := enemySkill
	selectedSkill.Use(enemy.Character, player.Character)
	fmt.Println("________________________________________________________________________________")
}
