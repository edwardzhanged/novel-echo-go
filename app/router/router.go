package router

import (
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/handlers/author_handler"
	"github.com/edwardzhanged/novel-go/app/handlers/resource_handler"
	"github.com/edwardzhanged/novel-go/app/handlers/user_handler"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func InitRouter() {
	e := echo.New()
	jwtConfig := echojwt.Config{
		Skipper:    skipper,
		KeyFunc:    verifyToken,
		ContextKey: "token",
	}
	e.Use(echojwt.WithConfig(jwtConfig))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(ExtractUserIDFromToken)
	e.Use(middleware.Recover())

	// User Router
	userRouter := e.Group("/api/front/user")
	userRouter.POST("/login", user_handler.LoginHandler)
	userRouter.POST("/register", user_handler.RegisterHandler)
	userRouter.GET("", user_handler.GetUserInfoHandler)

	//userRouter.POST("/info", user_handler.EditUserInfoHandler)
	//userRouter.POST("/addBookToShelf", user_handler.AddBookToShelfHandler)
	//userRouter.POST("/getBookShelf", user_handler.getBookShelf)

	// Resource Router
	resourceRouter := e.Group("/api/front/resource")
	resourceRouter.GET("/img_verify_code", resource_handler.GetImgVerifyCodeHandler)
	resourceRouter.POST("/verify_img_answer", resource_handler.VerifyImgAnswerHandler)

	authorRouter := e.Group("/api/author")
	authorRouter.GET("/status", author_handler.GetAuthorStatusHandler)
	authorRouter.POST("/register", author_handler.RegisterAuthorHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
func ExtractUserIDFromToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("token") == nil {
			return next(c)
		}
		rawUserId := c.Get("token").(*jwt.Token).Claims.(jwt.MapClaims)["userid"].(float64)
		userId := uint64(rawUserId)
		c.Set("uid", userId)
		return next(c)
	}
}

func skipper(c echo.Context) bool {
	skipApis := map[string]bool{"/api/front/user/login": true, "/api/front/user/register": true, "/api/front/resource/img_verify_code": true}
	if skipApis[c.Path()] {
		return true
	}
	return false
}

func verifyToken(t *jwt.Token) (interface{}, error) {
	var token string
	token = t.Raw
	secret := []byte(conf.GbViper.GetString("jwt.secret"))
	_, err := jwt.Parse(strings.TrimPrefix(token, "Bearer "), func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	return secret, nil

}
