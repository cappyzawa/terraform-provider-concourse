TEST?=$$(go list ./... | grep -v 'vendor')
BINARY=terraform-provider-concourse
VERSION=0.0.1

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_linux_amd64
	chmod +x ./bin/*

install: build
	mv ${BINARY} ~/.terraform.d/plugins

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
