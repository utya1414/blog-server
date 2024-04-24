help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

start:
	docker start postgres16

bash:
	docker exec -it postgres16 bash

psql:
	docker exec -it postgres16 psql --username=root --dbname=blog

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root blog

migrateup:
	migrate -path infrastructure/db/migrations -database "postgresql://root:password@localhost:5434/blog?sslmode=disable" -verbose up

migratedown:
	migrate -path infrastructure/db/migrations -database "postgresql://root:password@localhost:5434/blog?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
	
server:
	go run main.go

mockdb:
	mockgen -source infrastructure/db/store.go -destination infrastructure/db/mock/store.go

mock:
	scripts/mock_gen.sh
PHONY:
	help bash psql start createdb migrateup migratedown test server mock