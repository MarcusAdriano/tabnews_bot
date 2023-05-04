BINARY_NAME=tabnews_bot
ZIP_FILE_NAME=tabnews_bot

all: build

build: clean
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) main.go

zip:
	zip -r $(ZIP_FILE_NAME).zip bin/$(BINARY_NAME)

aws_lambda: clean build zip

test:
	go test -v ./...

clean:
	rm -rf bin
	rm -f $(ZIP_FILE_NAME).zip