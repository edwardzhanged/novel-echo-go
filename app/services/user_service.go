package services

import (
	"errors"
	"fmt"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"github.com/edwardzhanged/novel-go/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login()
	Register()
	EditInfo()
	AddBookToShelf()
	GetUserBookShelf()
}

type UserApi struct{}

type BookShelf struct {
	BookName string `json:"book_name"`
}

// func (u *UserApi) Login(phone string, password string) (*UserInfo, error) {
//
//		var modelUser model.User
//		conf.GbGorm.Where("phone = ?", phone).First(&modelUser)
//		// 验证密码
//		err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(password))
//		if err != nil {
//			return nil, err
//		}
//		token := generateToken(modelUser.Uuid)
//		return &UserInfo{
//			Token: token, Uid: int(modelUser.ID), Nickname: modelUser.Nickname,
//		}, nil
//	}
func (u *UserApi) Register(username string, password string, verifyCode string, sessionId string) (uid uint, token string, err error) {
	// 验证码校验
	if !utils.Store.Verify(sessionId, verifyCode, true) {
		return 0, "", errors.New("验证码错误")
	}
	// 验证手机号是否已注册
	var count int64
	conf.GbGorm.Where("username = ?", username).First(&model.UserInfo{}).Count(&count)
	if count > 0 {
		return 0, "", errors.New("手机号已注册")
	}

	// 生成密码的哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return 0, "", err
	}
	newUser := &model.UserInfo{Username: username, Password: string(hashedPassword), NickName: username}
	conf.GbGorm.Create(newUser)
	token, _ = utils.GenerateToken(newUser.ID)
	return newUser.ID, token, nil
}

//func (u *UserApi) EditInfo(uid int, userSex string, userPhoto string, nickname string) error {
//	err := conf.GbGorm.Model(&model.User{}).Where("idx = ?", uid).Updates(
//		map[string]interface{}{"Nickname": nickname, "UserSex": userSex, "UserPhoto": userPhoto}).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *UserApi) AddBookToShelf(userID, bookInfoID int, preContentId int) error {
//
//	if err := conf.GbGorm.Create(&model.UserBookShelf{UserID: userID, BookInfoID: bookInfoID, PreContentId: preContentId}); err != nil {
//		return err.Error
//
//	}
//	return nil
//}
//
//func (u *UserApi) GetUserBookShelf(userID int) ([]BookShelf, error) {
//	var result []struct {
//		BookName string
//	}
//	conf.GbGorm.Model(&model.UserBookShelf{}).Joins("left join book_infos on user_book_shelves.book_info_id = book_infos.id").
//		Where("user_book_shelves.user_id = ?", userID).
//		Find(&result)
//	fmt.Println(result)
//	return nil, nil
//}
//
//func generateToken(userID string) string {
//	// 将用户ID转换为字节数组
//	data := []byte(userID)
//
//	// 使用HMAC-SHA256算法生成消息认证码
//	h := hmac.New(sha256.New, secretKey)
//	h.Write(data)
//	mac := h.Sum(nil)
//
//	// 将消息认证码和用户ID进行拼接，并进行base64编码
//	token := base64.StdEncoding.EncodeToString(append(mac, data...))
//
//	return token
//}
