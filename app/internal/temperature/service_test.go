package temperature

import (
	"context"
	"database/sql"
	"errors"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errCRUD = errors.New("error crud")

func TestCreateTemperatureRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     CreateTemperatureRequest
		wantError bool
	}{
		{"success", CreateTemperatureRequest{
      CityId:      "123",
      Max: 30,
      Min: 32,
}, false},
		//{"required", CreateTemperatureRequest{Name: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}


func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger)

	ctx := context.Background()

	// initial count
	count, _ := s.Count(ctx)
	assert.Equal(t, 0, count)

	// successful creation
	temperature, err := s.Create(ctx, CreateTemperatureRequest{
    CityId:      "123",
    Max: 30,
    Min: 32})
	assert.Nil(t, err)
	assert.Equal(t, "test", temperature.CityId)
	assert.NotEmpty(t, temperature.Min)
	assert.NotEmpty(t, temperature.Max)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// validation error in creation
	_, err = s.Create(ctx, CreateTemperatureRequest{})
	assert.NotNil(t, err)

	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

}

type mockRepository struct {
	items []entity.Temperature
}

func (m mockRepository) Get(ctx context.Context, id string) (entity.Temperature, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Temperature{}, sql.ErrNoRows
}
func (m mockRepository) GetWebhookLists(ctx context.Context, id string) ([]entity.Webhook, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return []entity.Webhook{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(ctx context.Context, offset, limit int) ([]entity.Temperature, error) {
	return m.items, nil
}

func (m *mockRepository) Create(ctx context.Context, temperature entity.Temperature) error {
	if temperature.CityId == "error" {
		return errCRUD
	}
	m.items = append(m.items, temperature)
	return nil
}

