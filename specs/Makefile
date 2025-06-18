OAS3_API_SPEC_FILEPREFIX=openapi-attio
OAS3_API_SPEC=$(OAS3_API_SPEC_FILEPREFIX).json
GEN_DIR=../
# This make file task produces the go code from the openapi specificaton code. 
# Tool reference: https://openapi-generator.tech/docs/configuration/
# Tool reference: https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go.md
codegen: 
	openapi-generator generate -i $(OAS3_API_SPEC) -g go -o $(GEN_DIR) -p packageName="github.com/fastah/libattio" -p useDefaultValuesForRequiredVars=true 
	sed -i '' 's|^module github.com/[^ ]*|module github.com/fastah/libattio|' $(GEN_DIR)/go.mod

deps: 
	brew install openapi-generator

deps-oas2:
	GOPRIVATE=github.com/fastah/*
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	npm i @apiture/openapi-down-convert

# Look inside the yaml specified in the --config flag to see where the code is generated.
codegen-oas2: 
	oapi-codegen --config=config-codegen.yaml $(OAS3_API_SPEC)

$(OAS3_API_SPEC_FILEPREFIX)-downgraded.json: openapi-attio.json
	openapi-down-convert --input $(OAS3_API_SPEC) --output $(OAS3_API_SPEC_FILEPREFIX)-downgraded.json