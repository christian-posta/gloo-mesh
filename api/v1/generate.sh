#!/usr/bin/env bash

set -ex

# for symlink compatibility
# https://stackoverflow.com/questions/59895/getting-the-source-directory-of-a-bash-script-from-within
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done

IN="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
OUT=${IN}/../../pkg/api/v1/

GOGO_OUT_FLAG="--gogo_out=Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:${GOPATH}/src/"
PROTOC_FLAGS="-I=${GOPATH}/src/ \
    -I=${GOPATH}/src/github.com/solo-io/solo-kit/api/external/proto \
    -I=${GOPATH}/src/github.com/gogo/protobuf/ \
    -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
    ${GOGO_OUT_FLAG}"

mkdir -p ${OUT}
protoc -I=${IN} \
    ${PROTOC_FLAGS} \
    ${IN}/*.proto
