package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string   `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	ProblemIdentity string   `gorm:"column:problem_identity;type:varchar(36);index;" json:"problem_identity"`
	Problem         *Problem `gorm:"foreignKey:identity;references:problem_identity" json:"problem"`
	UserIdentity    string   `gorm:"column:user_identity;type:varchar(36);index;" json:"user_identity"`
	User            *User    `gorm:"foreignKey:identity;references:user_identity" json:"user"`
	Path            string   `gorm:"column:path;type:varchar(150)" json:"path"`
	Status          int      `gorm:"column:status;type:tinyint;" json:"status"`
}

func (s *Submit) TableName() string {
	return "submit"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm.DB {
	db := DB.Model(new(Submit)).
		Preload("Problem", func(db *gorm.DB) *gorm.DB {
			return db.Omit("content")
		}).
		Preload("User").
		Preload("Problem.ProblemCategories")
	if problemIdentity != "" {
		db.Where("problem_identity = ?", problemIdentity)
	}
	if userIdentity != "" {
		db.Where("user_identity = ?", userIdentity)
	}
	if status != -1 {
		db.Where("status = ?", status)
	}
	return db
}
