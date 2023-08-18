package routes

import (
	"kafka_events/internal/handlers"
	"kafka_events/pkg/database/influxdb"

	"github.com/gin-gonic/gin"
)

func SetupNewEventRoutes() *gin.Engine {
	router := gin.Default()
	eventHandler := handlers.NewEventHandler()

	router.POST("/event", eventHandler.CreateEventHandler)

	return router
}

func SetupLiveEventRoutes(db *influxdb.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/live-events", func(c *gin.Context) {
		handlers.HandleLiveEventsWebSocket(c, db)
	})
	return router
}
