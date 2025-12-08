tidy:
	go mod tidy

golangci-lint-run:
	golangci-lint run

golangci-lint-fix:
	golangci-lint run --fix

golangci-lint-fmt:
	golangci-lint fmt

generate-oapi:
	go generate ./internal/controller/rest/v1

generate-ent:
	go generate ./pkg/ent

entviz:
	go tool ariga.io/entviz ./ent/schema
