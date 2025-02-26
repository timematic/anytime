all: deps fmt vet cmd aliyun_fc

ragel-go: FORCE
	@echo "ragel generate go code"
	# @ragel -Z -G2 -e -o ragel/parse_datetime.go ragel/parse_datetime.go.rl
	@ragel -Z -e -o ragel/parse_datetime.go ragel/parse_datetime.go.rl

deps: FORCE
	@echo "update dependency"
	@go mod tidy

fmt: FORCE
	@echo "Formatting code"
	@go list -f '{{.Dir}}' ./  | xargs -L1 gofmt -l

vet:
	@echo "Vetting code"
	@go vet ./.

aliyun_fc: FORCE
	@echo "Building aliyun_fc"
	cd aliyun_fc && GOOS=linux GOARCH=amd64 go build -o ../bin/aliyun_fc
	@GIT_COMMIT=$$(git log -1 --pretty=format:"%h"); \
	zip -r "aliyun_fc.$$GIT_COMMIT.zip" bin/aliyun_fc assets/*.json && echo "package success: aliyun_fc.$$GIT_COMMIT.zip"

cmd: fmt
	cd cmd && go build -o ../bin/make_cedict_json

clean: FORCE
	@echo "Cleaning"
	rm -rf ./bin

FORCE: