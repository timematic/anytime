all: deps fmt vet cmd aliyun_fc

ragel-go: FORCE
	@echo "ragel generate go code"
	@ragel -Z -G2 -e -o ragel_parse_datetime.go ragel_parse_datetime.go.rl
	@ragel -Z -G2 -e -o atoi.go atoi.go.rl

fmt: FORCE
	@echo "Formatting code"
	@go list -f '{{.Dir}}' ./  | xargs -L1 gofmt -l

FORCE: