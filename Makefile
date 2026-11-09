.PHONY: update-schema gen-go gen-ts gen

# REF: branch, tag or commit of reearth/reearth-cms to fetch (default: main)
#   e.g. make update-schema REF=0123abcd
REF ?= main

update-schema:
	@echo "Updating schemas/integration.yml from upstream (ref: $(REF))..."
	@cd tools/oapiupdater && REF=$(REF) go run . -o ../../schemas/integration.yml
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
