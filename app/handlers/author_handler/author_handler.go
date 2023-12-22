package author_handler

import (
	"fmt"
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo/v4"
)

var AuthorApi = &services.AuthorApi{}

func GetAuthorStatusHandler(c echo.Context) error {
	status, err := AuthorApi.GetAuthorStatus(c.Get("uid").(uint64))
	fmt.Println(status, err)
	if err != nil && status != 0 {
		resp := handlers.CustomResponse{Code: "A0001", Message: "对不起，您还不是作家", Data: nil}
		return c.JSONPretty(200, &resp, "  ")
	} else {
		resp := handlers.CustomResponse{Code: "00000", Message: "", Data: status}
		return c.JSONPretty(200, &resp, "  ")
	}
}
