package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RToken   string `json:"r_token"`
}

type UserToken struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
	RToken string `json:"r_token"`
}

func CreateAuth(user User, rToken string) UserToken {
	db := Connect()
	uToken := UserToken{User: user, RToken: rToken}
	db.Create(&uToken)
	return uToken
}

func GetUser(id uint) User {
	db := Connect()
	user := User{ID: id}
	db.Select("Name", "Email", "RToken").Find(&user)
	return user
}

func GetUserByEmail(email string) User {
	db := Connect()
	user := User{Email: email}
	db.Find(&user)
	return user
}
func UpdateUser(newUser User) User {
	db := Connect()
	oldUser := User{ID: newUser.ID}
	db.Model(&oldUser).Updates(newUser)
	return oldUser
}

func DeleteUser(id uint) uint {
	db := Connect()
	user := User{ID: id}
	db.Delete(&user)
	return id
}

func CreateUser(user User) User {
	db := Connect()
	db.Create(&user)
	return user
}
