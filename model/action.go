package model

type Action struct {
	Action string `json:"action"`
	Source string `json:"source"`
	Target string `json:"target"`
	Extra  string `json:"extra"`
}
