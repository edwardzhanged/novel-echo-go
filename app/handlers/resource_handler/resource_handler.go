package resource_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo/v4"
)

var ResourceApi = &services.ResourceApi{}

func GetImgVerifyCodeHandler(c echo.Context) error {
	id, b64s, err := ResourceApi.GetImgVerifyCode()
	if err != nil {
		resp := handlers.CustomResponse{Code: "A0001", Message: "用户名或密码错误"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: "00000", Message: "OK", Data: map[string]string{"sessionId": id, "b64s": b64s}}
	return c.JSONPretty(200, &resp, "  ")
}

func VerifyImgAnswerHandler(c echo.Context) error {
	id := c.FormValue("sessionId")
	answer := c.FormValue("answer")
	if ResourceApi.VerifyImgAnswer(id, answer) {
		resp := handlers.CustomResponse{Code: "200", Message: "OK"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: "201", Message: "验证码错误"}
	return c.JSONPretty(200, &resp, "  ")
}
