package temperature

import (
  "bytes"
  "context"
  "encoding/json"
  "errors"
  "fmt"
  validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/log"
	_ "fmt"
	"net/http"
  _ "net/url"
  "time"
)

// Service encapsulates usecase logic for temperatures.
type Service interface {
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTemperatureRequest) (Temperature, error)
}

// Temperature represents the data about an temperature.
type Temperature struct {
	entity.Temperature
}

// CreateTemperatureRequest represents an temperature creation request.
type CreateTemperatureRequest struct {
	CityId string `json:"city_id"`
	Max int `json:"max"`
	Min int `json:"min"`
}

func validateCity(value interface{}) error {
  s, _ := value.(int)
  //ctx := context.Background()
  //fmt.Println(city.Service.Get(context.TODO(),'1'))
  fmt.Println("temperature")
  if s != 12 {
    return errors.New("not a valid city")
  }
  return nil
}

// Validate validates the CreateTemperatureRequest fields.
func (m CreateTemperatureRequest) Validate() error {

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

// NewService creates a new temperature service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the temperature with the specified the temperature ID.
func (s service) Get(ctx context.Context, id string) (Temperature, error) {
	temperature, err := s.repo.Get(ctx, id)
	if err != nil {
		return Temperature{}, err
	}
	return Temperature{temperature}, nil
}

// Create creates a new temperature.
func (s service) Create(ctx context.Context, req CreateTemperatureRequest) (Temperature, error) {
	if err := req.Validate(); err != nil {
		return Temperature{}, err
	}

  id := entity.GenerateID()
  data := entity.Temperature{
    ID:        id,
    CityId:      req.CityId,
    Max:  req.Max,
    Min: req.Min,
    Timestamp: int(time.Now().Unix()),

  }
	err := s.repo.Create(ctx, data)

	if err != nil {
		return Temperature{}, err
	}

	go s.callWebhooks(ctx,data)

  return s.Get(ctx, id)
}

func (s service) callWebhooks(ctx context.Context,temperature entity.Temperature) {
  fmt.Println("print webhooknow")
  webhooks,err := s.repo.GetWebhookLists(ctx,temperature.CityId)
  if err == nil {
    for data,_ := range webhooks {
      webhookUrl := webhooks[data].CallbackUrl
      formData,_ := json.Marshal(temperature)
      res,err := http.Post(webhookUrl,"application/json",bytes.NewBuffer(formData))
      fmt.Println("webhook sent")
      fmt.Println(res)
      fmt.Println(err)
    }
  }

}

// Count returns the number of temperatures.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}
