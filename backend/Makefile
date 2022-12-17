migrate_up:
	migrate -path ./internal/database/mysql/migrations -database "mysql://root:Math@2109@tcp(localhost:3306)/bdgolang?charset=utf8mb4&parseTime=true&loc=Local" -verbose up

migrate_down:
	migrate -path ./internal/database/mysql/migrations -database "mysql://root:Math@2109@tcp(localhost:3306)/bdgolang?charset=utf8mb4&parseTime=true&loc=Local" -verbose down

update_swagger:
	./bin/update_swagger.sh