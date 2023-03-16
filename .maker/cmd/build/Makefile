include $(CURDIR)/../../lib/default.mk

###################
#╔═══════════════╗#
#║ BUILD ACTIONS ║#
#╚═══════════════╝#
###################

TARGET_OS = "linux" "darwin"
TARGET_ARCH = "amd64" "arm64"

current-platform: ## Build project
	rm -rf $(GWD)/.build/bin/*
	env CGO_ENABLED=0 go build -o $(GWD)/.build/bin/fizzbuzz $(GWD)/project/cmd/main.go

cross-platform: ## Build project for defined platform
	rm -rf $(GWD)/.build/bin/*
	for os in $(TARGET_OS); do                                                                                                        \
		for arch in $(TARGET_ARCH); do                                                                                                \
			echo $$os/$$arch;                                                                                                         \
			env CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -o $(GWD)/.build/bin/fizzbuzz-$$os-$$arch $(GWD)/project/cmd/main.go ; \
		done;                                                                                                                         \
	done;

docker: ## Build docker images
	docker build -t aubeio/fizzbuzz:latest -f $(GWD)/deploy/Dockerfile $(GWD)/.