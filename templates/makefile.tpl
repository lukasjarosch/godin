GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GODIRS = $(shell go list -f '{{.Dir}}' ./...)
COMMIT=$(shell git rev-parse --short HEAD)
GOFILES=$(shell go list -f {{.Dir}} ./... | grep -v /vendor/)

DOCKER_IMAGE="<< .Docker.Registry >>/<< .Service.Namespace >>-<< .Service.Name >>"
DOCKER_TAG="v-${COMMIT}-dev"
PROTO_SRC="<< .Protobuf.Path >>"

.PHONY: build
build: check vendor
	@echo "--> building"
	@go build -o ${GOBIN}/<< .Service.Name >> ./cmd/<< .Service.Name >>/*.go

.PHONY: run
run:
	@echo "--> starting locally"
	@LOG_LEVEL=debug go run cmd/<< .Service.Name >>/*.go

.PHONY: vendor
vendor:
	@echo "--> vendoring dependencies"
	@go mod vendor

.PHONY: grpcui
grpcui:
	@grpcui -proto ${PROTO_SRC} -port 5000 -plaintext localhost:50051

.PHONY: clean
clean:
	@echo "--> cleaning up binaries"
	@go clean -tags netgo -i ./...
	@rm -rf $(GOBIN)/*

.PHONY: docker
docker: check vendor
	@echo "--> building docker image"
	docker build . -t ${DOCKER_IMAGE}:${DOCKER_TAG} -f ./Dockerfile

.PHONY: docker-run
docker-run:
	@echo "--> starting container ${DOCKER_IMAGE}:${DOCKER_TAG}"
	@docker run \
		--rm \
		--name << .Service.Name >> \
		--network host \
		${DOCKER_IMAGE}:${DOCKER_TAG}

.PHONY: check
check: format.check lint

.PHONY: format.check
format.check: tools.goimports

		@echo "--> checking code formatting with 'goimports'"
		@goimports -w ${GOFILES}


.PHONY: lint
lint: tools.golint
	@echo "--> checking code style with 'golint'"
	@echo $(GODIRS) | xargs -n 1 golint

#---------------
#-- tools
#---------------
.PHONY: tools tools.goimports tools.golint tools.govet

tools: tools.goimports tools.golint tools.govet

tools.goimports:
	@command -v goimports >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing goimports"; \
		go get golang.org/x/tools/cmd/goimports; \
	fi

tools.govet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		echo "--> installing govet"; \
		go get golang.org/x/tools/cmd/vet; \
	fi

tools.golint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get -u golang.org/x/lint/golint; \
	fi
