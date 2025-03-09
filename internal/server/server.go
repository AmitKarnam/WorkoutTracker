package server

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
)

func Start(port string) error {
	engine, err := initEngine()
	if err != nil {
		return err
	}
	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		return err
	}
	// Initialise routes
	initRoutes(engine, dbConn)
	return engine.Run(fmt.Sprintf(":%s", port))
}
