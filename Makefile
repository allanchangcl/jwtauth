build:
	go build -o bin/app

run:
	./bin/app

test:
	grc go test -v ./... -count=1
