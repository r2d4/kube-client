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

package v1beta1

import (
	v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	watch "github.com/r2d4/kube-client/apimachinery/pkg/watch"
	internalinterfaces "github.com/r2d4/kube-client/client-go/informers/internalinterfaces"
	kubernetes "github.com/r2d4/kube-client/client-go/kubernetes"
	v1beta1 "github.com/r2d4/kube-client/client-go/listers/rbac/v1beta1"
	rbac_v1beta1 "github.com/r2d4/kube-client/client-go/pkg/apis/rbac/v1beta1"
	cache "github.com/r2d4/kube-client/client-go/tools/cache"
	time "time"
)

// RoleBindingInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.RoleBindingLister
}

type roleBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newRoleBindingInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.RbacV1beta1().RoleBindings(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.RbacV1beta1().RoleBindings(v1.NamespaceAll).Watch(options)
			},
		},
		&rbac_v1beta1.RoleBinding{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbac_v1beta1.RoleBinding{}, newRoleBindingInformer)
}

func (f *roleBindingInformer) Lister() v1beta1.RoleBindingLister {
	return v1beta1.NewRoleBindingLister(f.Informer().GetIndexer())
}
