

.EXPORT_ALL_VARIABLES:

TMPDIR = ./build/tmp

default: build/irck

setuplocal:
	mkdir -p build/tmp

cleanuplocal:
	rm -rf build

build/irck: setuplocal irck/main.go
	go build -o build/irck irck/main.go

clean: setuplocal
	go clean # we literally need the tmpdirs to clean up things
	rm -rf build

test: setuplocal
	go test

