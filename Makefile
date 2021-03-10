REPOSITORY?=
TAG?=

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o chart-exporter cmd/chart-exporter/*.go 
	docker build -t ${REPOSITORY}:${TAG} .
	rm chart-exporter

run:
	go run cmd/chart-exporter/main.go
