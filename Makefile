
all: build

build: 
	go build 

run: build
	./aurgit --config=config/config.json

clean: 
	go clean -i

test: build
	go test

update-deps: 
	go get -u



