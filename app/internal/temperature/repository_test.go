package temperature

import (
	"context"
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/internal/test"
	"github.com/user/sites/app/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "temperature")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.Temperature{
		ID:        "test1",
		CityId:      "123",
		Max: 30,
		Min: 32,
		Timestamp : 124124124,

	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)


}
