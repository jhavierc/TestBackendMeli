package model

type Content struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	State       string `json:"state"`
	Type        int    `json:"type"`
}
