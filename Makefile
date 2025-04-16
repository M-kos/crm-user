generate-openapi:
	openapi-generator generate -g go-server -i ./openapi/openapi.yaml -o src/generated/ -c ./openapi/config.yaml

.PHONY: generate-openapi