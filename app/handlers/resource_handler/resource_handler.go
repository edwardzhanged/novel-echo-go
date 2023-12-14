package resource_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
)

var resourceApi services.ResourceApi

func GetImgVerifyCodeHandler(c echo.Context) error {
	id, b64s, _, err := resourceApi.GetImgVerifyCode()
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: "用户名或密码错误"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: map[string]string{"sessionId": id, "b64s": b64s}}
	return c.JSONPretty(200, &resp, "  ")
}
