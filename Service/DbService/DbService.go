package DbService

import (
	"fmt"
	"golang-blog/Model/Entity"
	"golang-blog/Service/ConfigService"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var (
		err error
	)
	dbConfig := ConfigService.GetAppConfig("database")
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.USER,
		dbConfig.PASSWORD,
		dbConfig.HOST, dbConfig.NAME)
	Db, err = gorm.Open(mysql.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动生成表结构
	dbErr := Db.AutoMigrate(&Entity.UserEntity{})
	if dbErr != nil {
		println(err)
	}
}
