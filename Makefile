all: deps fmt vet cmd aliyun_fc

ragel-go: FORCE
	@echo "ragel generate go code"
	# @ragel -Z -G2 -e -o ragel/parse_datetime.go ragel/parse_datetime.go.rl
	@ragel -Z -G2 -e -o ragel/parse_datetime.go ragel/parse_datetime.go.rl

deps: FORCE
	@echo "update dependency"
	@go mod tidy

fmt: FORCE
	@echo "Formatting code"
	@go list -f '{{.Dir}}' ./  | xargs -L1 gofmt -l

vet:
	@echo "Vetting code"
	@go vet ./.

cli: fmt
	cd cli && go build -o ../bin/anytime

clean: FORCE
	@echo "Cleaning"
	rm -rf ./bin

FORCE: