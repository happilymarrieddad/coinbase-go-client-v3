install:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

generate:
	go generate ./...

test: test.unit test.integ

test.integ:
	ginkgo -v --fail-fast --tags=integ

test.unit:
	ginkgo -v --fail-fast
