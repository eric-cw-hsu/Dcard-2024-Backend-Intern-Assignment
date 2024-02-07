package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Implement the BaseDatabase interface
type MysqlDatabase struct {
	config DatabaseConfigs
	pool   *sql.DB
}

func NewMysqlDatabase(config DatabaseConfigs) *MysqlDatabase {
	return &MysqlDatabase{config: config}
}

func (db *MysqlDatabase) Connect() {
	dsn := db.GetConnectionString()
	pool, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)
	pool.SetConnMaxLifetime(0)

	db.pool = pool
}

func (db *MysqlDatabase) Close() {
	err := db.pool.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (db *MysqlDatabase) GetConnectionString() string {
	// return the connection string for the mysql database
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", db.config.User, db.config.Password, db.config.Host, db.config.Port, db.config.DatabaseName)
}
