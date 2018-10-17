package model

type Weapon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ApCost int `json:"ap_cost"`
	ApDamage int `json:"ap_damage"`
	Durability Var `json:"durability"`
}