package router

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/labstack/echo"
)

func InitRouter() {
	e := echo.New()
	e.GET("/user/login", handlers.LoginHandler)
	e.GET("/user/register", handlers.RegisterHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
