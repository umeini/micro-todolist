package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWordDigest string
}

const (
	PassWordCost = 12 //密码加密难度
)

// 设置加密密码
func (user *User)SetPassWord(password string )error  {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// 检验密码
func (user *User)CheckPassWord(password string)bool  {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}