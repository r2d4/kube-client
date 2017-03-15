// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	conversion "github.com/r2d4/kube-client/apimachinery/pkg/conversion"
	runtime "github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Flunder, InType: reflect.TypeOf(&Flunder{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_FlunderList, InType: reflect.TypeOf(&FlunderList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_FlunderSpec, InType: reflect.TypeOf(&FlunderSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_FlunderStatus, InType: reflect.TypeOf(&FlunderStatus{})},
	)
}

func DeepCopy_v1alpha1_Flunder(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Flunder)
		out := out.(*Flunder)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

func DeepCopy_v1alpha1_FlunderList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FlunderList)
		out := out.(*FlunderList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Flunder, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Flunder(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_FlunderSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FlunderSpec)
		out := out.(*FlunderSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1alpha1_FlunderStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*FlunderStatus)
		out := out.(*FlunderStatus)
		*out = *in
		return nil
	}
}
