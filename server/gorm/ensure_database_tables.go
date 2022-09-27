package gorm

import (
	"fmt"
	model "github.com/jackyuan2010/gpaas/server/gorm/model"
	"gorm.io/gorm"
)

func MigrateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Enterprise{})
	db.AutoMigrate(&model.EnterpriseUser{})
	db.AutoMigrate(&model.DatabaseInfo{})

	object := model.User{Model: model.Model{Id: "1"}, Name: "admin", Password: "123456", Mobile: "12345"}

	fmt.Println(object)
	// db.Create(&object)

	fmt.Println("Get data from db")
	db.First(&object, "id=?", "1")
	fmt.Println(object)
	fmt.Println(object.Id)

}
