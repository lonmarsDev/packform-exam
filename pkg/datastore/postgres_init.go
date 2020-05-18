package datastore

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Initialize postgres driver
)

// DBHandle creates and returns a database handle
var (
	PostresDb  *gorm.DB
	onceConfig sync.Once
	onceDB     sync.Once
	DbHandle   *gorm.DB
	err        error
)

func DbPgInit(dbusername, dbpassword, dbhost, dbname string, dbport int) error {
	onceDB.Do(func() {
		connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbhost, dbport, dbusername, dbname, dbpassword)
		PostresDb, err = gorm.Open("postgres", connectionString)
	})

	return err
}
