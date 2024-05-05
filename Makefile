ARTIFACT_NAME := nhlscores

build:
	@go build -o bin/${ARTIFACT_NAME}/${ARTIFACT_NAME} cmd/${ARTIFACT_NAME}/main.go 

clean:
	rm -rf bin/

run:
	@go run cmd/${ARTIFACT_NAME}/main.go

quay:
	docker buildx create --use --platform linux/amd64,linux/arm64 --name quaybuilder
	docker buildx ls
	docker buildx build --platform linux/amd64,linux/arm64 --push -t quay.io/omaciel/nhlscores:latest .
	docker buildx stop quaybuilder
	docker buildx rm quaybuilder
