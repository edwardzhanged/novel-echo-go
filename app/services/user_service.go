package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login()
	Register()
}

type User struct {
	Username string `json:"username" omitempty:"username" `
	Password string `json:"password" omitempty:"password"`
	Phone    string `json:"phone"`
}

type UserApi struct{}

type UseInfo struct {
	Token    string `json:"token"`
	Uid      int    `json:"uid"`
	Nickname string `json:"nickName"`
}

var secretKey = []byte("my-secret-key")

func (u *UserApi) Login(phone string, password string) (*UseInfo, error) {

	var modelUser model.User
	conf.GbGorm.Where("phone = ?", phone).First(&modelUser)
	// 验证密码
	fmt.Println(modelUser.Password)
	err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid password:", err)
		//return nil, err
	}

	fmt.Println("Password is valid!")
	token := generateToken("123")
	return &UseInfo{
		Token: token, Uid: 123, Nickname: "edwardxx",
	}, nil
}

func (u *UserApi) Register(username string, password string, phone string) {
	// 生成密码的哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return
	}
	conf.GbGorm.Create(&model.User{Username: username, Password: string(hashedPassword), Phone: phone})

	fmt.Println("Hashed password:", string(hashedPassword))
}

func (u *UserApi) Update() {

}

func generateToken(userID string) string {
	// 将用户ID转换为字节数组
	data := []byte(userID)

	// 使用HMAC-SHA256算法生成消息认证码
	h := hmac.New(sha256.New, secretKey)
	h.Write(data)
	mac := h.Sum(nil)

	// 将消息认证码和用户ID进行拼接，并进行base64编码
	token := base64.StdEncoding.EncodeToString(append(mac, data...))

	return token
}
