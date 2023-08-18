package main

import (
	"fmt"
	"kafka_events/internal/routes"
	"kafka_events/pkg/database/influxdb"
	"log"
)

func main() {
	db := influxdb.Init()
	defer db.Close()

	router := routes.SetupLiveEventRoutes(db)
	fmt.Println("Server started on port 4001")
	log.Fatal(router.Run(":4001"))
}
