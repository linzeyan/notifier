.PHONY: stop
stop: ## stop docker-compose.dev.yml service 
	docker-compose -f docker-compose.dev.yml stop

.PHONY: clean
clean: stop ## clean docker-compose.dev.yml service
	docker-compose -f docker-compose.dev.yml rm

.PHONY: pull
pull: ## pull docker-compose.dev.yml service
	docker-compose -f docker-compose.dev.yml pull

.PHONY: start-dev
start-dev: 
	docker-compose -f docker-compose.dev.yml up dev

.PHONY: start-prod
start-prod: 
	docker-compose -f docker-compose.dev.yml up prod

.PHONY: dev
dev: stop ## start dev docker-compose.dev.yml service (no daemon)
	@-$(MAKE) start-dev

.PHONY: prod
prod: stop pull ## start prod docker-compose.dev.yml service (no daemon)
	@-$(MAKE) start-prod
	
.PHONY: test
test: ## test docker-compose.dev.yml service
	@curl -X "POST" "http://localhost:20001/send" -H 'Content-Type: application/json; charset=utf-8' -d '{"message": "🚖 *bold*\n*message*: `test` "}'
	@curl -X "POST" "http://localhost:20002/send" -H 'Content-Type: application/json; charset=utf-8' -d '{"message": "🚖 *bold*\n*message*: `test` "}'

.PHONY: help
help: ## list command:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)



