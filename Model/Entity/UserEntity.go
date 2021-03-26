package Entity

type UserEntity struct {
	Id        string `gorm:"primaryKey"`
	Age       int
	Password  string
	LoginName string
	NickName  string
}

func (UserEntity) TableName() string {
	return "User"
}
