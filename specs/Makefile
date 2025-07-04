OAS3_API_SPEC_FILEPREFIX=openapi-attio
OAS3_API_SPEC=$(OAS3_API_SPEC_FILEPREFIX).json
GEN_DIR=../
# This make file task produces the go code from the openapi specificaton code. 
# Tool reference: https://openapi-generator.tech/docs/configuration/
# Tool reference (Go specific): https://openapi-generator.tech/docs/generators/go
codegen: 
	openapi-generator generate -i $(OAS3_API_SPEC) -g go -o $(GEN_DIR) -p packageName="libattio" -p useDefaultValuesForRequiredVars=true 
	find $(GEN_DIR) -type f \( -name '*.go' -o -name '*.mod'  -o -name '*.md'  -o -name 'go.*' \) -exec sed -i '' \
		-e 's|openapiclient|libattio|g' \
		-e 's|github\.com/GIT_USER_ID/GIT_REPO_ID|github.com/fastah/libattio|g' {} +
	cd $(GEN_DIR) && go mod tidy && go list -m -u all && go get -u -t ./... && go mod tidy 

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