.DEFAULT_GOAL := build
.PHONY : build

IMAGE_NAME:=susu-wechat-bot

build:
	docker build -f build/Dockerfile -t ${IMAGE_NAME} .
test:
	docker build -f build/Dockerfile -t ${IMAGE_NAME}:test .
