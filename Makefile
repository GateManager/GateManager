build:
	go build -o bin/GateManager ./cmd/GateManager/main.go

run: build
	./bin/GateManager

test:
	go test -v ./... -count=1

tidy:
	go mod tidy

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down