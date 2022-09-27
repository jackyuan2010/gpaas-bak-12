package model

type User struct {
	Model
	Name        string       `json:"name" gorm:"size:100;not null"`
	Mobile      string       `json:"mobile" gorm:"size:100;not null;unique"`
	Password    string       `json:"password" gorm:"size:100;not null"`
	DenyLogin   bool         `json:"deny_login" gorm:"default:false"`
	Status      int          `json:"status" gorm:"default:1"`
	Enterprises []Enterprise `json:"authorities" gorm:"many2many:sys_enterprise_user;"`
}

func (entity *User) TableName() string {
	return "sys_user"
}
