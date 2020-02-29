.PHONY: setup run test test-with-coverage coverage doc

IMG=lana-app
PORT?=8080

setup:
	mkdir -p storage/

run:
	go run main.go

test:
	go test ./... -v

test-with-coverage:
	go test ./... -v -coverprofile coverage/cover.out

coverage: test-with-coverage
	go tool cover -html=coverage/cover.out

doc:
	godoc -http=:8081

docker-build:
	docker build --rm --build-arg port=${PORT} --no-cache -t ${IMG} .

docker-run:
	docker run -d -e ENV_PORT=${PORT} --rm -p ${PORT}:${PORT} ${IMG}:latest	