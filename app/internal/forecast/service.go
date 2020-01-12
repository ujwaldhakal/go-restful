package forecast

import (
	"context"
  "errors"
  "fmt"
  //"github.com/user/sites/app/internal/city"
  validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/log"
	_ "fmt"
)

// Service encapsulates usecase logic for forecasts.
type Service interface {
  Get(ctx context.Context, id string) (Forecast, error)
}

// Forecast represents the data about an forecast.
type Forecast struct {
	entity.Forecast
}

// CreateForecastRequest represents an forecast creation request.
type CreateForecastRequest struct {
	CityId int `json:"city_id"`
	Max int `json:"max"`
	Min int `json:"min"`
}

func validateCity(value interface{}) error {
  s, _ := value.(int)
  //ctx := context.Background()
  //fmt.Println(city.Service.Get(context.TODO(),'1'))
  fmt.Println("forecast")
  if s != 12 {
    return errors.New("not a valid city")
  }
  return nil
}

// Validate validates the CreateForecastRequest fields.
func (m CreateForecastRequest) Validate() error {

	return validation.ValidateStruct(&m,
		validation.Field(&m.CityId, validation.Required),
		validation.Field(&m.Max, validation.Required),
		validation.Field(&m.Min, validation.Required),
	)
}




type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new forecast service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the forecast with the specified the forecast ID.
func (s service) Get(ctx context.Context, id string) (Forecast, error) {
	forecast, err := s.repo.Get(ctx, id)
	if err != nil {
		return Forecast{}, err
	}
	return Forecast{forecast}, nil
}
