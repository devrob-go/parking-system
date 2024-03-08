#!/usr/bin/env bash
set -ex

# fail on OS X old bash version (for **)
shopt -s globstar
set -euo pipefail

./install_deps.sh

source ./devenv.sh

if [[ -d ../vendor ]]; then rm -r ../vendor; fi

gunk download protoc --version v3.12.0

gunk format ./...
# generate the remaining packages, avoiding contention on dependencies.
pgunk $(find ./api/v1 -mindepth 1 -maxdepth 1 -type d)

gofumports -w $(find . -iname '*.go')
