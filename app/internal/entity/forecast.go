package entity

type Forecast struct {
  CityId      int    `json:"city_id"`
  Max int `json:"max"`
  Min int `json:"min"`
  Sample int `json:"sample"`
}
