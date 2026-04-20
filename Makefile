.PHONY: update-schema gen-go gen-ts gen

update-schema:
	@echo "Updating schemas/integration.yml from upstream..."
	@cd tools/oapiupdater && go run . > ../../schemas/integration.yml
	@echo "Done"

gen-go:
	@echo "Generating Go code..."
	@cd go && go generate ./...
	@echo "Done"

gen-ts:
	@echo "Generating TypeScript types..."
	@cd ts && npm run generate
	@echo "Done"

gen: update-schema gen-go gen-ts
