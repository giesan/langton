# ------------------------------------------------------------- #
# HELPERS
# ------------------------------------------------------------- #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sort |  sed -e 's/^/  /'

# ------------------------------------------------------------- #
# TARGETS
# ------------------------------------------------------------- #

.DEFAULT_GOAL = qa

## analyze: analyze the code
analyze:
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

## build: build project binary
build: qa
	@go build -o ./build/langton .

## coverage: create coverage
coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/cover.out ./...
	@go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

## qa: run analysis and tests
qa: analyze test

## test: run all tests
test:
	@go test --cover ./...

.PHONY:
	analyze \
	build \
	coverage \
	help \
	qa \
	test \
