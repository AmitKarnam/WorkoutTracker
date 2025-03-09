package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func initEngine() (*gin.Engine, error) {

	gin.SetMode(gin.ReleaseMode)

	// Initialise New gin engine
	engine := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%v %v %v %v", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	return engine, nil
}
