#!/bin/sh

# This script compiles the protobuf files in the current directory

if ! [ -x "$(command -v protoc)" ]; then
  echo 'Error: protoc is not installed.' >&2
  exit 1
fi

# Compile the protobuf files in the current directory

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"


mkdir -p "${SCRIPT_DIR}"/../../agenda_server
mkdir -p "${SCRIPT_DIR}"/../../agenda_server/v1

cd "${SCRIPT_DIR}" || exit

echo "Compiling agenda.proto"
protoc --go_out="${SCRIPT_DIR}"/../../agenda_server/v1 --go-grpc_out="${SCRIPT_DIR}"/../../ --go_opt=paths=source_relative agenda.proto || exit 1
echo 'agenda.proto compiled'
