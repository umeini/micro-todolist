package model

//迁移   将定义模型映射到数据库上面

func migration(){
	DB.Set("gorm:table_options","charset=utf8mb4").AutoMigrate(&User{})
}
