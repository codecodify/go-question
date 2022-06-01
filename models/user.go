package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);uniqueIndex;" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Email    string `gorm:"column:email;type:varchar(100)" json:"email"`
	Phone    string `gorm:"column:phone;type:varchar(100)" json:"phone"`
}

func (u *User) TableName() string {
	return "user"
}

func FindUserByIdentity(identity string) (*User, error) {
	var user User
	err := DB.Model(new(User)).Where("identity = ?", identity).First(&user).Error
	return &user, err
}

func Login(name, password string) (*User, error) {
	var user User
	err := DB.Model(new(User)).Where("name = ? AND password = ?", name, password).First(&user).Error
	return &user, err
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func GetUserByName(name string) bool {
	var user User
	err := DB.Model(new(User)).Where("name = ?", name).First(&user).Error
	return err == nil
}

func GetUserByEmail(email string) bool {
	var user User
	err := DB.Model(new(User)).Where("email = ?", email).First(&user).Error
	return err == nil
}
