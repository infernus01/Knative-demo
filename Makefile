# Basic Makefile for NamespaceCleaner CRD

.PHONY: install-kind apply-crd apply-cr setup clean

# Install kind cluster
install-kind:
	@echo "Creating kind cluster..."
	kind create cluster --name namespacecleaner-demo

# Apply the CRD
apply-crd:
	@echo "Applying CRD..."
	kubectl apply -f config/crd/namespacecleaners.yaml

# Apply the Custom Resource
apply-cr:
	@echo "Applying Custom Resource..."
	kubectl apply -f example-namespacecleaner.yaml

# Full setup: install kind + apply CRD + apply CR
setup: install-kind apply-crd apply-cr
	@echo "Setup complete!"
	@kubectl get namespacecleaners

# Clean up
clean:
	@echo "Deleting kind cluster..."
	kind delete cluster --name namespacecleaner-demo 