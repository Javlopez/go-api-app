.PHONY: run test test-with-coverage coverage

run:
	go run main.go

test:
	go test ./... -v

test-with-coverage:
	go test ./... -v -coverprofile coverage/cover.out

coverage: test-with-coverage
	go tool cover -html=coverage/cover.out