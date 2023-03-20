#!/bin/bash -eu

go install github.com/AdamKorcz/go-118-fuzz-build@latest
go get github.com/AdamKorcz/go-118-fuzz-build/testing

export IMPORT_PATH="github.com/clfs/fuzz"

# TODO: Locate and compile all fuzz targets automatically.
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzBMPDecode image_FuzzBMPDecode
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzOpenTypeParse image_FuzzOpenTypeParse
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzTIFFDecode image_FuzzTIFFDecode
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzWEBPDecode image_FuzzWEBPDecode
compile_native_go_fuzzer $IMPORT_PATH/projects/protobuf FuzzUnmarshal protobuf_FuzzUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/projects/protobuf FuzzUnmarshalDeprecated protobuf_FuzzUnmarshalDeprecated
compile_native_go_fuzzer $IMPORT_PATH/projects/syncthing FuzzSanitizePath syncthing_FuzzSanitizePath
