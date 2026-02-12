default:
	@echo "Welcome Cozy Prop Tech"

.PHONY: run-web
run-web:
	cd ./frontend/web/
	bun -v
	bun run dev

