#!/bin/bash -eu

go install github.com/AdamKorcz/go-118-fuzz-build@latest
go get github.com/AdamKorcz/go-118-fuzz-build/testing

export IMPORT_PATH="github.com/clfs/fuzz"

# TODO: Locate and compile all fuzz targets automatically.
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzBMPDecode image_FuzzBMPDecode
compile_native_go_fuzzer $IMPORT_PATH/projects/image FuzzWEBPDecode image_FuzzWEBPDecode
compile_native_go_fuzzer $IMPORT_PATH/projects/jsoniter FuzzParseBytes jsoniter_FuzzParseBytes
compile_native_go_fuzzer $IMPORT_PATH/projects/jsoniter FuzzUnmarshal jsoniter_FuzzUnmarshal
compile_native_go_fuzzer $IMPORT_PATH/projects/rclone FuzzFilenameRoundTrip rclone_FuzzFilenameRoundTrip
compile_native_go_fuzzer $IMPORT_PATH/projects/rclone FuzzFspathParse rclone_FuzzFspathParse
compile_native_go_fuzzer $IMPORT_PATH/projects/syncthing FuzzSanitizePath syncthing_FuzzSanitizePath

# graphviz_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzGraphviz FuzzGraphviz

# pem_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzPEM FuzzPEM

# protobuf_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzProtobuf FuzzProtobuf
compile_native_go_fuzzer $IMPORT_PATH FuzzProtobufJSON FuzzProtobufJSON
#compile_native_go_fuzzer $IMPORT_PATH FuzzProtobufText FuzzProtobufText

# sfnt/sfnt_test.go
compile_native_go_fuzzer $IMPORT_PATH/sfnt FuzzParse sfnt_FuzzParse

# tiff_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzTIFF FuzzTIFF

# toml_test.go
compile_native_go_fuzzer $IMPORT_PATH FuzzTOML_BurntSushi FuzzTOML_BurntSushi
compile_native_go_fuzzer $IMPORT_PATH FuzzTOML_pelletier FuzzTOML_pelletier
