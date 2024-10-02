package server

import (
	"fmt"
	"log"
)

func Start(port string) error {
	engine, err := initEngine()
	if err != nil {
		log.Fatalf("Failed to initialize router: %v", err)
	}

	// Initialise routes
	initRoutes(engine)
	return engine.Run(fmt.Sprintf(":%s", port))
}
