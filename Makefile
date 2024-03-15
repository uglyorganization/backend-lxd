# Define variables
IMAGE_NAME := my-backend-image
CONTAINER_NAME := my-backend-container

# Build Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run Docker container
run:
	docker run -d -p 8080:8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

# Health check
health:
	curl http://localhost:8080/health

# Stop Docker container
stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# Target to build, run, check health, and stop
test: build run health stop

# Phony targets to prevent conflicts with file names
.PHONY: build run health stop test
