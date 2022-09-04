package model

type Model struct{
	Id string `gorm:"primaryKey" json:"id"`
	Create_Time int64 `gorm:"autoCreateTime:nano" json:"create_time"`
	Update_Time int64 `gorm:"autoUpdateTime:nano" json:"update_time"`
}