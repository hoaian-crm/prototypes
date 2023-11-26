#!/bin/bash

# Go compilte
protoc --plugin ~/go/bin/protoc-gen-go --go_out gen/go --proto_path=. interfaces/*.proto

# Ts Compile
protoc --plugin /Users/tranhoaian/.nvm/versions/node/v14.21.2/lib/node_modules/ts-proto/protoc-gen-ts_proto --ts_proto_out=gen/ts --proto_path=. interfaces/*.proto