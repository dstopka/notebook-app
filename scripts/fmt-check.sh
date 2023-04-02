#!/bin/bash

readonly check_result=$(gofmt -s -l .)
if [ ! -z "${check_result}" ]; then
    echo "The following files are not correctly formatted:"
    echo ${check_result} | tr "[:blank:]" "\n"
    exit 1
fi