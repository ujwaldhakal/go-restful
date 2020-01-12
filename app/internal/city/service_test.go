package city

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

func TestCreateCityRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     CreateCityRequest
		wantError bool
	}{
		{"success", CreateCityRequest{Name: "test", Latitude:123.4,Longitude:54.21}, false},
		//{"required", CreateCityRequest{Name: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestUpdateCityRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     UpdateCityRequest
		wantError bool
	}{
		{"success", UpdateCityRequest{Name: "test"}, false},
		//{"required", UpdateCityRequest{Name: ""}, false},
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
	city, err := s.Create(ctx, CreateCityRequest{Name: "test",Longitude:45.03,Latitude:99.99})
	assert.Nil(t, err)
	assert.NotEmpty(t, city.ID)
	id := city.ID
	assert.Equal(t, "test", city.Name)
	assert.NotEmpty(t, city.Latitude)
	assert.NotEmpty(t, city.Longitude)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// validation error in creation
	_, err = s.Create(ctx, CreateCityRequest{Name: ""})
	assert.NotNil(t, err)

	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	_, _ = s.Create(ctx, CreateCityRequest{Name: "test2"})

	// update
	city, err = s.Update(ctx, id, UpdateCityRequest{Name: "test updated"})
	assert.Nil(t, err)
	assert.Equal(t, "test updated", city.Name)
	_, err = s.Update(ctx, "none", UpdateCityRequest{Name: "test updated"})
	assert.NotNil(t, err)

	// validation error in update
	_, err = s.Update(ctx, id, UpdateCityRequest{Name: ""})
	assert.Nil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// unexpected error in update
	_, err = s.Update(ctx, id, UpdateCityRequest{Name: "error"})
	assert.Equal(t, errCRUD, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)

	// get
	_, err = s.Get(ctx, "none")
	assert.NotNil(t, err)
	city, err = s.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "test updated", city.Name)
	assert.Equal(t, id, city.ID)

	// query
	citys, _ := s.Query(ctx, 0, 0)
	assert.Equal(t, 1, len(citys))

	// delete
	_, err = s.Delete(ctx, "none")
	assert.NotNil(t, err)
	city, err = s.Delete(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, id, city.ID)
	count, _ = s.Count(ctx)
	assert.Equal(t, 0, count)
}

type mockRepository struct {
	items []entity.City
}

func (m mockRepository) Get(ctx context.Context, id string) (entity.City, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.City{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(ctx context.Context, offset, limit int) ([]entity.City, error) {
	return m.items, nil
}

func (m *mockRepository) Create(ctx context.Context, city entity.City) error {
	if city.Name == "error" {
		return errCRUD
	}
	m.items = append(m.items, city)
	return nil
}

func (m *mockRepository) Update(ctx context.Context, city entity.City) error {
	if city.Name == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.ID == city.ID {
			m.items[i] = city
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	for i, item := range m.items {
		if item.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}
