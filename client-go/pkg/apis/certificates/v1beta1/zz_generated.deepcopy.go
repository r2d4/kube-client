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

package v1beta1

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
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CertificateSigningRequest, InType: reflect.TypeOf(&CertificateSigningRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CertificateSigningRequestCondition, InType: reflect.TypeOf(&CertificateSigningRequestCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CertificateSigningRequestList, InType: reflect.TypeOf(&CertificateSigningRequestList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CertificateSigningRequestSpec, InType: reflect.TypeOf(&CertificateSigningRequestSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_CertificateSigningRequestStatus, InType: reflect.TypeOf(&CertificateSigningRequestStatus{})},
	)
}

func DeepCopy_v1beta1_CertificateSigningRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequest)
		out := out.(*CertificateSigningRequest)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_CertificateSigningRequestSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_CertificateSigningRequestStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_CertificateSigningRequestCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestCondition)
		out := out.(*CertificateSigningRequestCondition)
		*out = *in
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1beta1_CertificateSigningRequestList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestList)
		out := out.(*CertificateSigningRequestList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CertificateSigningRequest, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CertificateSigningRequest(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_CertificateSigningRequestSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestSpec)
		out := out.(*CertificateSigningRequestSpec)
		*out = *in
		if in.Request != nil {
			in, out := &in.Request, &out.Request
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		if in.Usages != nil {
			in, out := &in.Usages, &out.Usages
			*out = make([]KeyUsage, len(*in))
			copy(*out, *in)
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} else {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_CertificateSigningRequestStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CertificateSigningRequestStatus)
		out := out.(*CertificateSigningRequestStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]CertificateSigningRequestCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_CertificateSigningRequestCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Certificate != nil {
			in, out := &in.Certificate, &out.Certificate
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
