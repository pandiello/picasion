package model

type Warrior struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Health Var `json:"health"`
	Ap Var `json:"ap"`
	Weapons []Weapon `json:"weapons"`
	Items []Item `json:"items"`
}