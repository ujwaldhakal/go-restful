package entity

type Webhook struct {
	ID        string    `json:"id"`
	CityId      string    `json:"city_id"`
	CallbackUrl string `json:"callback_url"`
}
