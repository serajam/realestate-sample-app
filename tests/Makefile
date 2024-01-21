export GOPROXY=http://proxy.golang.org,direct
export GOSUMDB=off
export GO111MODULE=on
export GOPRIVATE=bitbucket.org/iskatel-ua

.PHONY: test.e2e
test.e2e:
	if go test ./e2e -v -count=1; then \
		echo "Testing succeeded"; \
	else \
		echo "Testing failed"; \
		exit 1; \
	fi

t:
	go test -json -v ./... 2>&1 | tee /tmp/gotest.log | docker run -i ghcr.io/gotesttools/gotestfmt:latest

.PHONY: fmt
fmt:
	gofmt -w -s cmd
	gofmt -w -s pkg

.PHONY: help
help:
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "help", "print this help"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "test.e2e", "run go e2e tests"}'
