package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	Name     string `gorm:"column:name;varchar(100);" json:"name"`
	ParentId int    `gorm:"column:parent_id;Index;" json:"parent_id"`
}

func (c *Category) TableName() string {
	return "category"
}
