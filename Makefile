IMAGE_NAME = ca-simulator-image
CONTAINER_NAME = simulator-app
PORT = 8080

DOCKERFILE = Dockerfile

.PHONY: build
build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

.PHONY: run
run:
	@echo "Running Docker container..."
	docker run -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

.PHONY: run-detached
run-detached:
	@echo "Running Docker container in detached mode..."
	docker run -d -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

.PHONY: stop
stop:
	@echo "Stopping and removing Docker container..."
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)

#
