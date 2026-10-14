BINARY ?= ticket

.PHONY: clean

$(BINARY): $(shell find . -type f -name '*.go')
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)
