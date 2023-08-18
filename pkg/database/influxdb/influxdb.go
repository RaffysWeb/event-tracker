package influxdb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// TODO: Move secret to .env file and config out of here
const (
	url   = "http://localhost:8086"
	token = "iIXvA0k0tCNpZM5yOaT0-mHh4GUP8OEe2Fylx6pXziHXrnUUp7HV8nfotLOIP310zd9vh9m_y9NfBkM_kUd_3w=="
	org   = "my-org"
)

type DB struct {
	influxClient influxdb2.Client
}

func Init() *DB {
	influxClient := influxdb2.NewClient(url, token)
	return &DB{influxClient: influxClient}
}

func (db *DB) Close() {
	db.influxClient.Close()
}

func (db *DB) GetClient() influxdb2.Client {
	return db.influxClient
}

func (db *DB) QueryAPI() api.QueryAPI {
	return db.influxClient.QueryAPI(org)
}
