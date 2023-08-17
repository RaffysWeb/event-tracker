package database

// Consider if this useful or not
type Database interface {
	Init()
	Close()
	GetSession() interface{}
}
