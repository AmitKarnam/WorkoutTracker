package mysql

import (
	"github.com/AmitKarnam/WorkoutTracker/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Defines the implementing type DB, which in this case is mysql
type mysqlDB struct {
	path   string
	DbConn *gorm.DB
}

// Asserts that the implementing type DB implements the GetConnection method
var _ database.Database = &mysqlDB{}

// Singleton Pattern is implemented here.
var dbInstance *mysqlDB

// Assigns the singleton reference variable the mysql DB instance
func DBGetConnection(path string) (*gorm.DB, error) {
	var err error
	if dbInstance == nil {
		dbInstance = new(path)
		dbInstance.DbConn, err = dbInstance.GetConnection()
		if err != nil {
			return nil, err
		}
	}

	return dbInstance.DbConn, nil
}

// Creation of mysql db instance
func new(path string) *mysqlDB {
	return &mysqlDB{
		path: path,
	}
}

// Implementation of the getConnection method
func (s *mysqlDB) GetConnection() (*gorm.DB, error) {
	return (gorm.Open(mysql.Open(s.path), &gorm.Config{
		SkipDefaultTransaction: true,
	}))
}
