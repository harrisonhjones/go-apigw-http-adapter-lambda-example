# default make target.
.PHONY: release
release: clean fmt vet test build zip

.PHONY: clean
clean:
	rm -fr bin

.PHONY: build
build:
	mkdir -p bin
	go build -v -o bin ./...

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
zip: build
	cd bin && zip httpapi.zip httpapi
	cd bin && zip restapi.zip restapi