package forecast

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/user/sites/app/pkg/log"
	_ "fmt"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.Get("/forecasts/<id>", res.get)
}

type resource struct {
	service Service
	logger  log.Logger
}



func (r resource) get(c *routing.Context) error {
  forecast, err := r.service.Get(c.Request.Context(), c.Param("id"))
  if err != nil {
    return err
  }

  return c.Write(forecast)
}
