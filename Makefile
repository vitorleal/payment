all: build

install: deps
	govendor sync

.PHONY: test
test:
	go test -v -covermode=count -coverprofile=coverage.out

deps:
	@hash govendor > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/kardianos/govendor; \
	fi

