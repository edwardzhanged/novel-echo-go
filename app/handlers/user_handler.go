package handlers

import (
	"fmt"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
)

//type User struct {
//	Name  string `json:"name" form:"name" query:"name"`
//	Email string `json:"email" form:"email" query:"email"`
//}

var userApi services.UserApi

func LoginHandler(c echo.Context) error {
	u := new(services.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	userApi.Login(u.Phone, u.Password)
	return c.JSONPretty(200, "Login", "  ")
}

func RegisterHandler(c echo.Context) error {
	u := new(services.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u)
	userApi.Register(u.Username, u.Password, u.Phone)
	return c.JSONPretty(200, "Register", "  ")
}
