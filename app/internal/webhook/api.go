package webhook

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
	r.Post("/webhooks", res.create)
  r.Delete("/webhooks/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}



func (r resource) delete(c *routing.Context) error {
  album, err := r.service.Delete(c.Request.Context(), c.Param("id"))
  if err != nil {
    return err
  }

  return c.Write(album)
}

func (r resource) create(c *routing.Context) error {
	var input CreateWebhookRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	webhook, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	c.Response.WriteHeader(http.StatusCreated)
	return c.Write(webhook)
}

