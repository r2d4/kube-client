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

package autoscaling

import (
	v1 "github.com/r2d4/kube-client/client-go/informers/autoscaling/v1"
	v2alpha1 "github.com/r2d4/kube-client/client-go/informers/autoscaling/v2alpha1"
	internalinterfaces "github.com/r2d4/kube-client/client-go/informers/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
	// V2alpha1 provides access to shared informers for resources in V2alpha1.
	V2alpha1() v2alpha1.Interface
}

type group struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &group{f}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.SharedInformerFactory)
}

// V2alpha1 returns a new v2alpha1.Interface.
func (g *group) V2alpha1() v2alpha1.Interface {
	return v2alpha1.New(g.SharedInformerFactory)
}
