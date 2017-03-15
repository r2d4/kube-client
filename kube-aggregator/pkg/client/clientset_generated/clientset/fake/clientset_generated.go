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

package fake

import (
	"github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	"github.com/r2d4/kube-client/apimachinery/pkg/watch"
	"github.com/r2d4/kube-client/client-go/discovery"
	fakediscovery "github.com/r2d4/kube-client/client-go/discovery/fake"
	"github.com/r2d4/kube-client/client-go/testing"
	clientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	apiregistrationv1alpha1 "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1alpha1"
	fakeapiregistrationv1alpha1 "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/apiregistration/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(registry, scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o, registry.RESTMapper()))

	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return &fakediscovery.FakeDiscovery{Fake: &c.Fake}
}

var _ clientset.Interface = &Clientset{}

// ApiregistrationV1alpha1 retrieves the ApiregistrationV1alpha1Client
func (c *Clientset) ApiregistrationV1alpha1() apiregistrationv1alpha1.ApiregistrationV1alpha1Interface {
	return &fakeapiregistrationv1alpha1.FakeApiregistrationV1alpha1{Fake: &c.Fake}
}

// Apiregistration retrieves the ApiregistrationV1alpha1Client
func (c *Clientset) Apiregistration() apiregistrationv1alpha1.ApiregistrationV1alpha1Interface {
	return &fakeapiregistrationv1alpha1.FakeApiregistrationV1alpha1{Fake: &c.Fake}
}
