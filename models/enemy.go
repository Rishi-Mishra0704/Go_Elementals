package models

type Enemy struct {
	*Character
	Skills []Skill
}
