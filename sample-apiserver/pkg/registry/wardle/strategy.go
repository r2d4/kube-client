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

package wardle

import (
	"fmt"

	"github.com/r2d4/kube-client/apimachinery/pkg/fields"
	"github.com/r2d4/kube-client/apimachinery/pkg/labels"
	"github.com/r2d4/kube-client/apimachinery/pkg/runtime"
	"github.com/r2d4/kube-client/apimachinery/pkg/util/validation/field"
	genericapirequest "github.com/r2d4/kube-client/apiserver/pkg/endpoints/request"
	"github.com/r2d4/kube-client/apiserver/pkg/registry/generic"
	"github.com/r2d4/kube-client/apiserver/pkg/storage"
	"github.com/r2d4/kube-client/apiserver/pkg/storage/names"

	"k8s.io/sample-apiserver/pkg/apis/wardle"
)

type apiServerStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func NewStrategy(typer runtime.ObjectTyper) apiServerStrategy {
	return apiServerStrategy{typer, names.SimpleNameGenerator}
}

func (apiServerStrategy) NamespaceScoped() bool {
	return false
}

func (apiServerStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (apiServerStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (apiServerStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
	// return validation.ValidateFlunder(obj.(*wardle.Flunder))
}

func (apiServerStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (apiServerStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (apiServerStrategy) Canonicalize(obj runtime.Object) {
}

func (apiServerStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
	// return validation.ValidateFlunderUpdate(obj.(*wardle.Flunder), old.(*wardle.Flunder))
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*wardle.Flunder)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Flunder.")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), FlunderToSelectableFields(apiserver), nil
}

// MatchFlunder is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchFlunder(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// FlunderToSelectableFields returns a field set that represents the object.
func FlunderToSelectableFields(obj *wardle.Flunder) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}
