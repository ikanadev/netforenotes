//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./types.config.yaml ./openapi/openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./server.config.yaml ./openapi/openapi.yaml
package api
