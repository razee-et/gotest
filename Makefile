export GO111MODULE=auto

REPO_ROOT = $(shell git rev-parse --show-toplevel)
AMQ_CONF = "$(REPO_ROOT)/configs/test/amq"
AMQ_VER = "5.14.5"
AMQ_DOCKER = "message-broker-amq"

amq-start:
	docker build . -f activemq.Dockerfile -t $(AMQ_DOCKER)
	docker run -d --rm -p 61613:61613 -p 8161:8161 --name activemq $(AMQ_DOCKER)

amq-stop:
	docker stop activemq -t 0

amq-logs:
	docker logs -f -t activemq

docker-build:
	docker build . -f activemq.Dockerfile -t $(AMQ_DOCKER)

docker-run:
	docker run -d --rm -p 61613:61613 -p 8161:8161 --name activemq $(AMQ_DOCKER)

lint:
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)

lint-fix:
	gofmt -w .

unit-test:
	go test -v github.com/acquia/message-broker-sdk-go/pkg/client -tags=unit

integration-test:
	MESSAGE_BROKER_DEBUG=true go test -v github.com/acquia/message-broker-sdk-go/pkg/client -tags=integration

BUFFER := $(shell mktemp)

.PHONY: amq-start amq-stop amq-logs docker-build docker-run lint lint-fix unit-test integration-test
