package city

import (
	"context"
	"database/sql"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/internal/test"
	"github.com/user/sites/app/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "city")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.City{
		ID:        "test1",
		Name:      "city1",
		Longitude: 51.25,
		Latitude: 75.03,
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	city, err := repo.Get(ctx, "test1")
	assert.Nil(t, err)
	assert.Equal(t, "city1", city.Name)
	_, err = repo.Get(ctx, "test0")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, entity.City{
		ID:        "test1",
		Name:      "city1 updated",
    Longitude: 51.25,
    Latitude: 75.03,
	})
	assert.Nil(t, err)
	city, _ = repo.Get(ctx, "test1")
	assert.Equal(t, "city1 updated", city.Name)

	// query
	citys, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(citys))

	// delete
	err = repo.Delete(ctx, "test1")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
}
