package city

import (
	"context"
  "fmt"
  "github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/dbcontext"
	"github.com/user/sites/app/pkg/log"
)

// Repository encapsulates the logic to access citys from the data source.
type Repository interface {
	// Get returns the city with the specified city ID.
	Get(ctx context.Context, id string) (entity.City, error)
	// Count returns the number of citys.
	Count(ctx context.Context) (int, error)
	// Query returns the list of citys with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.City, error)
	// Create saves a new city in the storage.
	Create(ctx context.Context, city entity.City) error
	// Update updates the city with given ID in the storage.
	Update(ctx context.Context, city entity.City) error
	// Delete removes the city with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists citys in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new city repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

//// Get reads the city with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.City, error) {
	var city entity.City
	err := r.db.With(ctx).Select().Model(id, &city)
	return city, err
}

// Create saves a new city record in the database.
// It returns the ID of the newly inserted city record.
func (r repository) Create(ctx context.Context, city entity.City) error {
  fmt.Println("ok now its upto here");
	return r.db.With(ctx).Model(&city).Insert()
}

// Update saves the changes to an city in the database.
func (r repository) Update(ctx context.Context, city entity.City) error {
	return r.db.With(ctx).Model(&city).Update()
}

// Delete deletes an city with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	city, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&city).Delete()
}

// Count returns the number of the city records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("city").Row(&count)
	return count, err
}

// Query retrieves the city records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.City, error) {
	var citys []entity.City
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&citys)
	return citys, err
}
