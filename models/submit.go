package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);index;" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);index;" json:"user_identity"`
	Path            string `gorm:"column:path;type:varchar(150)" json:"path"`
	Status          int    `gorm:"column:status;type:tinyint;" json:"status"`
}

func (s *Submit) TableName() string {
	return "submit"
}
