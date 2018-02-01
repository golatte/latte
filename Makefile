.PHONY: bin test all fmt deploy docs cli latte

all: fmt bin

bin: cli

cli: latte

latte:
	(cd ./latte; go build latte.go)

docs:
	./makedocs.sh

fmt:
	-go fmt ./...
