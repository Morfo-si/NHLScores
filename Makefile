ARTIFACT_NAME := nhlscores

build:
	@go build -o bin/${ARTIFACT_NAME}/${ARTIFACT_NAME} cmd/${ARTIFACT_NAME}/main.go 

clean:
	rm -rf bin/

run:
	@go run cmd/${ARTIFACT_NAME}/main.go