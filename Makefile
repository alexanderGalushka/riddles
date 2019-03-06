lint:
	@echo "Linting source with lint and vet"
	go vet ./api/... && golint -set_exit_status=true ./api/...

docker-push:
	docker push $(repo)/$(app_name):$(version)

unit:
	go test -v -tags=unit ./api/...

integration-ci:
	./api/scripts/test/integration-test.sh

golang:
	go run api/main.go

nodejs:
	node ui/server.js

run: golang nodejs