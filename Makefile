# Note: ./.... means recursively match all go packages in/under the current directory

CC = gcc


all: build

build:
	go build -o bin/iago ./....

run:
	go run main.go

test:
	go test ./....

clean:
	rm -rf bin/*

clean_binaries:
	rm -rf test_programs/bin/*

# add some logic to compile all c src files in test_programs
binaries:
	$(CC) -o test_programs/bin/square test_programs/square.c