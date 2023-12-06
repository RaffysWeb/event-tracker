package influxdb

import (
	"fmt"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/joho/godotenv"
)

type DB struct {
	influxClient influxdb2.Client
	org          string
	token        string
	url          string
}

func Init() *DB {
	fmt.Println("init")
	err := godotenv.Load("../../config/influxdb/influxdb.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("INFLUXDB_INIT_URl")
	token := os.Getenv("INFLUXDB_INIT_ADMIN_TOKEN")

	influxClient := influxdb2.NewClient(url, token)
	return &DB{
		influxClient: influxClient,
		org:          os.Getenv("INFLUXDB_INIT_ORG"),
	}
}

func (db *DB) Close() {
	db.influxClient.Close()
}

func (db *DB) GetClient() influxdb2.Client {
	return db.influxClient
}

func (db *DB) QueryAPI() api.QueryAPI {
	return db.influxClient.QueryAPI(db.org)
}
