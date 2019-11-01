.PHONY: deps lint test build run

SOURCES := $(shell find . -type d -name node_modules -prune -o -name '*.js' -print)

deps: node_modules/.installed

node_modules/.installed: package.json
	npm install
	touch node_modules/.installed

lint: node_modules/.installed $(SOURCES)
	npm run lint

test: deps lint
	npm test

clean:
	rm -rf node_modules

build: lint deps
	docker build --tag petewall/ota-service --file Dockerfile .

run:
	PORT=3000 index./.js