# Note: ./.... means recursively match all go packages in/under the current directory

CC = gcc


all: build

build:
	go build -o bin/iago ./....

run:
	go run main.go

test:
	go test ./....

# add some logic to compile all c src files in test_programs
binaries:
	$(CC) -o test_programs/square test_programs/square.c