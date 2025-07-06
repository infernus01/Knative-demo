# NamespaceCleaner CRD Guide

This is a **beginner-friendly** guide for working with the NamespaceCleaner Custom Resource Definition (CRD).

## What is this CRD?

The NamespaceCleaner CRD allows you to automatically clean up namespaces in your Kubernetes cluster based on a schedule and label selector.

## 📋 Prerequisites

- Kubernetes cluster (kind, minikube, or any K8s cluster)
- `kubectl` installed and configured
- Go 1.24.4 or later (for code generation)

## 🚀 Step-by-Step Setup

### Step 1: Generate the Code

```bash
# Generate the client, informers, and listers
bash hack/update-codegen.sh
```

This will generate:
- `pkg/generated/clientset/` - Client to interact with NamespaceCleaner resources
- `pkg/generated/listers/` - Listers for efficient resource queries
- `pkg/generated/informers/` - Informers for watching resource changes
- `pkg/apis/clusterops/v1alpha1/zz_generated.deepcopy.go` - DeepCopy methods

### Step 2: Install the CRD

```bash
# Install the CRD into your cluster
kubectl apply -f config/crd/bases/clusterops.io_namespacecleaners.yaml
```

### Step 3: Verify the CRD

```bash
# Check if the CRD is installed
kubectl get crds | grep namespacecleaners

# You should see: namespacecleaners.clusterops.io
```

### Step 4: Create a NamespaceCleaner Resource

```bash
# Apply the example resource
kubectl apply -f config/samples/clusterops_v1alpha1_namespacecleaner.yaml
```

### Step 5: Verify Your Resource

```bash
# Check your NamespaceCleaner resource
kubectl get namespacecleaners

# Or use the short name
kubectl get nsc

# Get detailed information
kubectl describe namespacecleaner example-namespacecleaner
```

## 📝 Understanding the Resource Structure

### Spec Fields:
- `schedule`: Cron format string (e.g., "0 2 * * *" = every day at 2 AM)
- `maxAge`: How old namespaces should be before cleanup (e.g., "7d", "24h", "30m")
- `selector`: Label selector to match namespaces to clean up

### Status Fields:
- `lastCleanup`: Timestamp of the last cleanup
- `nextCleanup`: Timestamp of the next scheduled cleanup
- `conditions`: Current status conditions

## 🔧 Common Commands

```bash
# List all NamespaceCleaner resources
kubectl get namespacecleaners -A

# Edit a NamespaceCleaner
kubectl edit namespacecleaner example-namespacecleaner

# Delete a NamespaceCleaner
kubectl delete namespacecleaner example-namespacecleaner

# Delete the CRD (this will also delete all NamespaceCleaner resources!)
kubectl delete crd namespacecleaners.clusterops.io
```

## 🎯 What's Next?

After completing this setup, you have:

1. ✅ Defined a CRD Go type with proper annotations
2. ✅ Generated client, informer, and lister code
3. ✅ Installed the CRD in your cluster
4. ✅ Created and verified a custom resource

**Next steps for learning:**
- Build a controller that watches NamespaceCleaner resources
- Implement the actual cleanup logic
- Add validation webhooks
- Add more advanced features like status updates

## 🐛 Troubleshooting

### Code Generation Issues
```bash
# Make sure dependencies are downloaded
go mod download

# Check if code-generator is available
go list -m k8s.io/code-generator
```

### CRD Installation Issues
```bash
# Check if you have proper permissions
kubectl auth can-i create customresourcedefinitions

# Check CRD status
kubectl get crd namespacecleaners.clusterops.io -o yaml
```

## 📚 Learn More

- [Kubernetes CRD Documentation](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
- [Code Generator Documentation](https://github.com/kubernetes/code-generator)
- [Writing Controllers](https://book.kubebuilder.io/) 