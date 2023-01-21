lint:
	go mod tidy || exit
	golangci-lint run

test:
	go mod tidy || exit
	go test -v -covermode=count -coverprofile=coverage_unit.out -coverpkg=./... -gcflags="-N -l" ./...
