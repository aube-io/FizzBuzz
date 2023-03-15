include $(CURDIR)/../../lib/default.mk

###################
#╔═══════════════╗#
#║ BUILD ACTIONS ║#
#╚═══════════════╝#
###################

TARGET_OS = "linux" "darwin"
TARGET_ARCH = "amd64" "arm64" 

current-platform: ## Build project
	rm -rf $(CURDIR)/.build/*
	go build -o $(CURDIR)/.build/fizzbuzz ./app/main.go

cross-platform: ## Build project for defined platform
	for os in $(TARGET_OS); do                                                                            \
		for arch in $(TARGET_ARCH); do                                                                    \
			echo $$os/$$arch;                                                                             \
			env GOOS=$$os GOARCH=$$arch go build -o $(CURDIR)/.build/fizzbuzz-$$os-$$arch ./app/main.go ; \
		done;                                                                                             \
	done;

docker: project ## Build docker images
