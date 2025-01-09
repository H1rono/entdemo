package router

import (
	"net/http"
	"time"

	"github.com/H1rono/entdemo/repository"
	"github.com/labstack/echo/v4"
)

func (r *Router) SetupCarRoutes(e *echo.Group) {
	e.POST("", r.createCar)
	e.GET("", r.getCars)
	e.GET("/:id", r.getCar)
	e.PUT("/:id", r.updateCar)
	e.DELETE("/:id", r.deleteCar)
}

type requestBodyCar struct {
	Model string `json:"model"`
}

type paramCarId struct {
	ID int `param:"id"`
}

type responseCar struct {
	ID           int       `json:"id"`
	Model        string    `json:"model"`
	RegisteredAt time.Time `json:"registered_at"`
}

func (r *Router) createCar(c echo.Context) error {
	var reqBody requestBodyCar
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	car, err := r.r.CreateCar(ctx, &repository.CreateCar{
		Model:        reqBody.Model,
		RegisteredAt: time.Now(),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := &responseCar{
		ID:           car.ID,
		Model:        car.Model,
		RegisteredAt: car.RegisteredAt,
	}
	return c.JSON(http.StatusCreated, res)
}

func (r *Router) getCars(c echo.Context) error {
	ctx := c.Request().Context()
	cars, err := r.r.GetCars(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := make([]*responseCar, 0, len(cars))
	for _, car := range cars {
		res = append(res, &responseCar{
			ID:           car.ID,
			Model:        car.Model,
			RegisteredAt: car.RegisteredAt,
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (r *Router) getCar(c echo.Context) error {
	binder := &echo.DefaultBinder{}
	var param paramCarId
	if err := binder.BindPathParams(c, &param); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	car, err := r.r.GetCar(ctx, param.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := &responseCar{
		ID:           car.ID,
		Model:        car.Model,
		RegisteredAt: car.RegisteredAt,
	}
	return c.JSON(http.StatusOK, res)
}

func (r *Router) updateCar(c echo.Context) error {
	binder := &echo.DefaultBinder{}
	var param paramCarId
	if err := binder.BindPathParams(c, &param); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var reqBody requestBodyCar
	if err := binder.BindBody(c, &reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	car, err := r.r.GetCar(ctx, param.ID)
	if err != nil {
		// TODO: return 404
		return c.JSON(http.StatusInternalServerError, err)
	}
	car, err = r.r.UpdateCar(ctx, param.ID, &repository.CreateCar{
		Model:        reqBody.Model,
		RegisteredAt: car.RegisteredAt,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := &responseCar{
		ID:           car.ID,
		Model:        car.Model,
		RegisteredAt: car.RegisteredAt,
	}
	return c.JSON(http.StatusOK, res)
}

func (r *Router) deleteCar(c echo.Context) error {
	binder := &echo.DefaultBinder{}
	var param paramCarId
	if err := binder.BindPathParams(c, &param); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err := r.r.DeleteCar(ctx, param.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
