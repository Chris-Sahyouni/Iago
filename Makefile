CC = gcc
ARMCC = arm-linux-gnueabi-gcc

.PHONY: all
all: build

build:
	go build -o bin/iago ./src/main.go

.PHONY: run
run:
	go run ./src/main.go

.PHONY: test
test:
	go test ./src/... -v

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: clean_binaries
clean_binaries:
	rm -rf test_programs/bin/*

binaries:
	$(CC) -o test_programs/bin/square64 test_programs/square.c
	$(CC) -o test_programs/bin/square32 test_programs/square.c -m32
	$(ARMCC) -o test_programs/bin/squareARM test_programs/square.c
	$(CC) -o test_programs/bin/vuln64 test_programs/vuln.c -fno-stack-protector
	$(CC) -o test_programs/bin/vuln32 test_programs/vuln.c -m32 -fno-stack-protector