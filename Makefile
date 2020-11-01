.PHONY: all

all: build run

build:
	@echo "===> $@"
	@go build -i -o bin/to-do-ctl -v $(REPO_PATH)/Microsoft-To-Do-API/cmd/task-ctl

run:
	@echo "===> $@"
	@./bin/to-do-ctl