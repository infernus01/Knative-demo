#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(go env GOPATH)/pkg/mod/k8s.io/code-generator@v0.32.5}

source "${CODEGEN_PKG}"/kube_codegen.sh

kube::codegen::gen_helpers \
  --boilerplate "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  "${SCRIPT_ROOT}"/pkg/apis

kube::codegen::gen_client \
  --with-watch \
  --output-dir "${SCRIPT_ROOT}"/pkg/generated \
  --output-pkg github.com/infernus01/knative-demo/pkg/generated \
  --boilerplate "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  "${SCRIPT_ROOT}"/pkg/apis
