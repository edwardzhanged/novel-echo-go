package book_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo/v4"
)

var bookApi = &services.BookApi{}

type CreateBookReq struct {
	WorkDirection int    `json:"workDirection"`
	CategoryId    int    `json:"categoryId"`
	CategoryName  string `json:"categoryName"`
	PicUrl        string `json:"picUrl"`
	BookName      string `json:"bookName"`
	BookDesc      string `json:"bookDesc"`
	IsVip         int    `json:"isVip"`
}

func CreateBookHandler(c echo.Context) error {
	request := new(CreateBookReq)
	if err := c.Bind(request); err != nil {
		return err
	}
	userId := c.Get("uid").(uint64)

	err := bookApi.CreateBook(userId, request.WorkDirection, request.CategoryId, request.CategoryName, request.PicUrl,
		request.BookName, request.BookDesc, request.IsVip)
	if err != nil {
		resp := handlers.CustomResponse{Code: "A0001", Message: err.Error()}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: "00000", Message: "OK"}
	return c.JSONPretty(200, &resp, "  ")
}
