package model

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB //

//Database 在中间件中初始化mysql连接
func DataBase(connString string)  {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release"{
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)//设置最大连接池，空闲
	db.DB().SetMaxOpenConns(100)//打开
	db.DB().SetConnMaxLifetime(time.Second*30)
	DB = db
	migration()
}
