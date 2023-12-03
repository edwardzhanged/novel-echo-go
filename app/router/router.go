package router

import (
	"github.com/labstack/echo"
)

func InitRouter() {
	e := echo.New()
	//e.POST("/api/front/user/login", user_handler.LoginHandler)
	//e.POST("/api/front/user/register", user_handler.RegisterHandler)
	//e.POST("/api/front/user/info", user_handler.EditUserInfoHandler)
	//e.POST("/api/front/user/addBookToShelf", user_handler.AddBookToShelfHandler)
	//e.GET("/api/front/user/getBookShelf", user_handler.GetBookShelfHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
