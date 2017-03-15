/*
Copyright 2014 The Kubernetes Authors.

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

// Package install installs the v1 monolithic api, making it available as an
// option to all of the API encoding/decoding machinery.
package install

import (
	"github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/announced"
	"github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/registered"
	"github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	"github.com/r2d4/kube-client/apimachinery/pkg/util/sets"
	"github.com/r2d4/kube-client/client-go/pkg/api"
	"github.com/r2d4/kube-client/client-go/pkg/api/v1"
)

func init() {
	Install(api.GroupFactoryRegistry, api.Registry, api.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  api.GroupName,
			VersionPreferenceOrder:     []string{v1.SchemeGroupVersion.Version},
			ImportPrefix:               "k8s.io/client-go/pkg/api",
			AddInternalObjectsToScheme: api.AddToScheme,
			RootScopedKinds: sets.NewString(
				"Node",
				"Namespace",
				"PersistentVolume",
				"ComponentStatus",
			),
			IgnoredKinds: sets.NewString(
				"ListOptions",
				"DeleteOptions",
				"Status",
				"PodLogOptions",
				"PodExecOptions",
				"PodAttachOptions",
				"PodPortForwardOptions",
				"PodProxyOptions",
				"NodeProxyOptions",
				"ServiceProxyOptions",
				"ThirdPartyResource",
				"ThirdPartyResourceData",
				"ThirdPartyResourceList",
			),
		},
		announced.VersionToSchemeFunc{
			v1.SchemeGroupVersion.Version: v1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}