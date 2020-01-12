package entity

type Temperature struct {
	ID        string    `json:"id"`
	CityId      string    `json:"city_id"`
	Max int `json:"max"`
	Min int `json:"min"`
	Timestamp int `json:"timestamp"`
}
