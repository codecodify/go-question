package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	Title    string `gorm:"column:title;type:varchar(256);" json:"title"`
	Content  string `gorm:"column:content;type:text;" json:"content"`
	TotalNum int    `gorm:"column:total_num;type:int(11);" json:"total_num"`

	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`
}

func (p *Problem) TableName() string {
	return "problem"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	db := DB.Model(new(Problem)).
		Preload("ProblemCategories").
		Preload("ProblemCategories.Category").
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if len(categoryIdentity) > 0 {
		fmt.Println("categoryIdentity:", categoryIdentity)
		db.Joins("RIGHT JOIN problem_category ON problem_category.problem_id = problem.id").
			Where("problem_category.category_id = (select id from category as cg where cg.identity = ?)", categoryIdentity)
	}

	return db
}

func FindProblemByIdentity(identity string) (*Problem, error) {
	var problem Problem
	err := DB.Model(new(Problem)).
		Preload("ProblemCategories").
		Preload("ProblemCategories.Category").
		Where("identity = ?", identity).
		First(&problem).Error
	return &problem, err
}
