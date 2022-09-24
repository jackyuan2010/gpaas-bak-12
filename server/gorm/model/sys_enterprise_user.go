package model

type EnterpriseUser struct {
	EnterpriseId string `gorm:"column:sys_enterprise_id"`
	UserId       string `gorm:"column:sys_user_id"`
}

func (entity *EnterpriseUser) TableName() string {
	return "sys_enterprise_user"
}
