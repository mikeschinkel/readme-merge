# Define variables
DOCKER_IMAGE = mikeschinkel/readme-merge
DOCKER_TAG = latest

# Default target
.PHONY: all
all: help

# Build the Docker image
.PHONY: build-docker
build-docker:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Publish the Docker image to DockerHub
.PHONY: publish-docker
publish-docker: build-docker
	@echo "Publishing Docker image to DockerHub..."
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

# Clean up Docker images
.PHONY: clean-docker
clean-docker:
	@echo "Cleaning up Docker images..."
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true

# Help message
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make publish-docker  - Build and publish the Docker image to DockerHub"
	@echo "  make clean-docker    - Clean up Docker images"
	@echo "  make help            - Show this help message"
