package temperature

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/user/sites/app/internal/errors"
	"github.com/user/sites/app/pkg/log"
	"net/http"
	_ "fmt"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.Post("/temperatures", res.create)

}

type resource struct {
	service Service
	logger  log.Logger
}



func (r resource) create(c *routing.Context) error {
	var input CreateTemperatureRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	temperature, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	c.Response.WriteHeader(http.StatusCreated)
	return c.Write(temperature)
}

