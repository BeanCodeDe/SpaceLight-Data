SRC_PATH?=./internal
APP_NAME?=spacelight

build:
	go build -buildvcs=false -o $(APP_NAME) $(SRC_PATH)