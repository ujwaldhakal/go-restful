package city

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/log"
	"fmt"
)

// Service encapsulates usecase logic for citys.
type Service interface {
	Get(ctx context.Context, id string) (City, error)
	Query(ctx context.Context, offset, limit int) ([]City, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateCityRequest) (City, error)
	Update(ctx context.Context, id string, input UpdateCityRequest) (City, error)
	Delete(ctx context.Context, id string) (City, error)
}

// City represents the data about an city.
type City struct {
	entity.City
}

// CreateCityRequest represents an city creation request.
type CreateCityRequest struct {
	Name string `json:"name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Validate validates the CreateCityRequest fields.
func (m CreateCityRequest) Validate() error {

  fmt.Println(m)
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required),
		validation.Field(&m.Latitude, validation.Required),
		validation.Field(&m.Longitude, validation.Required),
	)
}

// UpdateCityRequest represents an city update request.
type UpdateCityRequest struct {
	Name string `json:"name,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
}

// Validate validates the CreateCityRequest fields.
func (m UpdateCityRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new city service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the city with the specified the city ID.
func (s service) Get(ctx context.Context, id string) (City, error) {
	city, err := s.repo.Get(ctx, id)
	if err != nil {
		return City{}, err
	}
	return City{city}, nil
}

// Create creates a new city.
func (s service) Create(ctx context.Context, req CreateCityRequest) (City, error) {
	if err := req.Validate(); err != nil {
		return City{}, err
	}
	id := entity.GenerateID()
	//_, now := time.Now()
	fmt.Println("is it upto here")
	err := s.repo.Create(ctx, entity.City{
		ID:        id,
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,

	})
	if err != nil {
		return City{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the city with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateCityRequest) (City, error) {
	if err := req.Validate(); err != nil {
		return City{}, err
	}

	city, err := s.Get(ctx, id)
	if err != nil {
		return city, err
	}
	if req.Name != "" {
   city.Name = req.Name
  }
  if req.Latitude > 0 {
   city.Latitude = req.Latitude
  }

  if req.Longitude > 0 {
   city.Longitude = req.Longitude
  }

	if err := s.repo.Update(ctx, city.City); err != nil {
		return city, err
	}
	return city, nil
}

// Delete deletes the city with the specified ID.
func (s service) Delete(ctx context.Context, id string) (City, error) {
	city, err := s.Get(ctx, id)
	if err != nil {
		return City{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return City{}, err
	}
	return city, nil
}

// Count returns the number of citys.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the citys with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]City, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []City{}
	for _, item := range items {
		result = append(result, City{item})
	}
	return result, nil
}
