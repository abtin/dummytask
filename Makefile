############
# Commands
############
help:
	@echo Usage:
	@echo make clean	- to delete ${APP_NAME} executable file
	@echo make linux	- to build ${APP_NAME} executable for linux.
	@echo make docker	- to build docker image for ${APP_NAME}

APP_NAME := dummytask

clean:
	@rm ${APP_NAME}
linux:
	@GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o ${APP_NAME}

docker:
	@docker build  -t abtin/dummytask:0.1 .

.PHONY: clean, build, docker-build
