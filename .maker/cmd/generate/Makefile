include $(shell readlink -f $(CURDIR)/../../lib/default.mk)

###################
#╔═══════════════╗#
#║ CERTS ACTIONS ║#
#╚═══════════════╝#
###################

certificate: ## Generate ssl certificate
	if [ ! -f $(GWD)/.build/certs/server.key ] || [ ! -f $(GWD)/.build/certs/server.crt ]; then \
		rm -rf $(GWD)/.build/certs; \
		mkdir -p $(GWD)/.build/certs; \
		openssl genrsa -out $(GWD)/.build/certs/server.key 4096; \
		openssl req -new -key $(GWD)/.build/certs/server.key -out $(GWD)/.build/certs/server.csr -sha512 -subj "/CN=localhost"; \
		openssl x509 -req -days 3650 -in $(GWD)/.build/certs/server.csr -signkey $(GWD)/.build/certs/server.key -out $(GWD)/.build/certs/server.crt -sha512; \
	fi

swagger: ## Generate swagger documentation
	rm -rf $(GWD)/project/internal/docs
	swag init -d $(GWD)/project -g cmd/main.go -o $(GWD)/project/internal/docs
	chmod -R 644 $(GWD)/project/internal/docs/*