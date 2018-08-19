

.PHONY: start
start: ## start docker-compose.dev.yml service
	@-docker-compose -f docker-compose.dev.yml up

.PHONY: stop
stop: ## stop docker-compose.yml service 
	docker-compose -f docker-compose.dev.yml stop

.PHONY: dev
dev: ## start docker-compose.dev.yml service
	docker-compose -f docker-compose.dev.yml up -d

.PHONY: restart
restart: stop dev ## restart docker-compose.dev.yml service

.PHONY: test
test: stop ## start docker-compose.dev.yml service (no daemon)
	@-$(MAKE) start

.PHONY: help
help: ## list command:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)



