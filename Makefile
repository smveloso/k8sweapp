.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o k8swebapp ./cmd/main.go
.PHONY:build

docker-build: build
	docker build -t smveloso/k8swebapp:latest  -f build/package/Dockerfile .

docker-push: docker-build
	docker push smveloso/k8swebapp:latest
 
