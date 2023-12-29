package author_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo/v4"
)

var AuthorApi = &services.AuthorApi{}

func GetAuthorStatusHandler(c echo.Context) error {
	status, err := AuthorApi.GetAuthorStatus(c.Get("uid").(uint64))
	if err != nil && status != 0 {
		resp := handlers.CustomResponse{Code: "00000", Message: "对不起，您还不是作家", Data: nil}
		return c.JSONPretty(200, &resp, "  ")
	} else {
		resp := handlers.CustomResponse{Code: "00000", Message: "", Data: status}
		return c.JSONPretty(200, &resp, "  ")
	}
}

type RegisterAuthorReq struct {
	UserId        uint64 `json:"userId"`
	PenName       string `json:"penName"`
	Telephone     string `json:"telephone"`
	ChatAccount   string `json:"chatAccount"`
	Email         string `json:"email"`
	WorkDirection string `json:"workDirection"`
}

func RegisterAuthorHandler(c echo.Context) error {
	request := new(RegisterAuthorReq)
	if err := c.Bind(request); err != nil {
		return err
	}
	userId := c.Get("uid").(uint64)
	err := AuthorApi.RegisterAuthor(userId, request.PenName, request.Telephone, request.ChatAccount,
		request.Email, request.WorkDirection)
	if err != nil {
		resp := handlers.CustomResponse{Code: "A0001", Message: "注册失败"}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: "00000", Message: "OK"}
	return c.JSONPretty(200, &resp, "  ")
}
