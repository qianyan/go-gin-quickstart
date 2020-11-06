package infra

import (
	"github.com/jinzhu/gorm"
	"github.com/qianyan/go-gin-quickstart/infra/logger"
	"go.uber.org/zap"
	//import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQL struct {
	db *gorm.DB
}

var (
	log = logger.MustGetLogger("MySQL")
)

func (mysql *MySQL) OpenDB(uri string) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	mysql.db = db
}

func (mysql *MySQL) CloseDB() {
	err := mysql.db.Close()
	log.Error("Close MySQL", zap.Error(err))
}

func (mysql *MySQL) Get() *gorm.DB {
	return mysql.db
}
