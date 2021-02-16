package models

type Crew struct {
	Name              string  `json:"name"`
	Image             *string `json:"image"`
	GroupID           *int    `json:"groupID"`
	GroupIDSemiFinals *int    `json:"groupIDSemiFinals"`
	GroupIDFinals     *int    `json:"groupIDFinals"`
	Playoffs          int     `json:"playoffs"`
	Disqualified      int     `json:"disqualified"`
	CustomText        *string `json:"customText"`
	FPP               int     `json:"fpp"`
	TPP               int     `json:"tpp"`
	Active            bool    `json:"active"`
}
