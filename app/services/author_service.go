package services

import (
	"errors"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"gorm.io/gorm"
	"strconv"
)

type AuthorService interface {
	GetAuthorStatus()
	RegisterAuthor()
}
type AuthorApi struct{}

func (author *AuthorApi) GetAuthorStatus(uid uint64) (status uint8, err error) {
	var authorInfo model.AuthorInfo
	result := conf.GbGorm.Where("user_id = ?", uid).First(&authorInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 1, result.Error
	}
	if authorInfo.Status == 0 {
		return 0, nil
	} else {
		return 1, nil
	}
}

func (author *AuthorApi) RegisterAuthor(uid uint64, penName string, telephone string, chatAccount string, email string,
	workDirection string) (err error) {
	parsedUint64, _ := strconv.ParseUint(workDirection, 10, 8)
	result := conf.GbGorm.Create(&model.AuthorInfo{UserID: uid, PenName: penName, TelPhone: telephone,
		ChatAccount: chatAccount, WorkDirection: uint8(parsedUint64), Email: email, InviteCodeID: "1"})
	if result.Error == nil {
		return nil
	} else {
		return result.Error
	}
}
