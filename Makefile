build:
	go build -o practical-go cmd/hello/main.go

run:
	./practical-go

clean:
	rm practical-go

check: build run clean