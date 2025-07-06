# Makefile for NamespaceCleaner CRD
# This is a beginner-friendly Makefile with common commands

.PHONY: generate install-crd uninstall-crd apply-sample clean-samples help

# Generate code (clientset, informers, listers, deepcopy)
generate:
	@echo "🔄 Generating code..."
	bash hack/update-codegen.sh

# Install the CRD in the cluster
install-crd:
	@echo "🚀 Installing CRD..."
	kubectl apply -f config/crd/bases/clusterops.io_namespacecleaners.yaml
	@echo "✅ CRD installed successfully!"

# Uninstall the CRD from the cluster
uninstall-crd:
	@echo "🗑️  Uninstalling CRD..."
	kubectl delete -f config/crd/bases/clusterops.io_namespacecleaners.yaml
	@echo "✅ CRD uninstalled successfully!"

# Apply the sample NamespaceCleaner resource
apply-sample:
	@echo "📝 Applying sample resource..."
	kubectl apply -f config/samples/clusterops_v1alpha1_namespacecleaner.yaml
	@echo "✅ Sample resource applied successfully!"

# Delete the sample NamespaceCleaner resource
clean-samples:
	@echo "🧹 Cleaning up sample resources..."
	kubectl delete -f config/samples/clusterops_v1alpha1_namespacecleaner.yaml
	@echo "✅ Sample resources cleaned up!"

# Verify the CRD and resources
verify:
	@echo "🔍 Verifying CRD installation..."
	kubectl get crds | grep namespacecleaners || echo "❌ CRD not found"
	@echo "🔍 Listing NamespaceCleaner resources..."
	kubectl get namespacecleaners -A || echo "ℹ️  No NamespaceCleaner resources found"

# Complete setup (generate code + install CRD + apply sample)
setup: generate install-crd apply-sample
	@echo "🎉 Complete setup finished!"
	@echo "📖 Check the status with: make verify"

# Complete teardown (remove sample + uninstall CRD)
teardown: clean-samples uninstall-crd
	@echo "🧹 Complete teardown finished!"

# Show help
help:
	@echo "Available commands:"
	@echo "  generate      - Generate client, informers, and listers"
	@echo "  install-crd   - Install the CRD in the cluster"
	@echo "  uninstall-crd - Remove the CRD from the cluster"
	@echo "  apply-sample  - Apply the sample NamespaceCleaner resource"
	@echo "  clean-samples - Delete the sample NamespaceCleaner resource"
	@echo "  verify        - Verify CRD and resources"
	@echo "  setup         - Complete setup (generate + install + sample)"
	@echo "  teardown      - Complete teardown (remove sample + uninstall)"
	@echo "  help          - Show this help message" 