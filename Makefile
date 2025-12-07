generate-ent:
	go generate ./ent

generate-oapi:
	go generate ./api

tidy:
	go mod tidy

entviz:
	go tool ariga.io/entviz ./ent/schema