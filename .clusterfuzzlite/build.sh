#!/bin/bash -eu

go install github.com/AdamKorcz/go-118-fuzz-build@latest
go get github.com/AdamKorcz/go-118-fuzz-build/testing

export IMPORT_PATH="github.com/clfs/fuzz"

# TODO: Locate and compile all fuzz targets automatically.
compile_native_go_fuzzer $IMPORT_PATH/projects/jsoniter FuzzParseBytes jsoniter_FuzzParseBytes
compile_native_go_fuzzer $IMPORT_PATH/projects/jsoniter FuzzUnmarshal jsoniter_FuzzUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/projects/rclone FuzzFilenameRoundTrip rclone_FuzzFilenameRoundTrip
compile_native_go_fuzzer $IMPORT_PATH/projects/rclone FuzzFspathParse rclone_FuzzFspathParse
compile_native_go_fuzzer $IMPORT_PATH/projects/syncthing FuzzSanitizePath syncthing_FuzzSanitizePath

# graphviz_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzGraphviz FuzzGraphviz

# pem_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzPEM FuzzPEM

compile_native_go_fuzzer $IMPORT_PATH/protobuf FuzzProtoUnmarshal protobuf_FuzzProtoUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/protobuf FuzzProtojsonUnmarshal protobuf_FuzzProtojsonUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/protobuf FuzzPrototextUnmarshal protobuf_FuzzPrototextUnmarshal

compile_native_go_fuzzer $IMPORT_PATH/tiff FuzzDecode tiff_FuzzDecode
compile_native_go_fuzzer $IMPORT_PATH/sfnt FuzzParse sfnt_FuzzParse

# toml_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzTOML_BurntSushi FuzzTOML_BurntSushi
compile_native_go_fuzzer $IMPORT_PATH FuzzTOML_pelletier FuzzTOML_pelletier

compile_native_go_fuzzer $IMPORT_PATH/bmp FuzzDecode bmp_FuzzDecode

compile_native_go_fuzzer $IMPORT_PATH/webp FuzzDecode webp_FuzzDecode
