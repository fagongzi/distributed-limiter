#!/bin/bash
#
# Generate all elasticell protobuf bindings.
# Run from repository root.
#
set -e

# directories containing protos to be built
DIRS="./pb"

GOGOPROTO_ROOT="${GOPATH}/src/github.com/gogo/protobuf"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"
PB_PATH="${GOPATH}/src/github.com/fagongzi/distributed-limiter/pkg"

for dir in ${DIRS}; do
	pushd ${dir}
		protoc --gofast_out=plugins=grpc,import_prefix=github.com/fagongzi/distributed-limiter/pkg/:. -I=.:"${GOGOPROTO_PATH}":"${PB_PATH}":"${GOPATH}/src" *.proto
		sed -i.bak -E 's/github\.com\/fagongzi\/'"distributed-limiter"'\/pkg\/(gogoproto|github\.com|golang\.org|google\.golang\.org)/\1/g' *.pb.go
		sed -i.bak -E 's/github\.com\/fagongzi\/'"distributed-limiter"'\/pkg\/(errors|fmt|io)/\1/g' *.pb.go
		sed -i.bak -E 's/import _ \"gogoproto\"//g' *.pb.go
		sed -i.bak -E 's/import fmt \"fmt\"//g' *.pb.go
		sed -i.bak -E 's/import math \"github.com\/fagongzi\/'"distributed-limiter"'\/pkg\/math\"//g' *.pb.go
		rm -f *.bak
		goimports -w *.pb.go
	popd
done
