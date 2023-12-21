package user_handler

import (
	"github.com/edwardzhanged/novel-go/app/handlers"
	"github.com/edwardzhanged/novel-go/app/services"
	"github.com/labstack/echo"
	"strconv"
)

var userApi = &services.UserApi{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Uid      uint   `json:"uid"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

func LoginHandler(c echo.Context) error {
	request := new(LoginReq)
	if err := c.Bind(request); err != nil {
		return err
	}
	uid, nickname, token, err := userApi.Login(request.Username, request.Password)
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: err.Error()}
		return c.JSONPretty(200, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: LoginResp{Uid: uid, Nickname: nickname, Token: token}}
	return c.JSONPretty(200, &resp, "  ")
}

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
	request := new(RegisterReq)
	if err := c.Bind(request); err != nil {
		return err
	}
	userId, token, err := userApi.Register(request.Username, request.Password, request.VerifyCode, request.SessionId)
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: err.Error()}
		return c.JSONPretty(201, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: RegisterResp{Uid: userId, Token: token}}

	return c.JSONPretty(200, &resp, "  ")
}

type GetUserInfoResp struct {
	Nickname string `json:"nickName"`
	Photo    string `json:"userPhoto"`
	Sex      uint8  `json:"userSex"`
}

func GetUserInfoHandler(c echo.Context) error {
	userId := c.Param("id")
	userIdInt, err := strconv.ParseUint(userId, 10, 64)
	nickname, sex, photo, err := userApi.GetUserInfo(uint(userIdInt))
	if err != nil {
		resp := handlers.CustomResponse{Code: 201, Message: err.Error()}
		return c.JSONPretty(201, &resp, "  ")
	}
	resp := handlers.CustomResponse{Code: 200, Message: "OK", Data: GetUserInfoResp{
		Nickname: nickname,
		Photo:    photo,
		Sex:      sex,
	}}
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
