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

.PHONY: up
up:
	docker compose -f docker-compose.yml up -d
	docker compose ps

.PHONY: down
down:
	docker compose -f docker-compose.yml down

.PHONY: api-up
api-up:
	docker compose -f docker-compose.yml up -d api

.PHONY: api-down
api-down:
	docker compose -f docker-compose.yml down api

.PHONY: api-logs
api-logs:
	docker compose -f docker-compose.yml logs -f api

.PHONY: connect-db
connect-db:
	chmod +x ./scripts/connect-db.sh
	@./scripts/connect-db.sh $(ARGS)
