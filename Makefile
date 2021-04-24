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

test-firmware:
	mkdir -p data/firmware/test/1.2.3/
	touch data/firmware/test/1.2.3/test-1.2.3.bin

clean:
	rm -rf node_modules

build: lint deps
	docker build --tag petewall/ota-service --file Dockerfile .

run: deps
	PORT=8266 DATA_DIR=$(shell pwd)/data ./index.js

# CI targets
set-pipeline:
	fly -t wallhouse set-pipeline \
		--load-vars-from ../secrets/pipeline-creds.json \
		--pipeline ota-service \
		--config ci/pipeline.yaml
