package user_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
)

var userApi services.UserApi

type User struct {
	ID        uint   `json:"id"`
	Nickname  string `json:"nickname" omitempty:"nickname" `
	Password  string `json:"password" omitempty:"password"`
	Phone     string `json:"phone" omitempty:"phone"`
	UserSex   string `json:"user_sex" omitempty:"user_sex"`
	UserPhoto string `json:"user_photo" omitempty:"user_photo"`
}

func LoginHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := userApi.Login(u.Phone, u.Password)
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: "用户名或密码错误"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: &user}
	return c.JSONPretty(200, &resp, "  ")
}

func RegisterHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	newUser, _ := userApi.Register(u.Nickname, u.Password, u.Phone)
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: &newUser}

	return c.JSONPretty(200, &resp, "  ")
}

func EditUserInfoHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	userApi.EditInfo(int(u.ID), u.UserSex, u.UserPhoto, u.Nickname)
	return c.JSONPretty(200, &handlers.CustomResponse{Code: 200, Message: "OK"}, "  ")
}
