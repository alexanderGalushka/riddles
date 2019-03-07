lint:
	@echo "Linting source with lint and vet"
	go vet ./api/... && golint -set_exit_status=true ./api/...

unit:
	go test -v -tags=unit ./api/...

goapi:
	go run api/main.go

jsclient:
	node client/server.js

run: golang client