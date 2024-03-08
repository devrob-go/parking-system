#!/usr/bin/env bash

set -euo pipefail

source ./devenv.sh

pushd "$tools_dir"

# Relying on pinned version in go.mod due to bug in `go install`
# https://github.com/golang/go/issues/54908
GOBIN=$gunk_dir/bin go install github.com/gunk/gunk
GOBIN=$gunk_dir/bin go install mvdan.cc/gofumpt/gofumports@v0.1.1
GOBIN=$gunk_dir/bin go install github.com/gunk/scopegen@v0.1.4
GOBIN=$gunk_dir/bin go install cirello.io/openapigen@v1.1.0
GOBIN=$gunk_dir/bin go install github.com/go-swagger/go-swagger/cmd/swagger@v0.27.0
GOBIN=$gunk_dir/bin go install github.com/akuity/grpc-gateway-client/protoc-gen-grpc-gateway-client@main

pushd "$tools_dir/pgunk"
GOBIN=$gunk_dir/bin go install
popd

pushd "$gunk_dir"
pnpm install --frozen-lockfile
popd
