package forecast

import (
	"context"
  _ "fmt"
  "github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/dbcontext"
	"github.com/user/sites/app/pkg/log"
)

// Repository encapsulates the logic to access forecasts from the data source.
type Repository interface {
	// Get returns the forecast with the specified forecast ID.
	Get(ctx context.Context, id string) (entity.Forecast, error)
	// Count returns the number of forecasts.
}

// repository persists forecasts in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new forecast repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the forecast with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Forecast, error) {
	var forecast entity.Forecast
	query := r.db.With(ctx).NewQuery("SELECT temperature.city_id,temperature.min,temperature.max,temperature.id as saWmple FROM temperature where city_id = "+"'"+id+"'")

	err := query.One(&forecast)

	return forecast, err
}


