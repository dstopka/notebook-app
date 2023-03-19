#!/bin/bash

usage() {
    cat <<EOM
usage: "$0" SERVICE_NAME

    Generate go code for the protobuf file of the service given by name
EOM
}

validate_service() {
    services=(users notebooks)

    for name in "${services[@]}"; do 
        [[ "$1" == "$name" ]] && return 0
    done
    echo \'"$1"\' is not valid, must be one of: ["${services[*]}"] && usage && exit 1
}

[[ $# -ne 1 ]] && usage && exit 1

readonly service="$1"
validate_service $service

readonly out_dir_path="backend/common/genproto/${service}"
readonly protobuf_path="api/protobuf"

protoc --go_out=${out_dir_path} --go_opt=paths=source_relative \
       --go-grpc_out=${out_dir_path} --go-grpc_opt=paths=source_relative \
       --proto_path=${protobuf_path} ${protobuf_path}/${service}.proto
