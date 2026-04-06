APP_NAME=sysmon

run:
	go run .

build:
	go build -o $(APP_NAME)

build-win:
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)

clean:
	rm -f $(APP_NAME) $(APP_NAME).exe