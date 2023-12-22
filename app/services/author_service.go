package services

import (
	"errors"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"gorm.io/gorm"
)

type AuthorService interface {
	GetAuthorStatus()
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
