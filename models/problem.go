package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	CategoryId string `gorm:"column:category_id;type:varchar(256)" json:"category_id"`
	Title      string `gorm:"column:title;type:varchar(256);" json:"title"`
	Content    string `gorm:"column:content;type:text;" json:"content"`
	TotalNum   int    `gorm:"column:total_num;type:int(11);" json:"total_num"`
}

func (p Problem) TableName() string {
	return "problem"
}
