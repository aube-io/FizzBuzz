include $(CURDIR)/../../lib/default.mk

###################
#╔═══════════════╗#
#║ BUILD ACTIONS ║#
#╚═══════════════╝#
###################

TARGET_OS = "linux" "darwin"
TARGET_ARCH = "amd64" "arm64" 

current-platform: ## Build project
	rm -rf $(GWD)/.build/*
	go build -o $(GWD)/.build/bin/fizzbuzz $(GWD)/project/cmd/main.go

cross-platform: ## Build project for defined platform
	rm -rf $(GWD)/.build/*
	for os in $(TARGET_OS); do                                                                                          \
		for arch in $(TARGET_ARCH); do                                                                                  \
			echo $$os/$$arch;                                                                                           \
			env GOOS=$$os GOARCH=$$arch go build -o $(GWD)/.build/bin/fizzbuzz-$$os-$$arch $(GWD)/project/cmd/main.go ; \
		done;                                                                                                           \
	done;

docker: project ## Build docker images
