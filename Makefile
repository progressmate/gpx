install-tools:
	go install aqwari.net/xml/cmd/xsdgen

generate:
	go generate ./...

test:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
