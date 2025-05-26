# Docker 相关变量
DOCKER_IMAGE_NAME = ccr.ccs.tencentyun.com/cloudmonitor/wework-weixin-bot
VERSION ?= latest

# 构建 Docker 镜像
.PHONY: docker-build
build-image:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) -f Dockerfile .


build-linux-image:
	DOCKER_BUILDKIT=1 DOCKER_DEFAULT_PLATFORM=linux/amd64 docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) -f Dockerfile .

push-image:
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)




