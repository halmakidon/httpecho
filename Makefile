IMAGE_NAME := 788213621982.dkr.ecr.ap-northeast-1.amazonaws.com/httpecho
TAG := 1.0.0

build:
	DOCKER_BUILDKIT=1 docker build -t $(IMAGE_NAME):$(TAG) .

run:
	docker run -it --rm -p 8080:8080 $(IMAGE_NAME):$(TAG)

push:
	docker push $(IMAGE_NAME):$(TAG)