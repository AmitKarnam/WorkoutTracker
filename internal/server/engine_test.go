package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitEngine(t *testing.T) {
	engine, err := initEngine()

	// Test that the engine is created successfully
	assert.NoError(t, err)
	assert.NotNil(t, engine)

	// Test that the engine is in ReleaseMode
	assert.Equal(t, gin.ReleaseMode, gin.Mode())

	// Test that the engine has the expected middleware
	handlers := engine.Handlers
	assert.Len(t, handlers, 2)
	assert.IsType(t, gin.LoggerWithConfig(gin.LoggerConfig{}), handlers[0])
	assert.IsType(t, gin.RecoveryWithWriter(nil), handlers[1])

	// Test that DebugPrintRouteFunc is set
	assert.NotNil(t, gin.DebugPrintRouteFunc)
}
