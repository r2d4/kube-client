/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1

import (
	meta_v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	watch "github.com/r2d4/kube-client/apimachinery/pkg/watch"
	internalinterfaces "github.com/r2d4/kube-client/client-go/informers/internalinterfaces"
	kubernetes "github.com/r2d4/kube-client/client-go/kubernetes"
	v1 "github.com/r2d4/kube-client/client-go/listers/core/v1"
	api_v1 "github.com/r2d4/kube-client/client-go/pkg/api/v1"
	cache "github.com/r2d4/kube-client/client-go/tools/cache"
	time "time"
)

// EndpointsInformer provides access to a shared informer and lister for
// Endpoints.
type EndpointsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EndpointsLister
}

type endpointsInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newEndpointsInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().Endpoints(meta_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().Endpoints(meta_v1.NamespaceAll).Watch(options)
			},
		},
		&api_v1.Endpoints{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *endpointsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api_v1.Endpoints{}, newEndpointsInformer)
}

func (f *endpointsInformer) Lister() v1.EndpointsLister {
	return v1.NewEndpointsLister(f.Informer().GetIndexer())
}
