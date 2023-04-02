#!/bin/bash

usage() {
    cat <<EOM
usage: "$0" SERVICE_NAME

    Generate go code for the protobuf file of the service given by name
EOM
}

[[ $# -ne 1 ]] && usage && exit 1
readonly service="$1"

readonly base_dir=$(git rev-parse --show-toplevel)
readonly out_dir_path="${base_dir}/backend/${service}/pkg/${service}"
readonly proto_path="${base_dir}/api/protobuf"
readonly proto_file="${proto_path}/${service}.proto"

echo "Generating go code for ${proto_file}"

protoc --go_out=${out_dir_path} --go_opt=paths=source_relative \
       --go-grpc_out=${out_dir_path} --go-grpc_opt=paths=source_relative \
       --proto_path=${proto_path} ${proto_file}
