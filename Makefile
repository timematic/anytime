OS_NAME := $(shell uname -s)

all: deps fmt vet cli test

deps: FORCE
# install ragel if not exist
ifeq (, $(shell which ragel))
ifeq ($(OS_NAME),Darwin)
		@echo "brew install ragel"
		@brew install ragel
else
	@echo "dnf install -y ragel"
	@dnf install -y ragel
endif
endif
	@echo "go mod tidy"
	@go mod tidy

ragel-go: FORCE
	@echo "ragel -Z -G2 -e -o ragel_parse_datetime.go ragel/parse_datetime.go.rl"
	@ragel -Z -G2 -e -o ragel_parse_datetime.go ragel/parse_datetime.go.rl

fmt: FORCE
	@echo "Formatting code"
	@go list -f '{{.Dir}}' ./  | xargs -L1 gofmt -l

vet:
	@echo "Vetting code"
	@go vet ./.

test:
	@go test -v ./tests -rapid.checks=10000

bench:
	cd tests && go test -bench .

cli: fmt
	cd cli && go build -o ../bin/anytime

clean: FORCE
	@echo "Cleaning"
	rm -rf ./bin

FORCE: