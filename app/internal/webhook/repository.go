package webhook

import (
	"context"
  "fmt"
  "github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/pkg/dbcontext"
	"github.com/user/sites/app/pkg/log"
)

// Repository encapsulates the logic to access webhooks from the data source.
type Repository interface {
	// Get returns the webhook with the specified webhook ID.
	Get(ctx context.Context, id string) (entity.Webhook, error)
	// Count returns the number of webhooks.
	Count(ctx context.Context) (int, error)
	// Query returns the list of webhooks with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Webhook, error)
	// Create saves a new webhook in the storage.
	Create(ctx context.Context, webhook entity.Webhook) error
	// Update updates the webhook with given ID in the storage.
	Update(ctx context.Context, webhook entity.Webhook) error
	// Delete removes the webhook with given ID from the storage.
	Delete(ctx context.Context, id string) error

  FindById()
}

// repository persists webhooks in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new webhook repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the webhook with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Webhook, error) {
	var webhook entity.Webhook
	err := r.db.With(ctx).Select().Model(id, &webhook)
	return webhook, err
}

func (r repository) FindById() {
  fmt.Println("get me data man")
}

// Create saves a new webhook record in the database.
// It returns the ID of the newly inserted webhook record.
func (r repository) Create(ctx context.Context, webhook entity.Webhook) error {
  fmt.Println("ok now its upto here");
	return r.db.With(ctx).Model(&webhook).Insert()
}

// Update saves the changes to an webhook in the database.
func (r repository) Update(ctx context.Context, webhook entity.Webhook) error {
	return r.db.With(ctx).Model(&webhook).Update()
}

// Delete deletes an webhook with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	webhook, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&webhook).Delete()
}

// Count returns the number of the webhook records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("webhook").Row(&count)
	return count, err
}

// Query retrieves the webhook records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Webhook, error) {
	var webhooks []entity.Webhook
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&webhooks)
	return webhooks, err
}
