SERVICE_NAME=ticket

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GODIRS = $(shell go list -f '{{.Dir}}' ./...)
KUBE_CONTEXT=$(cat ~/.kube/config | grep "current-context:" | sed "s/current-context: //")

COMMIT=$(shell git rev-parse --short HEAD)

DOCKER_IMAGE="docker.coozzy.ch/ticket-ticket"
DOCKER_TAG="v-${COMMIT}-dev"

MODULE="bitbucket.org/jdbergmann/ticket/ticket"
PROTO_SRC="/home/lukas/devel/work/protobuf/ticket/ticket"
PROTO_DST="proto"

.PHONY: build
build:
	@echo "--> building"
	@go build -o ${GOBIN}/ticket ./cmd/ticket/main.go

.PHONY: run
run:
	echo "--> starting locally"
	@LOG_LEVEL=${LOG_LEVEL} AMQP_CONNECTION=${AMQP_CONNECTION} AMQP_EXCHANGE=${AMQP_EXCHANGE} go run cmd/ticket/main.go

.PHONY: grpcui
	grpcui:
	grpcui -proto ${PROTO_SRC}/ticket.proto -port 5000 -plaintext localhost:50051

.PHONY: clean
clean:
	echo "--> cleaning up binaries"
	go clean -tags netgo -i ./...
	rm -rf $(GOBIN)/*

.PHONY: clean
skaffold:
	@echo "--> starting skaffold, make sure minikube is running"
	@skaffold dev

.PHONY: docker
docker:
	@echo "--> building docker image"
	docker build . -t ${DOCKER_IMAGE}:${DOCKER_TAG} -f ./Dockerfile

.PHONY: docker-run
docker-run:
	@echo "--> starting container ${DOCKER_IMAGE}:${DOCKER_TAG}"
	@docker run \
		--rm \
		--name ${ticket} \
		--network host \
		${DOCKER_IMAGE}:${DOCKER_TAG}

.PHONY: proto
proto:
	@rm -rf ${PROTO_DST}
	@echo "--> fetching proto from ${PROTO_SRC}"
	@make -C ${PROTO_SRC} go
	@mkdir ${PROTO_DST}
	@echo "--< moving stubs into ${PROTO_DST}"
	@cp ${PROTO_SRC}/go/* ${PROTO_DST}

.PHONY: check
check: format.check lint

.PHONY: format.check
format.check: tools.goimports
	@echo "--> checking code formatting with 'goimports'"
	@goimports -w -l .

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
