BINARY_NAME=tabnews_bot
ZIP_FILE_NAME=tabnews_bot

all: build test

build:
	@go get -v ./...
	@go build -o bin/$(BINARY_NAME) -v

zip:
	@zip -r $(ZIP_FILE_NAME).zip bin/$(BINARY_NAME) config.json

publish_lambda:
	@aws lambda update-function-code --function-name tabnews_bot --zip-file fileb://$(ZIP_FILE_NAME).zip

test:
	@go test -v ./...

clean:
	@rm -f bin/$(BINARY_NAME)