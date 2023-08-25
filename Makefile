dev:
	@go run main.go

build:
	@go build -tags netgo -ldflags '-s -w' -o app

#Docker Related
REGION := asia-southeast1
IMAGE_NAME := amaz-corp/ac-be/ac-be
TAG_VERSION := 0.0.3

build-docker:
	#[[--platform linux/amd64 required for cloud run]]
	@docker build --platform linux/amd64 -t ${REGION}-docker.pkg.dev/${IMAGE_NAME} .

docker-tag:
	@docker tag ${REGION}-docker.pkg.dev/${IMAGE_NAME} ${REGION}-docker.pkg.dev/${IMAGE_NAME}:${TAG_VERSION}

docker-push:
	@docker push ${REGION}-docker.pkg.dev/${IMAGE_NAME}:${TAG_VERSION}

docker-history:
	@docker image history ${REGION}-docker.pkg.dev/${IMAGE_NAME}:${TAG_VERSION}

CONTAINER_NAME := suspicious_black

docker-inspect-fs:
	#[[check filesystem inside the image in container name ${CONTAINER_NAME} and export it to image-fs.txt]]
	@docker export ${CONTAINER_NAME} | tar t > image-fs.txt

#Test logic
test:
	@go -short ./...

#Test with mock call to the app
test-mock:
	@go test ./internal/app/route

clean-test-cache:
	@go clean -testcache
