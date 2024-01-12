package services

import (
	"errors"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/edwardzhanged/novel-go/app/model"
	"gorm.io/gorm"
	"time"
)

type BookService interface {
	CreateBook()
}

type BookApi struct{}

func (book *BookApi) CreateBook(userId uint64, workDirection int, categoryId int, categoryName string, picUrl string, bookName string,
	bookDesc string, isVip int) error {
	var authorInfo model.AuthorInfo
	result := conf.GbGorm.Where("user_id = ? AND status = ?", userId, 0).First(&authorInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("作者不存在或者状态不正确")
	}
	conf.GbGorm.Create(&model.BookInfo{
		WorkDirection:         uint8(workDirection),
		CategoryID:            uint(categoryId),
		CategoryName:          categoryName,
		PicURL:                picUrl,
		BookName:              bookName,
		BookDesc:              bookDesc,
		IsVIP:                 uint8(isVip),
		LastChapterUpdateTime: time.Now(),
		AuthorID:              authorInfo.ID,
	})
	return nil
}
