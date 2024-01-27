# Install sqlc and sql-migrate
setup:
	@go get
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@$(MAKE) gen-sql
	@go install github.com/rubenv/sql-migrate/...@latest

gen-sql:
	@sqlc generate

gen-sql-win:
	@docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate .\sqlc.yaml

dev:
	@go run main.go

build:
	@go build -tags netgo -ldflags '-s -w' -o amaz-corp-app

#Docker Related
REGION := asia-southeast1
IMAGE_NAME := amaz-corp/ac-be/ac-be
TAG_VERSION := 0.0.5

CREATE_CONTAINER_NAME := angry_goodall

docker-create:
	#[[create container]]
	@docker create --name=${CREATE_CONTAINER_NAME} ${REGION}-docker.pkg.dev/${IMAGE_NAME}:${TAG_VERSION}

docker-tar:
	#[[export docker container to tar file]]
	@docker export ${CREATE_CONTAINER_NAME} > ${CREATE_CONTAINER_NAME}.tar

docker-build:
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
	@go test -short ./...

#Test with mock call to the app
test-mock:
	@go test ./internal/app/route

clean-test-cache:
	@go clean -testcache
