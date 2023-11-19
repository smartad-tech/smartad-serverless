.PHONY: build
build:
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/ ./cmd/...
	
.PHONY: deploy	
deploy:
	cd cdk/ && cdk deploy --profile private