BINARY_NAME=tabnews_bot
ZIP_FILE_NAME=tabnews_bot

all: build

build: clean
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) main.go

zip:
	zip -r $(ZIP_FILE_NAME).zip bin/$(BINARY_NAME)

aws_lambda: clean build zip

build_tg_cli:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/telegram_linux script/telegram/main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/telegram_windows.exe script/telegram/main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/telegram_mac script/telegram/main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o bin/telegram_mac_arm script/telegram/main.go

test:
	go test -v ./...

clean:
	rm -rf bin
	rm -f $(ZIP_FILE_NAME).zip