install-go:
	sh script/install-go.sh 1.6.2

install-glide:
	sh script/install-glide.sh 0.8.3

deps: install-glide
	go get github.com/golang/lint/golint
	glide install

deps-update: install-glide
	rm -rf ./vendor
	glide update

test: go-test go-vet

go-test:
	go test $(shell go list ./... | grep -v vendor)

go-vet:
	go vet $(shell go list  ./... | grep -v vendor)

build:
	GOOS=$(ARCH) GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o bin/cloudinit-helper main.go

mock:
	mockgen -source client/ec2/client.go -package ec2 -destination client/ec2/client_mock.go
	mockgen -source client/ec2meta/client.go -package ec2meta -destination client/ec2meta/client_mock.go
