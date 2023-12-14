package router

import (
	"github.com/edwardzhanged/novel-go/app/handlers/resource_handler"
	"github.com/labstack/echo"
)

func InitRouter() {
	e := echo.New()
	// User Router
	//userRouter := e.Group("/api/front/user")
	//userRouter.POST("/login", user_handler.LoginHandler)
	//userRouter.POST("/register", user_handler.RegisterHandler)
	//userRouter.POST("/info", user_handler.EditUserInfoHandler)
	//userRouter.POST("/addBookToShelf", user_handler.AddBookToShelfHandler)
	//userRouter.POST("/getBookShelf", user_handler.getBookShelf)

	// Resource Router
	resourceRouter := e.Group("/api/front/resource")
	resourceRouter.GET("/img_verify_code", resource_handler.GetImgVerifyCodeHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
