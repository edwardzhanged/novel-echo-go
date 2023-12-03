package model

import (
	"gorm.io/gorm"
	"time"
)

type BookInfo struct {
	gorm.Model
	WorkDirection         uint8             `gorm:"type:tinyint unsigned;default:0" json:"work_direction" comment:"作品方向 0-男频 1-女频"`
	CategoryID            uint              `gorm:"type:bigint unsigned;" json:"category_id" comment:"类别ID"`
	CategoryName          string            `gorm:"type:varchar(50);not null" json:"category_name" comment:"类别名"`
	PicURL                string            `gorm:"type:varchar(200);not null" json:"pic_url" comment:"小说封面地址"`
	BookName              string            `gorm:"type:varchar(50);not null" json:"book_name" comment:"小说名"`
	AuthorID              uint              `gorm:"type:bigint unsigned; not null" json:"author_id"`
	AuthorName            string            `gorm:"type:varchar(50);not null" json:"author_name" comment:"作家名"`
	BookDesc              string            `gorm:"type:varchar(2000);not null" json:"book_desc" comment:"书籍描述"`
	Score                 uint8             `gorm:"type:tinyint unsigned;default:0" json:"score" comment:"评分 0-10"`
	BookStatus            uint8             `gorm:"type:tinyint unsigned;default:0" json:"book_status" comment:"书籍状态 0-连载中 1-已完结"`
	VisitCount            uint64            `gorm:"type:bigint unsigned;default:103" json:"visit_count" comment:"点击量"`
	WordCount             uint32            `gorm:"type:int unsigned;default:0" json:"word_count" comment:"总字数"`
	CommentCount          uint32            `gorm:"type:int unsigned;default:0" json:"comment_count" comment:"评论数"`
	LastChapterID         uint64            `gorm:"type:bigint unsigned" json:"last_chapter_id" comment:"最新章节ID"`
	LastChapterName       string            `gorm:"type:varchar(50)" json:"last_chapter_name" comment:"最新章节名"`
	LastChapterUpdateTime time.Time         `gorm:"type:datetime" json:"last_chapter_update_time" comment:"最新章节更新时间"`
	IsVIP                 uint8             `gorm:"type:tinyint unsigned;default:0" json:"is_vip" comment:"是否收费 0-免费 1-收费"`
	BookChapters          []BookChapter     `gorm:"foreignKey:BookID"`
	BookComments          []BookComment     `gorm:"foreignKey:BookID"`
	UserBookshelves       []UserBookshelf   `gorm:"foreignKey:BookID"`
	UserReadHistorys      []UserReadHistory `gorm:"foreignKey:BookID"`
}

type BookCategory struct {
	gorm.Model
	WorkDirection uint8      `gorm:"type:tinyint unsigned;not null" json:"work_direction" comment:"作品方向"`
	Name          string     `gorm:"type:varchar(20);not null" json:"name" comment:"类别名"`
	Sort          uint8      `gorm:"type:tinyint unsigned;not null;default:10" json:"sort" comment:"排序"`
	BookInfos     []BookInfo `gorm:"foreignKey:CategoryID"`
}

type BookChapter struct {
	gorm.Model
	BookID      uint64 `gorm:"type:bigint unsigned;not null" json:"book_id" comment:"小说ID"`
	ChapterNum  uint16 `gorm:"type:smallint unsigned;not null" json:"chapter_num" comment:"章节号"`
	ChapterName string `gorm:"type:varchar(100);not null" json:"chapter_name" comment:"章节名"`
	WordCount   uint32 `gorm:"type:int unsigned;not null" json:"word_count" comment:"章节字数"`
	IsVIP       uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"is_vip" comment:"是否收费 0-免费 1-收费"`
}

type BookContent struct {
	gorm.Model
	ChapterID   uint64      `gorm:"type:bigint unsigned;not null" json:"chapter_id" comment:"章节ID"`
	BookChapter BookChapter `gorm:"foreignKey:ChapterID"`
	Content     string      `gorm:"type:mediumtext;not null" json:"content" comment:"小说章节内容"`
}

type BookComment struct {
	gorm.Model
	BookID             uint64             `gorm:"type:bigint unsigned;not null" json:"book_id" comment:"评论小说ID"`
	UserID             uint64             `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"评论用户ID"`
	CommentContent     string             `gorm:"type:varchar(512);not null" json:"comment_content" comment:"评价内容"`
	ReplyCount         uint32             `gorm:"type:int unsigned;not null;default:0" json:"reply_count" comment:"回复数量"`
	AuditStatus        uint8              `gorm:"type:tinyint unsigned;not null;default:0" json:"audit_status" comment:"审核状态"`
	BookCommentReplies []BookCommentReply `gorm:"foreignKey:CommentID"`
}

type BookCommentReply struct {
	gorm.Model
	CommentID    uint64 `gorm:"type:bigint unsigned;not null" json:"comment_id" comment:"评论ID"`
	UserID       uint64 `gorm:"type:bigint unsigned;not null" json:"user_id" comment:"回复用户ID"`
	ReplyContent string `gorm:"type:varchar(512);not null" json:"reply_content" comment:"回复内容"`
	AuditStatus  uint8  `gorm:"type:tinyint unsigned;not null;default:0" json:"audit_status" comment:"审核状态"`
}
