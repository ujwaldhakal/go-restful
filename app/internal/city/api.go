package city

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/user/sites/app/internal/errors"
	"github.com/user/sites/app/pkg/log"
	"github.com/user/sites/app/pkg/pagination"
	"net/http"
	"fmt"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}

	fmt.Println("ok")
	r.Get("/citys/<id>", res.get)
	r.Get("/cities", res.query)
  //r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/cities", res.create)
	r.Patch("/cities/<id>", res.update)
	r.Delete("/cities/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	city, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(city)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	citys, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = citys
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateCityRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	fmt.Println("ok its here")
	city, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	c.Response.WriteHeader(http.StatusCreated)
	return c.Write(city)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateCityRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	city, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(city)
}

func (r resource) delete(c *routing.Context) error {
	city, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(city)
}
