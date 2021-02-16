build:
	GOFLAGS=-mod=vendor go build -o build/amicaldo main.go

clean:
	rm -rf build

run:
	GOFLAGS=-mod=vendor go run main.go
