package model

type DatabaseInfo struct {
	EnterpriseId string `gorm:"enterprise_id" json:"enterprise_id"`
	Host         string `gorm:"host" json:"host"`
	Port         string `gorm:"port" json:"port"`
	DbName       string `gorm:"dbname" json:"dbname" yaml:"dbname"`
}

func (entity *DatabaseInfo) TableName() string {
	return "sys_database_info"
}
