# 定义变量
HOMEDIR = /home/work/comment-service
SERVICENAME = comment-service

.PHONY: build run clean

build:
	@echo "Building ${SERVICENAME}..."
	@docker build -t ${SERVICENAME} .

run:
	@echo "Running ${SERVICENAME} container..."
	@docker run -d --name ${SERVICENAME}-container ${SERVICENAME}

clean:
	@echo "Cleaning up..."
	@docker stop ${SERVICENAME}-container
	@docker rm ${SERVICENAME}-container

