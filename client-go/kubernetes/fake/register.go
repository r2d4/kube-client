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
	announced "github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/announced"
	registered "github.com/r2d4/kube-client/apimachinery/pkg/apimachinery/registered"
	v1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	schema "github.com/r2d4/kube-client/apimachinery/pkg/runtime/schema"
	serializer "github.com/r2d4/kube-client/apimachinery/pkg/runtime/serializer"
	core "github.com/r2d4/kube-client/client-go/pkg/api/install"
	apps "github.com/r2d4/kube-client/client-go/pkg/apis/apps/install"
	authentication "github.com/r2d4/kube-client/client-go/pkg/apis/authentication/install"
	authorization "github.com/r2d4/kube-client/client-go/pkg/apis/authorization/install"
	autoscaling "github.com/r2d4/kube-client/client-go/pkg/apis/autoscaling/install"
	batch "github.com/r2d4/kube-client/client-go/pkg/apis/batch/install"
	certificates "github.com/r2d4/kube-client/client-go/pkg/apis/certificates/install"
	extensions "github.com/r2d4/kube-client/client-go/pkg/apis/extensions/install"
	policy "github.com/r2d4/kube-client/client-go/pkg/apis/policy/install"
	rbac "github.com/r2d4/kube-client/client-go/pkg/apis/rbac/install"
	settings "github.com/r2d4/kube-client/client-go/pkg/apis/settings/install"
	storage "github.com/r2d4/kube-client/client-go/pkg/apis/storage/install"
	os "os"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

var registry = registered.NewOrDie(os.Getenv("KUBE_API_VERSIONS"))
var groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	Install(groupFactoryRegistry, registry, scheme)
}

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	core.Install(groupFactoryRegistry, registry, scheme)
	apps.Install(groupFactoryRegistry, registry, scheme)
	authentication.Install(groupFactoryRegistry, registry, scheme)
	authorization.Install(groupFactoryRegistry, registry, scheme)
	autoscaling.Install(groupFactoryRegistry, registry, scheme)
	batch.Install(groupFactoryRegistry, registry, scheme)
	certificates.Install(groupFactoryRegistry, registry, scheme)
	extensions.Install(groupFactoryRegistry, registry, scheme)
	policy.Install(groupFactoryRegistry, registry, scheme)
	rbac.Install(groupFactoryRegistry, registry, scheme)
	settings.Install(groupFactoryRegistry, registry, scheme)
	storage.Install(groupFactoryRegistry, registry, scheme)

}
