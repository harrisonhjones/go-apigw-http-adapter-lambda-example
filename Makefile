# default make target.
.PHONY: release
release: clean fmt vet test build zip

.PHONY: clean
clean:
	rm -f go-apigw-http-adapter-lambda-example
	rm -f go-apigw-http-adapter-lambda-example.zip

.PHONY: build
build:
	go build -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: vet
vet:
	go vet -v ./...

.PHONY: zip
zip:
	zip go-apigw-http-adapter-lambda-example.zip go-apigw-http-adapter-lambda-example