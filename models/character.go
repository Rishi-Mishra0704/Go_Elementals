package models

type Character struct {
	HP      int8
	Mana    int8
	Stamina int8
}

type Skill interface {
	Use(c *Character)
}
