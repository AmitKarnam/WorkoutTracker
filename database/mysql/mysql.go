package mysql

import (
	"github.com/AmitKarnam/WorkoutTracker/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Defines the implementing type DB, which in this case is mysql
type MySQLDB struct {
	path string
}

// Asserts that the implementing type DB implements the GetConnection method
var _ database.Database = &MySQLDB{}

// Singleton Pattern is implemented using this reference variable, which is used by other modules to access this singleton reference.
var DB database.Database

// Initialise DB
func NewMySQLInit(path string) {
	DB = new(path)
}

// Creation of mysql db instance
func new(path string) *MySQLDB {
	return &MySQLDB{
		path: path,
	}
}

// Implementation of the getConnection method
func (s *MySQLDB) GetConnection() (*gorm.DB, error) {
	return (gorm.Open(mysql.Open(s.path), &gorm.Config{
		SkipDefaultTransaction: true,
	}))
}
