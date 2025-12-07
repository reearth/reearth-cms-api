update-schema:
	@echo "Updating schema..."
	@go run ./cmd/oapiupdater > ../schemas/integration.yml
	@echo "Done"

gen:
	@echo "Generating code..."
	@go generate ./...
	@echo "Done"

.PHONY: update-schema gen
