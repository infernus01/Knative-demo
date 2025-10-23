Kubernetes CRD Demo

Building a **Kubernetes Custom Resource Definition (CRD)** following professional patterns.

## 🎯 What we'll Learn

- How to define custom resources using Go types
- How Kubernetes code generation works
- How to install and use custom resources
- Professional patterns used by major Kubernetes tools

## 📋 Prerequisites

- Kubernetes cluster (kind, minikube, or any K8s cluster)
- `kubectl` installed and configured

## 📋 The 3-Step Process

### **Step 1: Define the CRD Go Type**
Create Go structs that define your custom resource schema

### **Step 2: Generate Boilerplate Code**  
Use Kubernetes code generators to create client libraries automatically

### **Step 3: Install the CRD into the Cluster**
Register your custom resource type with Kubernetes

## k8s code generator
This project uses Kubernetes code-generator to generate clients, informers, and listers for the NamespaceCleaner CRD.
Starting with Kubernetes 1.30+ (code-generator v0.30.x+), generate-groups.sh has been removed.
To make code generation work, I have pin code-generator to v0.27.3 or lower:
```
go get k8s.io/code-generator@v0.27.3
go mod tidy
```

Then you can generate your CRD clients by running:

```
./hack/update-codegen.sh
```

## Quick Start

### 1. Generate the code
```bash
bash hack/update-codegen.sh
```

### 2. Install the CRD
```bash
kubectl apply -f config/crd/namespacecleaners.yaml
```

### 3. Create a NamespaceCleaner
```bash
kubectl apply -f example-namespacecleaner.yaml
```

### 4. Check it works
```bash
kubectl get namespacecleaners
```

## Files

- `pkg/apis/clusterops/v1alpha1/types.go` - Go type definitions
- `pkg/apis/clusterops/v1alpha1/register.go` - API registration
- `config/crd/namespacecleaners.yaml` - CRD definition
- `hack/update-codegen.sh` - Code generation script
