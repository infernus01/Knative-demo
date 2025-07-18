/*
Copyright 2024 The Namespace Cleaner Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	clusteropsv1alpha1 "github.com/infernus01/knative-demo/pkg/apis/clusterops/v1alpha1"
	scheme "github.com/infernus01/knative-demo/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NamespaceCleanersGetter has a method to return a NamespaceCleanerInterface.
// A group's client should implement this interface.
type NamespaceCleanersGetter interface {
	NamespaceCleaners() NamespaceCleanerInterface
}

// NamespaceCleanerInterface has methods to work with NamespaceCleaner resources.
type NamespaceCleanerInterface interface {
	Create(ctx context.Context, namespaceCleaner *clusteropsv1alpha1.NamespaceCleaner, opts v1.CreateOptions) (*clusteropsv1alpha1.NamespaceCleaner, error)
	Update(ctx context.Context, namespaceCleaner *clusteropsv1alpha1.NamespaceCleaner, opts v1.UpdateOptions) (*clusteropsv1alpha1.NamespaceCleaner, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*clusteropsv1alpha1.NamespaceCleaner, error)
	List(ctx context.Context, opts v1.ListOptions) (*clusteropsv1alpha1.NamespaceCleanerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusteropsv1alpha1.NamespaceCleaner, err error)
	NamespaceCleanerExpansion
}

// namespaceCleaners implements NamespaceCleanerInterface
type namespaceCleaners struct {
	*gentype.ClientWithList[*clusteropsv1alpha1.NamespaceCleaner, *clusteropsv1alpha1.NamespaceCleanerList]
}

// newNamespaceCleaners returns a NamespaceCleaners
func newNamespaceCleaners(c *ClusteropsV1alpha1Client) *namespaceCleaners {
	return &namespaceCleaners{
		gentype.NewClientWithList[*clusteropsv1alpha1.NamespaceCleaner, *clusteropsv1alpha1.NamespaceCleanerList](
			"namespacecleaners",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *clusteropsv1alpha1.NamespaceCleaner { return &clusteropsv1alpha1.NamespaceCleaner{} },
			func() *clusteropsv1alpha1.NamespaceCleanerList { return &clusteropsv1alpha1.NamespaceCleanerList{} },
		),
	}
}
