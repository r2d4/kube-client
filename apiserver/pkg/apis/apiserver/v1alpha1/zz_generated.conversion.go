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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "github.com/r2d4/kube-client/apimachinery/pkg/conversion"
	runtime "github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	apiserver "github.com/r2d4/kube-client/apiserver/pkg/apis/apiserver"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration,
		Convert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration,
		Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration,
		Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration,
	)
}

func autoConvert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]apiserver.AdmissionPluginConfiguration, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Plugins = nil
	}
	return nil
}

func Convert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]AdmissionPluginConfiguration, len(*in))
		for i := range *in {
			if err := Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Plugins = nil
	}
	return nil
}

func Convert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionConfiguration_To_v1alpha1_AdmissionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Configuration, &out.Configuration, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Configuration, &out.Configuration, s); err != nil {
		return err
	}
	return nil
}

func Convert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionPluginConfiguration_To_v1alpha1_AdmissionPluginConfiguration(in, out, s)
}
