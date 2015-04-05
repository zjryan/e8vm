.PHONY: all fmt tags doc

all:
	go install -v ./...
	gofmt -s -w -l .
	go install -v ./...
	e8chk -path="lonnie.io/e8vm"
	golint ./...
	gotags -R . > tags

rall:
	touch `find . -name "*.go"`
	go install -v ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags -R . > tags

test:
	go test ./...

testv:
	go test -v ./...

lc:
	wc -l `find . -name "*.go"`

doc:
	godoc -http=:8000

lint:
	golint ./...
