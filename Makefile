


apps = 'problems' 'users' 'server' 'files' 'submissions' 'judgements'

.PHONY: build
build: wire
	for app in $(apps) ;\
	do \
		go build -o dist/$$app ./cmd/$$app/; \
		GOOS=linux GOARCH="amd64" go build -o dist/$$app-linux-amd64 ./cmd/$$app/; \
	done

.PHONY: dev
dev:
	for app in $(apps) ;\
	do \
		CompileDaemon -build="go build -o ./dist/$$app.exe ./cmd/$$app/" -command="./dist/$$app.exe -f configs/$$app.yml" & \
	done

.PHONY: run
run:
	for app in $(apps) ;\
	do \
		./dist/$$app -f configs/$$app.yml  & \
	done
.PHONY: cli
cli: wire
	go build ./cmd/cli

.PHONY: wire
wire:
	wire ./...

.PHONY: proto
proto:
	protoc -I api/protobuf-spec ./api/protobuf-spec/*.proto --go_out=plugins=grpc:api/protobuf-spec


.PHONY: deploy
deploy:
	docker-compose -f deployments/docker-compose.yml up --build


.PHONY: mock
mock:
	mockery --all