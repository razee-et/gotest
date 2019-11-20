export GO111MODULE=auto

lint:
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)

lint-fix:
	gofmt -w .

BUFFER := $(shell mktemp)

.PHONY: lint lint-fix
