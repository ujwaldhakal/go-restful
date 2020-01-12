package temperature

import (
	"github.com/user/sites/app/internal/entity"
	"github.com/user/sites/app/internal/test"
	"github.com/user/sites/app/pkg/log"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Temperature{
		{"123", "berlin", 12,15,1241241},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), logger)
	tests := []test.APITestCase{
		{"create ok", "POST", "/temperatures", `{"city_id":"1","max" : 45.1,"min" : 60.5}`, nil, http.StatusCreated, "*test*"},
		{"create ok count", "GET", "/cititemperatureses", "", nil, http.StatusOK, `*"total_count":2*`},
		{"create input error", "POST", "/temperatures", `"name":"test"}`, nil, http.StatusBadRequest, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
