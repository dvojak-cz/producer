DOCKER_IMAGE=producer
DOCKER_IMAGE_TAG=latest

.PHONY: build

docker:
	docker build \
		-f build/Dockerfile \
		-t $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) \
		.

run:
	ENVIRONMENT_FILE="./config/.env" go run cmd/producer/producer.go