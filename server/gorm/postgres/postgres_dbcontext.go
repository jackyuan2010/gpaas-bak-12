package gorm

import (
	"strings"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	gpaasgorm "github.com/jackyuan2010/gpaas/server/gorm"
)

type PostgresDbContext struct {
	DbConfig *gpaasgorm.DbConfig
}

func NewDbContext(dbconfig *gpaasgorm.DbConfig) PostgresDbContext {
	dbcontext := PostgresDbContext{DbConfig : dbconfig}
	return dbcontext
}

func (dc *PostgresDbContext) DSN() string {
	var sb strings.Builder
	sb.WriteString("host=")
	sb.WriteString(dc.DbConfig.Host)

	sb.WriteString(" user=")
	sb.WriteString(dc.DbConfig.Username)

	sb.WriteString(" password=")
	sb.WriteString(dc.DbConfig.Password)

	sb.WriteString(" dbname=")
	sb.WriteString(dc.DbConfig.DbName)

	sb.WriteString(" port=")
	sb.WriteString(dc.DbConfig.Port)

	sb.WriteString(" ")
	sb.WriteString(dc.DbConfig.Config)
	return sb.String()
}

func (dc *PostgresDbContext) GetDb() *gorm.DB {
	pgsqlconfig := postgres.Config{
		DSN:                  dc.DSN(),
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlconfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dc.DbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dc.DbConfig.MaxOpenConns)
		return db
	}
}