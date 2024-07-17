build:
	go build -o ./bin/raven.exe ./cmd/main.go

run:
	make build
	./bin/raven.exe

gqlgen:
	go run github.com/99designs/gqlgen@v0.17.49

.PHONY: run
.PHONY: build
.PHONY: gqlgen