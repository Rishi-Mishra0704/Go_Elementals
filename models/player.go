package models

type Player struct {
	*Character
	Skills []Skill
}
