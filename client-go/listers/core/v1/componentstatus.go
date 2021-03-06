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

package v1

import (
	"github.com/r2d4/kube-client/apimachinery/pkg/api/errors"
	meta_v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	"github.com/r2d4/kube-client/apimachinery/pkg/labels"
	api "github.com/r2d4/kube-client/client-go/pkg/api"
	v1 "github.com/r2d4/kube-client/client-go/pkg/api/v1"
	"github.com/r2d4/kube-client/client-go/tools/cache"
)

// ComponentStatusLister helps list ComponentStatuses.
type ComponentStatusLister interface {
	// List lists all ComponentStatuses in the indexer.
	List(selector labels.Selector) (ret []*v1.ComponentStatus, err error)
	// Get retrieves the ComponentStatus from the index for a given name.
	Get(name string) (*v1.ComponentStatus, error)
	ComponentStatusListerExpansion
}

// componentStatusLister implements the ComponentStatusLister interface.
type componentStatusLister struct {
	indexer cache.Indexer
}

// NewComponentStatusLister returns a new ComponentStatusLister.
func NewComponentStatusLister(indexer cache.Indexer) ComponentStatusLister {
	return &componentStatusLister{indexer: indexer}
}

// List lists all ComponentStatuses in the indexer.
func (s *componentStatusLister) List(selector labels.Selector) (ret []*v1.ComponentStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ComponentStatus))
	})
	return ret, err
}

// Get retrieves the ComponentStatus from the index for a given name.
func (s *componentStatusLister) Get(name string) (*v1.ComponentStatus, error) {
	key := &v1.ComponentStatus{ObjectMeta: meta_v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("componentstatus"), name)
	}
	return obj.(*v1.ComponentStatus), nil
}
