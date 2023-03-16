include $(shell readlink -f $(CURDIR)/.maker/lib/default.mk)

doc: documentation ## Display all documentation actions

documentation:
	make -C $(CURDIR)/.maker/cmd/doc $(ARGS)

build: ## Display all builds actions
	make -C $(CURDIR)/.maker/cmd/build $(ARGS)

test: ## Display all tests actions
	make -C $(CURDIR)/.maker/cmd/test $(ARGS)

generate: ## Display all generate actions
	make -C $(CURDIR)/.maker/cmd/generate $(ARGS)

update: ## Install/Update vendor
	echo "Update all dependencies"
	go get -u $(GWD)/...
	go mod vendor

install: 
	go get -u $(ARGS)
	go mod vendor

run:
	make generate certificate > /dev/null
	make generate swagger > /dev/null
	go run project/cmd/main.go -t $(GWD)/.build/certs

run-from-docker:
	make generate swagger > /dev/null
	make build docker
	docker run -ti --rm --pid=host -p 80:80 -p 443:443 aubeio/fizzbuzz:latest

fix-perms:
	find $(GWD) -type d -print0 | xargs -0 chmod 0775
	find $(GWD) -type f -print0 | xargs -0 chmod 0664