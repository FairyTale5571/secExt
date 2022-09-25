NAME=secExtNew

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

#.SILENT:

define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## ------------------------------------------------- Common commands: --------------------------------------------------
## Formats the code.
format:
	${call colored,formatting is running...}
	go vet ./...
	go fmt ./...

## Fix-imports order.
imports:
	${call colored,fixing imports...}
	./scripts/fix-imports-order.sh

## Fix fields.
fields:
	${call colored,fixing fields...}
	./scripts/fieldalignment.sh

## lint project
lint:
	${call colored,lint is running...}
	./scripts/linters.sh
.PHONY: lint

fumpt:
	${call colored,fumpt is running...}
	gofumpt -l -w .

build:
	${call colored,building x64...}
	GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o release/${NAME}_x64.dll cmd/main.go

	${call colored,building x86...}
	GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o release/${NAME}.dll cmd/main.go

test64:
	cd release && ./callExtension_x64.exe "secExt" callExtension ["Version",[]]

