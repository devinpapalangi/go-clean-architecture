RUN-DEV:
	air --build.cmd "go build -o bin/api cmd/api/main.go" --build.bin "./bin/api"

BUILD-DOCS:
	swag init -g cmd/api/main.go --parseDependency --parseInternal

RUN-TESTS:
	go test ./...

FORMAT-DOCS:
	swag fmt

CHECK-MIGRATION:
	atlas schema inspect --env gorm --url "env://src"

CREATE-MIGRATION:
	atlas migrate diff --env gorm

APPLY-MIGRATION:
	atlas migrate apply --url "postgres://postgres:postgres@localhost:5432/go_refreshments?sslmode=disable"


