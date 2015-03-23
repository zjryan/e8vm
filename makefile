.PHONY: all fmt tags doc

all:
	go install -v ./...

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

asmt:
	make -C asm/tests --no-print-directory

stayall:
	STAYPATH=`pwd`/stay-tests stayall

lint:
	golint ./...

symdep:
	symdep lonnie.io/e8vm/arch8
	symdep lonnie.io/e8vm/dasm8
	symdep lonnie.io/e8vm/link8
	symdep lonnie.io/e8vm/lex8
	symdep lonnie.io/e8vm/asm8
	symdep lonnie.io/e8vm/build8

check: fmt all lint symdep
