package main

import (
	"dcard-2024-backend-intern-assignment/databases"
	"dcard-2024-backend-intern-assignment/server"
	"dcard-2024-backend-intern-assignment/utils"
	"log"
)

func main() {
	config := utils.LoadConfig()
	db := databases.NewMysqlDatabase(databases.DatabaseConfigs{
		Host:         config.DB_HOST,
		Port:         config.DB_PORT,
		User:         config.DB_USER,
		Password:     config.DB_PASSWORD,
		DatabaseName: config.DB_NAME,
	})

	server := server.NewServer(config, db)
	server.InitRouter()

	server.StartDatabaseConnection()
	defer server.CloseDatabaseConnection()

	log.Fatal(server.Serve())
}
