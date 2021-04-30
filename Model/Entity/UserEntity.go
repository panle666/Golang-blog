package Entity

type UserEntity struct {
	Id        string `gorm:"primaryKey"`
	Age       int
	Password  string
	LoginName string
	NickName  string
}

// 自定义表名称
func (UserEntity) TableName() string {
	return "User"
}
