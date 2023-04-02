package main

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./configs/config_types.yaml ../../api/openapi/api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./configs/config_server.yaml ../../api/openapi/api.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=./configs/config_client.yaml ../../api/openapi/api.yaml

func main() {}
