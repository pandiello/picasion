package model

type State struct {
	Turns Var `json:"turns"`
	Allies []Warrior `json:"allies"`
	Enemies []Warrior `json:"enemies"`
	Weapons []Weapon `json:"weapons"`
	Items []Item `json:"items"`
}