SRC_PATH?=./internal
APP_NAME?=spacelight

build:
	go build -o $(APP_NAME) $(SRC_PATH)