BIN_DIR = bin
PROTO_DIR = proto
API_NAME = shelter

migration-up:
	@echo "migration-up";
migration-down:
	@echo "migration-down";
build-app:
	@go build -o ${BIN_DIR}/app ./
build-api:
	@protoc -I${PROTO_DIR} --go_out=. --go-grpc_out=. ${PROTO_DIR}/${API_NAME}.proto
	@echo "build api: DONE"
build-all:
	@echo "build-all";
