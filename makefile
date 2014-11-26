.PHONY: all fmt test testv tags doc vet lc

all: build

build:
	@ GOPATH=`pwd` go install -v ./src/...

fmt: 
	@ GOPATH=`pwd` gofmt -s -l -w src

vet: 
	@ GOPATH=`pwd` go vet ./src/...

testv:
	@ GOPATH=`pwd` go test -v ./src/...

testc:
	@ GOPATH=`pwd` go test -cover -coverprofile=cover.out ./src/...

test:
	@ GOPATH=`pwd` go test ./src/...

clean:
	@ rm -rf pkg bin

fix:
	@ GOPATH=`pwd` go fix ./src/...

tags:
	@ gotags -R src > tags

doc:
	@ GOPATH=`pwd` godoc -http=:8000

lc:
	@ wc -l `find src -name "*.go"`

lint:
	@ GOPATH=`pwd` golint ./src/...
