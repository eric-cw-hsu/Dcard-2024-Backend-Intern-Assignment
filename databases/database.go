package databases

import "database/sql"

type DatabaseConfigs struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type BaseDatabase interface {
	GetPool() *sql.DB
	Connect()
	Close()
	GetConnectionString() string
}
