.PHONY: build clean lint

build:
	go build -o testapp

clean:
	rm testapp

lint:
	gometalinter --enable=misspell --enable=unparam --enable=dupl --enable=gofmt --enable=goimports --disable=gotype --disable=gas --deadline=3m ./...

