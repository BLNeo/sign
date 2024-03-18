package models

import (
	"gorm.io/gorm"
	"sign/tool/log"
	"sign/tool/mysql"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDb(dbEngine *gorm.DB) {
	once.Do(func() {
		db = dbEngine
	})
}

func GetDb() *gorm.DB {
	return db
}

type Model struct {
	ID        int64 `gorm:"primaryKey;type:int4"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateTable 表同步
func CreateTable() {
	tables := []interface{}{
		new(User),
	}
	err := mysql.AutoMigrate(db, tables...)
	if err != nil {
		log.Logger.Error("mysql AutoMigrate err:" + err.Error())
		panic(err)
	}
}
