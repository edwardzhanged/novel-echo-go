package router

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/labstack/echo"
)

func InitRouter() {
	e := echo.New()
	e.POST("/api/front/user/login", handlers.LoginHandler)
	e.POST("/user/register", handlers.RegisterHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
