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
	make generate swagger > /dev/null
	go run project/cmd/main.go 