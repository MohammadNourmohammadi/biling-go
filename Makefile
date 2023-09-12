version := v0.0.1
json_path := ./json-file
BILING_PORT := 8000
run:
	json_path=$(json_path) BILING_PORT=$(BILING_PORT) go run ./cmd/biling/main.go
test:
	go test ./... -v
build-image:
	docker build . -t biling:$(version)
compose-up:
	docker-compose up
install-requirements:
	cat requirements | xargs go get -u

