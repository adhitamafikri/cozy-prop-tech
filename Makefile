default:
	@echo "Welcome Cozy Prop Tech"

.PHONY: run-web
run-web:
	cd ./frontend/web/
	bun -v
	bun run dev

.PHONY: run-listing-service
run-listing-service:
	docker compose -f infra/docker-compose.yml up listing-service

