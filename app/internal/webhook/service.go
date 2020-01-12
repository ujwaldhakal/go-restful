package webhook

import (
	"context"
  "errors"
  "fmt"
  //"github.com/user/sites/app/internal/city"
  validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/log"
	_ "fmt"
  _ "time"
)

// Service encapsulates usecase logic for webhooks.
type Service interface {
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateWebhookRequest) (Webhook, error)
  Delete(ctx context.Context, id string) (Webhook, error)
}

// Webhook represents the data about an webhook.
type Webhook struct {
	entity.Webhook
}

// CreateWebhookRequest represents an webhook creation request.
type CreateWebhookRequest struct {
	CityId string `json:"city_id"`
	CallbackUrl string `json:"callback_url"`
}

func validateCity(value interface{}) error {
  s, _ := value.(int)
  //ctx := context.Background()
  //fmt.Println(city.Service.Get(context.TODO(),'1'))
  fmt.Println("webhook")
  if s != 12 {
    return errors.New("not a valid city")
  }
  return nil
}

// Validate validates the CreateWebhookRequest fields.
func (m CreateWebhookRequest) Validate() error {

	return validation.ValidateStruct(&m,
		validation.Field(&m.CityId, validation.Required),
		validation.Field(&m.CallbackUrl, validation.Required),
	)
}



type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new webhook service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the webhook with the specified the webhook ID.
func (s service) Get(ctx context.Context, id string) (Webhook, error) {
	webhook, err := s.repo.Get(ctx, id)
	if err != nil {
		return Webhook{}, err
	}
	return Webhook{webhook}, nil
}

// Create creates a new webhook.
func (s service) Create(ctx context.Context, req CreateWebhookRequest) (Webhook, error) {
	if err := req.Validate(); err != nil {
		return Webhook{}, err
	}
	id := entity.GenerateID()
	err := s.repo.Create(ctx, entity.Webhook{
		ID:        id,
		CityId:      req.CityId,
		CallbackUrl:  req.CallbackUrl,

	})
	if err != nil {
		return Webhook{}, err
	}
	return s.Get(ctx, id)
}

// Count returns the number of webhooks.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Delete deletes the webhook with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Webhook, error) {
  webhook, err := s.Get(ctx, id)
  if err != nil {
    return Webhook{}, err
  }
  if err = s.repo.Delete(ctx, id); err != nil {
    return Webhook{}, err
  }
  return webhook, nil
}
