GO_FOLDERS=./cmd/server/... ./internal/... ./pkg/...
LOCAL_GO_FOLDERS=./src/github.com/zack-jack/pedal-tetris-api-v1/cmd/server/... ./src/github.com/zack-jack/pedal-tetris-api-v1/internal/... ./src/github.com/zack-jack/pedal-tetris-api-v1/pkg/...
ROOT_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOPATH=$(shell go env GOPATH)
BODYCLOSE=$(shell which bodyclose)
LINTER=golang.org/x/lint/golint
GOSEC=github.com/securego/gosec/v2/cmd/gosec

default: up

go-fmt:
	go fmt $(GO_FOLDERS)

go-vet:
	go vet $(GO_FOLDERS)

go-bodyclose:
	go vet -vettool=$(BODYCLOSE) $(GO_FOLDERS)

go-test:
	go test -mod=vendor -count=1 --short $(GO_FOLDERS)

go-lint:
	go get $(LINTER)
	go run $(LINTER) -set_exit_status $(GO_FOLDERS)

go-sec:
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.3.0
	./bin/gosec -quiet -exclude=G204,G201,G108,G601,G402 $(GO_FOLDERS)

go-fmt-check: go-fmt
	@if [ -z "`git status --porcelain`" ]; then exit 0; else echo "please run \`make go-fmt\` and commit the result" && exit 1; fi

vendor:
	go mod tidy && go mod vendor

up:
	docker-compose -f docker-compose.yml up --build -d

down:
	docker-compose -f docker-compose.yml down
