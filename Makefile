generate-ent:
	go generate ./pkg/ent

generate-oapi:
	go generate ./internal/controller/rest/v1

tidy:
	go mod tidy

entviz:
	go tool ariga.io/entviz ./ent/schema