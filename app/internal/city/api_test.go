package city

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
	repo := &mockRepository{items: []entity.City{
		{"123", "berlin", 12.1,15.1},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), logger)
	tests := []test.APITestCase{
		{"get all", "GET", "/cities", "", nil, http.StatusOK, `*"total_count":1*`},
		{"create ok", "POST", "/cities", `{"name":"test","latitude" : 45.1,"longitude" : 60.5}`, nil, http.StatusCreated, "*test*"},
		{"create ok count", "GET", "/cities", "", nil, http.StatusOK, `*"total_count":2*`},
		{"create input error", "POST", "/cities", `"name":"test"}`, nil, http.StatusBadRequest, ""},
		{"update ok", "PATCH", "/cities/123", `{"name":"cityxyz"}`, nil, http.StatusOK, "*cityxyz*"},
		{"update input error", "PATCH", "/cities/123", `"name":"cityxyz"}`, nil, http.StatusBadRequest, ""},
		{"delete ok", "DELETE", "/cities/123", ``, nil, http.StatusOK, "*cityxyz*"},
		{"delete verify", "DELETE", "/cities/123", ``, nil, http.StatusNotFound, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
