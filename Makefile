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
