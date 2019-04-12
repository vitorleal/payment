.PHONY: all install test deps

export GOPATH="${HOME}/Code/go"

all: build

install: deps
	govendor sync

test:
	go test -v -covermode=count -coverprofile=coverage.out

deps:
	@hash govendor > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/kardianos/govendor; \
	fi

