build:
	go build -o bin/app

run:
	./bin/app

test:
	grc go test ./... -count=1
