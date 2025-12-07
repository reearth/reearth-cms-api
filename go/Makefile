update-schema:
	@echo "Updating schema..."
	@go run ./cmd/oapiupdater > ../schemas/integration.yml
	@echo "Done"

gen:
	@echo "Generating code..."
	@go tool oapi-codegen -config cmsoapi/oapi-codegen.yml -o cmsoapi/integration_gen.go https://raw.githubusercontent.com/reearth/reearth-cms/main/server/schemas/integration.yml
	@echo "Done"

.PHONY: update-schema gen
