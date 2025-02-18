package server

import "fmt"

func Start(port string) error {
	engine, err := initEngine()
	if err != nil {
		return err
	}

	// Initialise routes
	initRoutes(engine)
	return engine.Run(fmt.Sprintf(":%s", port))
}
