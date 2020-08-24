package main

import (
	"github.com/labstack/echo"
)

type Usercontroller struct{}

func (u Usercontroller) RegisterRoute(g *echo.Group) {
	g.GET("/sign_up", u.signUp)
	g.POST("/user/register", u.register)
}

func (u Usercontroller) signUp(ctx echo.Context) error {
	return render.NoNavRender(ctx, "user/sign_up.html")
}
func (u Usercontroller) register(ctx echo.Context) error {
	return nil
}
