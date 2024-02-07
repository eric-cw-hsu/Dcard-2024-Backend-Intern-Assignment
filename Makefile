migrate-up: 
	migrate -path ./migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose up

migrate-down:
	migrate -path ./migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose down

migrate-create: 
	migrate create -ext sql -dir ./migrations -seq $(name)
