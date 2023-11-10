package router

import (
	"github.com/edwardzhanged/novel-go/app/handlers/user_handler"
	"github.com/labstack/echo"
)

func InitRouter() {
	e := echo.New()
	e.POST("/api/front/user/login", user_handler.LoginHandler)
	e.POST("/api/front/user/register", user_handler.RegisterHandler)
	e.POST("/api/front/user/info", user_handler.EditUserInfoHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
