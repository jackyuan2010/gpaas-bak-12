package model

type Enterprise struct {
	Model
	Name    string `json:"name" gorm:"size:500;not null"`
	Address string `json:"address" gorm:"size:500;"`
}

func (entity *Enterprise) TableName() string {
	return "sys_enterprise"
}
