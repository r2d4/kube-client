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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	"github.com/r2d4/kube-client/apimachinery/pkg/api/errors"
	v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	"github.com/r2d4/kube-client/apimachinery/pkg/labels"
	rbac "github.com/r2d4/kube-client/client-go/pkg/apis/rbac"
	v1alpha1 "github.com/r2d4/kube-client/client-go/pkg/apis/rbac/v1alpha1"
	"github.com/r2d4/kube-client/client-go/tools/cache"
)

// ClusterRoleLister helps list ClusterRoles.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRole, err error)
	// Get retrieves the ClusterRole from the index for a given name.
	Get(name string) (*v1alpha1.ClusterRole, error)
	ClusterRoleListerExpansion
}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	indexer cache.Indexer
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(indexer cache.Indexer) ClusterRoleLister {
	return &clusterRoleLister{indexer: indexer}
}

// List lists all ClusterRoles in the indexer.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the index for a given name.
func (s *clusterRoleLister) Get(name string) (*v1alpha1.ClusterRole, error) {
	key := &v1alpha1.ClusterRole{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbac.Resource("clusterrole"), name)
	}
	return obj.(*v1alpha1.ClusterRole), nil
}
