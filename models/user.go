package models

type User struct {
	Model
	Name     string `gorm:"column:name;type:varchar(128)" json:"name"`
	Phone    string `gorm:"column:phone;type:varchar(64)" json:"phone"`
	Password string `gorm:"column:password;type:varchar(128)" json:"password"`
	PassSalt string `gorm:"column:pass_salt;type:varchar(36)" json:"pass_salt"`
	Email    string `gorm:"column:email;type:varchar(50);comment:邮箱地址" json:"email"`
	Avatar   string `gorm:"column:avatar;comment:头像url" json:"avatar"`
	Gender   string `gorm:"column:gender;type:varchar(1);comment:性别" json:"gender"` // m-男 w-女 n-保密
	Nickname string `gorm:"column:nickname;type:varchar(128);comment:昵称" json:"nickname"`
}

func (u *User) User() string {
	return "user"
}
