package handlers

import (
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
)

type CustomResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var userApi services.UserApi

func LoginHandler(c echo.Context) error {
	u := new(services.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := userApi.Login(u.Phone, u.Password)
	if err != nil {
		resp := CustomResponse{Code: 201, Message: "用户名或密码错误"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := CustomResponse{Code: 200, Message: "OK", Data: &user}
	return c.JSONPretty(200, &resp, "  ")
}

func RegisterHandler(c echo.Context) error {
	u := new(services.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	newUser, _ := userApi.Register(u.Username, u.Password, u.Phone)
	resp := CustomResponse{Code: 200, Message: "OK", Data: &newUser}

	return c.JSONPretty(200, &resp, "  ")
}
