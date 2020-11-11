package infra

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
	"os"
)

type Sqlite struct {
	db  *gorm.DB
	uri string
}

var (
	DB *gorm.DB
)

func (s *Sqlite) OpenDB(uri string) {
	db, err := gorm.Open("sqlite3", uri)
	if err != nil {
		fmt.Errorf("db err: %v", err)
	}
	db.DB().SetMaxIdleConns(10)
	s.db = db
	s.uri = uri
	DB = db
}

func (s *Sqlite) CloseDB() {
	err := s.db.Close()
	err = os.Remove(s.uri)
	fmt.Println("Close Sqlite", zap.Error(err))
}

func (s *Sqlite) Get() *gorm.DB {
	return s.db
}

func (s *Sqlite) setMaxIdleConns(n int) {
	s.db.DB().SetMaxIdleConns(n)
}

func (s *Sqlite) logMode(b bool) {
	s.db.LogMode(b)
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *Sqlite {
	testdb := &Sqlite{}
	testdb.OpenDB("./../gorm_test.db")
	testdb.setMaxIdleConns(3)
	testdb.logMode(true)
	return testdb
}

// Delete the database after running testing cases.
func TestDBFree(test_db *Sqlite) {
	test_db.CloseDB()
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
