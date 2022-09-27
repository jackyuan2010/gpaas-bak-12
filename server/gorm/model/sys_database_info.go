package model

type DatabaseInfo struct {
	Model
	EnterpriseId string `json:"enterprise_id" gorm:"enterprise_id" `
	Host         string `json:"host" gorm:"host" `
	Port         string `json:"port" gorm:"port" `
	DbName       string `json:"dbname" gorm:"dbname"`
}

func (entity *DatabaseInfo) TableName() string {
	return "sys_database_info"
}
