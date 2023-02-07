
generate:
	go generate ./...

test: test.unit test.integ

test.integ:
	ginkgo -v --fail-fast --tags=integ

test.unit:
	ginkgo -v --fail-fast
