#!/bin/bash

# This script generates the clientset, listers, and informers for the NamespaceCleaner CRD
# It's a basic script suitable for beginners to understand Kubernetes code generation

set -o errexit
set -o nounset
set -o pipefail

# Get the script's directory
SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

# Find the code-generator location
CODEGEN_PKG=${CODEGEN_PKG:-$(go list -m -f '{{.Dir}}' k8s.io/code-generator 2>/dev/null)}

if [ -z "$CODEGEN_PKG" ]; then
    echo "Error: Could not find k8s.io/code-generator"
    echo "Make sure it's installed by running: go mod download"
    exit 1
fi

echo "Using code-generator from: $CODEGEN_PKG"

# Load the code generation functions
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "Generating deepcopy methods..."
# Generate deepcopy methods for all types
kube::codegen::gen_helpers \
  --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  "${SCRIPT_ROOT}/pkg/apis"

echo "Generating client, informers, and listers..."
# Generate client, informers, and listers
kube::codegen::gen_client \
  --with-watch \
  --output-dir "${SCRIPT_ROOT}/pkg/generated" \
  --output-pkg github.com/infernus01/knative-demo/pkg/generated \
  --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
  "${SCRIPT_ROOT}/pkg/apis"

echo "✅ Code generation completed successfully!"
echo ""
echo "Generated files:"
echo "- pkg/generated/clientset/    - Client to interact with NamespaceCleaner resources"
echo "- pkg/generated/listers/      - Listers for efficient resource queries"
echo "- pkg/generated/informers/    - Informers for watching resource changes"
echo "- pkg/apis/clusterops/v1alpha1/zz_generated.deepcopy.go - DeepCopy methods"