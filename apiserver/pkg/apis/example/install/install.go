/*
Copyright 2016 The Kubernetes Authors.

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

// Package install installs the certificates API group, making it available as
// an option to all of the API encoding/decoding machinery.
package install

import (
	"github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/announced"
	"github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/registered"
	"github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	"github.com/r2d4/kube-client/apiserver/pkg/apis/example"
	examplev1 "github.com/r2d4/kube-client/apiserver/pkg/apis/example/v1"
)

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  example.GroupName,
			VersionPreferenceOrder:     []string{examplev1.SchemeGroupVersion.Version},
			ImportPrefix:               "k8s.io/apiserver/pkg/apis/example",
			AddInternalObjectsToScheme: example.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			examplev1.SchemeGroupVersion.Version: examplev1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
