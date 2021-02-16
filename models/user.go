package models

type User struct {
	Name     string `json:"name"`
	IngameID string `json:"ingameID"`
	Image    string `json:"image"`
}
