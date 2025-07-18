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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apisclusteropsv1alpha1 "github.com/infernus01/knative-demo/pkg/apis/clusterops/v1alpha1"
	versioned "github.com/infernus01/knative-demo/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/infernus01/knative-demo/pkg/generated/informers/externalversions/internalinterfaces"
	clusteropsv1alpha1 "github.com/infernus01/knative-demo/pkg/generated/listers/clusterops/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NamespaceCleanerInformer provides access to a shared informer and lister for
// NamespaceCleaners.
type NamespaceCleanerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() clusteropsv1alpha1.NamespaceCleanerLister
}

type namespaceCleanerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNamespaceCleanerInformer constructs a new informer for NamespaceCleaner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceCleanerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespaceCleanerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceCleanerInformer constructs a new informer for NamespaceCleaner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceCleanerInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusteropsV1alpha1().NamespaceCleaners().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusteropsV1alpha1().NamespaceCleaners().Watch(context.TODO(), options)
			},
		},
		&apisclusteropsv1alpha1.NamespaceCleaner{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceCleanerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespaceCleanerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *namespaceCleanerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisclusteropsv1alpha1.NamespaceCleaner{}, f.defaultInformer)
}

func (f *namespaceCleanerInformer) Lister() clusteropsv1alpha1.NamespaceCleanerLister {
	return clusteropsv1alpha1.NewNamespaceCleanerLister(f.Informer().GetIndexer())
}
