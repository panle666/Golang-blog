package DbService

import (
	"golang-blog/Model/Entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var (
		err error
	)
	Db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动生成表结构
	dbErr := Db.AutoMigrate(&Entity.UserEntity{})
	if dbErr != nil {
		println(err)
	}
}
