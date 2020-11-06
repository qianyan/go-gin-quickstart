package infra

import "github.com/jinzhu/gorm"

type Database interface {
	OpenDB(uri string)
	CloseDB()
	Get() *gorm.DB // TODO: remove later because of violating DI principle
}
