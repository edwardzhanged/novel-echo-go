package services

import (
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
	Username string `json:"username" omitempty:"username"`
	Password string `json:"password" omitempty:"password"`
	Phone    string `json:"phone"`
}

type UserApi struct{}

func (u *UserApi) Login(phone string, password string) {

	var modelUser model.User
	conf.GbGorm.Where("phone = ?", phone).First(&modelUser)
	// 验证密码
	fmt.Println(modelUser.Password)
	err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(password))
	fmt.Println(modelUser.Phone)
	if err != nil {
		fmt.Println("Invalid password:", err)
		return
	}

	fmt.Println("Password is valid!")
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
