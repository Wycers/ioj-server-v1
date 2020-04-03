


apps = 'problems' 'users' 'server' 'files'

.PHONY: run
run: proto wire
	for app in $(apps) ;\
	do \
		go build -o build/$$app ./cmd/$$app/; \
	done
	for app in $(apps) ;\
	do \
		./build/$$app -f configs/$$app.yml  & \
	done

.PHONY: run-cli
run-cli: proto wire
	go build ./cmd/cli

.PHONY: wire
wire:
	wire ./...



.PHONY: proto
proto:
	protoc -I api/protobuf-spec ./api/protobuf-spec/*.proto --go_out=plugins=grpc:api/protobuf-spec