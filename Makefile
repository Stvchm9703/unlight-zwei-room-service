
	# Generate gogo, gRPC-Gateway, swagger, go-validators output.
	#
	# -I declares import folders, in order of importance
	# This is how proto resolves the protofile imports.
	# It will check for the protofile relative to each of these
	# folders and use the first one it finds.
	#
	# --gogo_out generates GoGo Protobuf output with gRPC plugin enabled.
	# --grpc-gateway_out generates gRPC-Gateway output.
	# --swagger_out generates an OpenAPI 2.0 specification for our gRPC-Gateway endpoints.
	# --govalidators_out generates Go validation files for our messages types, if specified.
	#
	# The lines starting with Mgoogle/... are proto import replacements,
	# which cause the generated file to import the specified packages
	# instead of the go_package's declared by the imported protof files.
	#
	# $$GOPATH/src is the output directory. It is relative to the GOPATH/src directory
	# since we've specified a go_package option relative to that directory.
	#
	# proto/example.proto is the location of the protofile we use.

LOGPATH = ./.log
LOGFILE = $(LOGPATH)/$(shell date --iso=seconds).log
test_run:
	go run ./build_cli/room_status.go start -c=config.test_server.yaml -m=test > $(LOGFILE) 2>&1 &

generate:
	protoc \
		-I proto \
		-I vendor/ \
		--go_out=plugins=grpc:proto \
		proto/service.proto

install:
	go get \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		github.com/rakyll/statik

build: 
	go build -o bin/roomstatus_server.exe main.try.go



generate_vcred:
	protoc \
		-I proto/ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--go_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--grpc-gateway_out=allow_patch_feature=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		--govalidators_out=gogoimport=true,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
$(CURDIR)/vendor/ \
		proto/*.proto
	# gvm issue :  move the genrated file to current directory
	mv $(CURDIR)/vendor/*.go $(CURDIR)/proto/
	

build_docker:
	go build -o room_server build_cli/room_status.go
	# docker build -t ulz-room-service .
	docker-compose build 

build_doc_message :
	protoc --doc_out=./doc --doc_opt=markdown,message.md proto/message.proto

build_doc_service :
	protoc --doc_out=./doc --doc_opt=markdown,service.md proto/service.proto
	
