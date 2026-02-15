default:
	@echo "Welcome Cozy Prop Tech"

.PHONY: start-web
start-web:
	cd ./frontend/web/
	bun -v
	bun run dev

.PHONY: start-admin
start-admin:
	cd ./frontend/admin/
	bun -v
	bun run dev

.PHONY: start-api
start-api:
	docker compose -f docker-compose.yml up -d api

.PHONY: stop-api
stop-api:
	docker compose -f docker-compose.yml down api
