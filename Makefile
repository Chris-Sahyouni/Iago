# Note: ./... means recursively match all go packages in/under the current directory

CC = gcc


all: build

build:
	go build -o bin/iago ./...

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./... -v

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: clean_binaries
clean_binaries:
	rm -rf test_programs/bin/*

# add some logic to compile all c src files in test_programs
binaries:
	$(CC) -o test_programs/bin/square64 test_programs/square.c
	$(CC) -o test_programs/bin/square32 test_programs/square.c -m32