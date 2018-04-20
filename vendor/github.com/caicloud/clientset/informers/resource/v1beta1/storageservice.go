/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1beta1 "github.com/caicloud/clientset/listers/resource/v1beta1"
	resource_v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	client_go_kubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// StorageServiceInformer provides access to a shared informer and lister for
// StorageServices.
type StorageServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.StorageServiceLister
}

type storageServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageServiceInformer constructs a new informer for StorageService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageServiceInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageServiceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageServiceInformer constructs a new informer for StorageService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageServiceInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().StorageServices().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1beta1().StorageServices().Watch(options)
			},
		},
		&resource_v1beta1.StorageService{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageServiceInformer) defaultInformer(client client_go_kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageServiceInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resource_v1beta1.StorageService{}, f.defaultInformer)
}

func (f *storageServiceInformer) Lister() v1beta1.StorageServiceLister {
	return v1beta1.NewStorageServiceLister(f.Informer().GetIndexer())
}
