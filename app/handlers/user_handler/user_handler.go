package user_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
)

var userApi = &services.UserApi{}

//	func LoginHandler(c echo.Context) error {
//		u := new(User)
//		if err := c.Bind(u); err != nil {
//			return err
//		}
//		user, err := userApi.Login(u.Phone, u.Password)
//		if err != nil {
//			resp := handlers.CustomResponse{Code: 201, Message: "用户名或密码错误"}
//			return c.JSONPretty(200, &resp, "  ")
//		}
//		resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: &user}
//		return c.JSONPretty(200, &resp, "  ")
//	}
type RegisterReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"`
	SessionId  string `json:"session_id"`
}
type RegisterResp struct {
	Uid   uint   `json:"uid"`
	Token string `json:"token"`
}

func RegisterHandler(c echo.Context) error {
	u := new(RegisterReq)
	if err := c.Bind(u); err != nil {
		return err
	}
	userId, token, err := userApi.Register(u.Username, u.Password, u.VerifyCode, u.SessionId)
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: err.Error()}
		return c.JSONPretty(201, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: RegisterResp{Uid: userId, Token: token}}

	return c.JSONPretty(200, &resp, "  ")
}

//func EditUserInfoHandler(c echo.Context) error {
//	u := new(User)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	if err := userApi.EditInfo(int(u.ID), u.UserSex, u.UserPhoto, u.Nickname); err != nil {
//		resp := handlers.CustomResponse{Code: 201, Message: error.Error(err)}
//		return c.JSONPretty(200, &resp, "  ")
//	}
//	return c.JSONPretty(200, &handlers.CustomResponse{Code: 200, Message: "OK"}, "  ")
//}
//
//type BookShelf struct {
//	UserID       int `json:"user_id"`
//	BookInfoID   int `json:"book_id"`
//	PreContentId int `json:"pre_content_id"`
//}
//
//func AddBookToShelfHandler(c echo.Context) error {
//	u := new(BookShelf)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	if err := userApi.AddBookToShelf(u.UserID, u.BookInfoID, u.PreContentId); err != nil {
//		resp := handlers.CustomResponse{Code: 201, Message: error.Error(err)}
//		return c.JSONPretty(200, &resp, "  ")
//	}
//	return c.JSONPretty(200, &handlers.CustomResponse{Code: 200, Message: "OK"}, "  ")
//}
//
//func GetBookShelfHandler(c echo.Context) error {
//	userID := c.QueryParam("userID")
//	i, err := strconv.Atoi(userID)
//
//	bookShelf, err := userApi.GetUserBookShelf(i)
//	if err != nil {
//		resp := handlers.CustomResponse{Code: 201, Message: error.Error(err)}
//		return c.JSONPretty(200, &resp, "  ")
//	}
//	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: &bookShelf}
//	return c.JSONPretty(200, &resp, "  ")
//}
