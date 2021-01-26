export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

fmt:
	go fmt ./...

build: 
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o ./bin/autorun .

test: 
	go test .
	
clean:
	rm -f ./bin/autorun