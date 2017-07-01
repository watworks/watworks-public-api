dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

dev-rebuild: dev-down
	docker-compose up --build -d

dev-init:
	@if [ ! "$(GOPATH)" ]; then \
        echo "GOPATH needs to be specified for code completion to work in your IDE"; \
        return 1; \
    fi
	echo $(GOPATH)
	mkdir -p $(GOPATH)/src/github.com/watworks
	ln -s $(shell pwd) $(GOPATH)/src/github.com/watworks/watworks-public-api

deps-update:
	docker-compose run --rm app dep ensure -update

build:
	docker-compose run --rm -e CGO_ENABLED=0 -e GOOS=linux app go build -o .bin/watworks-public-api .

run: build
	docker-compose run --rm app .bin/watworks-public-api

test:
	docker-compose run --rm app sh -c 'go list ./...| grep -v /vendor/ | xargs go test'

test-package:
	docker-compose run --rm app go test github.com/watworks/watworks-public-api/$(PKG)

fmt:
	docker-compose run --rm app go fmt

build-image: build
	docker build -t watworks/watworks-public-api .

publish-tag:
	echo TODO

.PHONY: dev-up dev-down dev-rebuild deps-update test test-package fmt build publish
