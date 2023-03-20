#!/bin/bash -eu

go install github.com/AdamKorcz/go-118-fuzz-build@latest
go get github.com/AdamKorcz/go-118-fuzz-build/testing

export IMPORT_PATH="github.com/clfs/fuzz"

# TODO: Locate and compile all fuzz targets automatically.
compile_native_go_fuzzer $IMPORT_PATH/projects/syncthing FuzzSanitizePath syncthing_FuzzSanitizePath
compile_native_go_fuzzer $IMPORT_PATH/projects/protobuf FuzzUnmarshal protobuf_FuzzUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/projects/protobuf FuzzUnmarshalDeprecated protobuf_FuzzUnmarshalDeprecated
