package UserService

import (
	"golang-blog/Model/Entity"
	"golang-blog/Service/DbService"
)

func ByLoginNameGetUser(loginName string) Entity.UserEntity {
	var (
		user Entity.UserEntity
	)
	DbService.Db.Where(&Entity.UserEntity{LoginName: loginName}).First(&user)
	return user
}

func SignIn(user Entity.UserEntity) string {
	user.Id = "123"
	DbService.Db.Create(&user)
	return user.Id
}

func Update(user Entity.UserEntity) {
	DbService.Db.Model(&Entity.UserEntity{}).Where(&Entity.UserEntity{Id: user.Id}).Updates(user)
}

func Delete(userId string) {
	DbService.Db.Delete(&Entity.UserEntity{}, userId)
}
