GOCMD=go
GOCLEAN=$(GOCMD) clean
BINARY_NAME=go-cutup

.PHONY: all build

build:
	$(GOCLEAN)
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/linux/$(BINARY_NAME) -v main.go

clean:
	$(GOCLEAN)
	rm -f ./build/linux/$(BINARY_NAME)

run:
	go build -o ./build/linux/$(BINARY_NAME) -v main.go
	./build/linux/$(BINARY_NAME)

