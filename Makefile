.PHONY: help build clean test lint lint-misspell lint-golint docs docs-tidy docs-test

export GO111MODULE=on

NOW=$(shell TZ=Asia/Shanghai date '+%Y-%m-%d_%H:%M:%S')
REV=$(shell git rev-parse HEAD)
LD_FLAGS="-X main.Build=${NOW}@${REV}"

help:
	@printf "Commands:\n"
	@printf "  build\t\tCompiles source code into binaries\n"
	@printf "  clean\t\tDeletes compiled binaries\n"
	@printf "  test\t\tTest all packages\n"

build:
	@mkdir -p bin && cd bin && go build -ldflags ${LD_FLAGS} ../cmd/...

clean:
	@rm -rf "$(PWD)/bin" "$(PWD)/cover"
	@find "$(PWD)" -type f -name '*.out' -delete
	@find "$(PWD)" -type f -name '*.xml' -delete

test:
	sh gofmtcheck.sh
	@go test -covermode=count -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

lint: lint-misspell lint-golint

lint-misspell:
	@[ -z "$(shell { command -v misspell; } 2>/dev/null)" ] && \
		go get -u github.com/client9/misspell/cmd/misspell || true
	@find . -not -path '*/vendor/*' -type f \( -name '*.adoc' -o -name '*.go' \) -exec misspell -error {} +

lint-golint:
	@[ -z "$(shell { command -v golint; } 2>/dev/null)" ] && \
		go get -u golang.org/x/lint/golint || true
	@go list ./... | xargs -L1 golint
