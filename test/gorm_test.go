package test

import (
	"fmt"
	"github.com/codecodify/go-question/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// 测试gorm连接
func TestGormConnect(t *testing.T) {
	dsn := "gin_pj:iA6hrNdNyYp2tNTb@tcp(42.51.5.91:3306)/gin_pj?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.Problem, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}
}
