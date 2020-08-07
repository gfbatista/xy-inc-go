package model

// Poi struct
type Poi struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	CoordenadaX int64  `json:"coordenadaX"`
	CoordenadaY int64  `json:"coordenadaY"`
}
