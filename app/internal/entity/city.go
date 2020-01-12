package entity

type City struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
