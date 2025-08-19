OAS3_API_SPEC_FILEPREFIX=openapi-attio
OAS3_API_SPEC=$(OAS3_API_SPEC_FILEPREFIX).json
GEN_DIR=../
# This make file task produces the go code from the openapi specificaton code. 
# Tool reference: https://openapi-generator.tech/docs/configuration/
# Tool reference (Go specific): https://openapi-generator.tech/docs/generators/go
codegen: 
	openapi-generator generate -i $(OAS3_API_SPEC) -g go -o $(GEN_DIR) -p packageName="libattio" -p useDefaultValuesForRequiredVars=true -p useOneOfDiscriminatorLookup=true
	find $(GEN_DIR) -type f \( -name '*.go' -o -name '*.mod'  -o -name '*.md'  -o -name 'go.*' \) -exec sed -i '' \
		-e 's|openapiclient|libattio|g' \
		-e 's|github\.com/GIT_USER_ID/GIT_REPO_ID|github.com/fastah/libattio|g' {} +
	cd $(GEN_DIR) && go mod tidy && go list -m -u all && go get -u -t ./... && go mod tidy 

deps: 
	brew install openapi-generator
