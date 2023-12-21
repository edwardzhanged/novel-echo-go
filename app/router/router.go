package router

import (
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/handlers/resource_handler"
	"github.com/edwardzhanged/novel-go/app/handlers/user_handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func InitRouter() {
	e := echo.New()

	skipJWT := func(c echo.Context) bool {
		skipApis := map[string]bool{"/api/front/user/login": true, "/api/front/user/register": true, "/api/front/resource/img_verify_code": true}
		if skipApis[c.Path()] {
			return true
		}
		return false
	}
	jwtConfig := middleware.JWTConfig{
		Skipper:    skipJWT,
		SigningKey: []byte(conf.GbViper.GetString("jwt.secret")),
		ErrorHandler: func(err error) error {
			// 在这里你可以自定义错误信息的处理逻辑
			return echo.NewHTTPError(http.StatusUnauthorized, "请先登录")
		},
	}
	e.Use(middleware.JWTWithConfig(jwtConfig))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// User Router
	userRouter := e.Group("/api/front/user")
	userRouter.POST("/login", user_handler.LoginHandler)
	userRouter.POST("/register", user_handler.RegisterHandler)
	userRouter.GET("/:id", user_handler.GetUserInfoHandler)

	//userRouter.POST("/info", user_handler.EditUserInfoHandler)
	//userRouter.POST("/addBookToShelf", user_handler.AddBookToShelfHandler)
	//userRouter.POST("/getBookShelf", user_handler.getBookShelf)

	// Resource Router
	resourceRouter := e.Group("/api/front/resource")
	resourceRouter.GET("/img_verify_code", resource_handler.GetImgVerifyCodeHandler)
	resourceRouter.POST("/verify_img_answer", resource_handler.VerifyImgAnswerHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
