package temperature

import (
	"context"
  "fmt"
  "github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/dbcontext"
	"github.com/user/sites/app/pkg/log"
)

// Repository encapsulates the logic to access temperatures from the data source.
type Repository interface {
	// Get returns the temperature with the specified temperature ID.
	Get(ctx context.Context, id string) (entity.Temperature, error)
	// Count returns the number of temperatures.
	Count(ctx context.Context) (int, error)
	// Create saves a new temperature in the storage.
	Create(ctx context.Context, temperature entity.Temperature) error

  GetWebhookLists(ctx context.Context, cityId string) ([]entity.Webhook, error)
	// Update updates the temperature with given ID in the storage.
}

// repository persists temperatures in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new temperature repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the temperature with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Temperature, error) {
	var temperature entity.Temperature
	err := r.db.With(ctx).Select().Model(id, &temperature)
	return temperature, err
}

// Get reads the temperature with the specified ID from the database.
func (r repository) GetWebhookLists(ctx context.Context, cityId string) ([]entity.Webhook, error) {
	var webhook[] entity.Webhook
  query := r.db.With(ctx).NewQuery("SELECT * FROM webhook where city_id = "+"'"+cityId+"'")
	err := query.All(&webhook)
	return webhook, err
}

// Create saves a new temperature record in the database.
// It returns the ID of the newly inserted temperature record.
func (r repository) Create(ctx context.Context, temperature entity.Temperature) error {
  fmt.Println("ok now its upto here");
	return r.db.With(ctx).Model(&temperature).Insert()
}

// Update saves the changes to an temperature in the database.
func (r repository) Update(ctx context.Context, temperature entity.Temperature) error {
	return r.db.With(ctx).Model(&temperature).Update()
}

// Delete deletes an temperature with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	temperature, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&temperature).Delete()
}

// Count returns the number of the temperature records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("temperature").Row(&count)
	return count, err
}

// Query retrieves the temperature records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Temperature, error) {
	var temperatures []entity.Temperature
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&temperatures)
	return temperatures, err
}
