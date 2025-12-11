tidy:
	go mod tidy
.PHONY: tidy

golangci-lint-run:
	golangci-lint run
.PHONY: golangci-lint-run

golangci-lint-fix:
	golangci-lint run --fix
.PHONY: golangci-lint-fix

golangci-lint-fmt:
	golangci-lint fmt
.PHONY: golangci-lint-fmt

generate-oapi:
	go generate ./internal/controller/rest/v1
.PHONY: generate-oapi

generate-mock:
	go generate ./internal/usecase
	go generate ./internal/repository
.PHONY: mock

generate-ent:
	go generate ./pkg/ent
.PHONY: generate-ent

entviz:
	go tool ariga.io/entviz ./ent/schema
.PHONY: entviz
