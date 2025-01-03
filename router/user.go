package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/H1rono/entdemo/repository"
)

func (r *Router) SetupUserRoutes(e *echo.Group) {
	e.POST("", r.createUser)
	e.GET("", r.getUsers)
	e.GET("/:id", r.getUser)
	e.PUT("/:id", r.updateUser)
	e.DELETE("/:id", r.deleteUser)
}

type requestBodyUser struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type paramUserId struct {
	ID int `param:"id"`
}

type responseUser struct {
	ID   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func (r *Router) createUser(c echo.Context) error {
	var reqBody requestBodyUser
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := r.r.CreateUser(c.Request().Context(), &repository.CreateUser{
		Age:  reqBody.Age,
		Name: reqBody.Name,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := &responseUser{
		ID:   user.ID,
		Age:  user.Age,
		Name: user.Name,
	}
	return c.JSON(http.StatusCreated, res)
}

func (r *Router) getUsers(c echo.Context) error {
	users, err := r.r.GetUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := make([]*responseUser, 0, len(users))
	for _, u := range users {
		res = append(res, &responseUser{
			ID:   u.ID,
			Age:  u.Age,
			Name: u.Name,
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (r *Router) getUser(c echo.Context) error {
	params := paramUserId{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user, err := r.r.GetUser(c.Request().Context(), params.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	res := &responseUser{
		ID:   user.ID,
		Age:  user.Age,
		Name: user.Name,
	}
	return c.JSON(http.StatusOK, res)
}

func (r *Router) updateUser(c echo.Context) error {
	binder := &echo.DefaultBinder{}
	var param paramUserId
	var body requestBodyUser
	if err := binder.BindPathParams(c, &param); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := binder.BindBody(c, &body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	user, err := r.r.UpdateUser(ctx, param.ID, &repository.UpdateUser{
		Age:  body.Age,
		Name: body.Name,
	})
	if err != nil {
		c.Logger().Printf("failed updating user: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (r *Router) deleteUser(c echo.Context) error {
	params := paramUserId{}
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := r.r.DeleteUser(c.Request().Context(), params.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
