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
	v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	labels "github.com/r2d4/kube-client/apimachinery/pkg/labels"
	schema "github.com/r2d4/kube-client/apimachinery/pkg/runtime/schema"
	types "github.com/r2d4/kube-client/apimachinery/pkg/types"
	watch "github.com/r2d4/kube-client/apimachinery/pkg/watch"
	v1beta1 "github.com/r2d4/kube-client/client-go/pkg/apis/extensions/v1beta1"
	testing "github.com/r2d4/kube-client/client-go/testing"
)

// FakeDaemonSets implements DaemonSetInterface
type FakeDaemonSets struct {
	Fake *FakeExtensionsV1beta1
	ns   string
}

var daemonsetsResource = schema.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "daemonsets"}

func (c *FakeDaemonSets) Create(daemonSet *v1beta1.DaemonSet) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daemonsetsResource, c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

func (c *FakeDaemonSets) Update(daemonSet *v1beta1.DaemonSet) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daemonsetsResource, c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

func (c *FakeDaemonSets) UpdateStatus(daemonSet *v1beta1.DaemonSet) (*v1beta1.DaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(daemonsetsResource, "status", c.ns, daemonSet), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

func (c *FakeDaemonSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(daemonsetsResource, c.ns, name), &v1beta1.DaemonSet{})

	return err
}

func (c *FakeDaemonSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daemonsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.DaemonSetList{})
	return err
}

func (c *FakeDaemonSets) Get(name string, options v1.GetOptions) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daemonsetsResource, c.ns, name), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}

func (c *FakeDaemonSets) List(opts v1.ListOptions) (result *v1beta1.DaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daemonsetsResource, c.ns, opts), &v1beta1.DaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DaemonSetList{}
	for _, item := range obj.(*v1beta1.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daemonSets.
func (c *FakeDaemonSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daemonsetsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched daemonSet.
func (c *FakeDaemonSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daemonsetsResource, c.ns, name, data, subresources...), &v1beta1.DaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DaemonSet), err
}
