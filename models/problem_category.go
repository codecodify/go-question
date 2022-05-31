package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId  int `gorm:"column:problem_id;type:int(11);" json:"problem_id"`
	CategoryId int `gorm:"column:category_id;type:int(11);" json:"category_id"`

	Category *Category `gorm:"foreignKey:id;references:category_id" json:"category"`
}

func (c *ProblemCategory) TableName() string {
	return "problem_category"
}
