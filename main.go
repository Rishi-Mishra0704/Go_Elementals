package main

import (
	"Go_Elemental/logic"
	"fmt"
)

func main() {
	// Create player and enemy characters
	player := logic.CreatePlayerCharacter()
	enemy := logic.CreateEnemyCharacter()

	// Main game loop
	for {
		// Player's turn
		fmt.Println("Player's turn:")
		logic.PlayerTurn(player, enemy)

		// Check if the game is over
		if player.Character.HP <= 0 || enemy.Character.HP <= 0 {
			break
		}

		// Enemy's turn
		fmt.Println("Enemy's turn:")
		logic.EnemyTurn(player, enemy)

		// Check if the game is over
		if player.Character.HP <= 0 || enemy.Character.HP <= 0 {
			break
		}
	}

	// Determine the outcome of the game
	if player.Character.HP <= 0 {
		fmt.Println("Player defeated! You lose!")
	} else if enemy.Character.HP <= 0 {
		fmt.Println("Enemy defeated! You win!")
	}
}
