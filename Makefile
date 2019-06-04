

default: irck.go
	mkdir -p build
	go build -o build/irck irck.go

clean:
	rm -rf build
	go clean
