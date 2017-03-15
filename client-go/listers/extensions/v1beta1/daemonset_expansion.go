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

package v1beta1

import (
	"fmt"

	metav1 "github.com/r2d4/kube-client/apimachinery/pkg/apis/meta/v1"
	"github.com/r2d4/kube-client/apimachinery/pkg/labels"
	"github.com/r2d4/kube-client/client-go/pkg/api/v1"
	"github.com/r2d4/kube-client/client-go/pkg/apis/extensions/v1beta1"
)

// DaemonSetListerExpansion allows custom methods to be added to
// DaemonSetLister.
type DaemonSetListerExpansion interface {
	GetPodDaemonSets(pod *v1.Pod) ([]*v1beta1.DaemonSet, error)
}

// DaemonSetNamespaceListerExpansion allows custom methods to be added to
// DaemonSetNamespaeLister.
type DaemonSetNamespaceListerExpansion interface{}

// GetPodDaemonSets returns a list of DaemonSets that potentially match a pod.
// Only the one specified in the Pod's ControllerRef will actually manage it.
// Returns an error only if no matching DaemonSets are found.
func (s *daemonSetLister) GetPodDaemonSets(pod *v1.Pod) ([]*v1beta1.DaemonSet, error) {
	var selector labels.Selector
	var daemonSet *v1beta1.DaemonSet

	if len(pod.Labels) == 0 {
		return nil, fmt.Errorf("no daemon sets found for pod %v because it has no labels", pod.Name)
	}

	list, err := s.DaemonSets(pod.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var daemonSets []*v1beta1.DaemonSet
	for i := range list {
		daemonSet = list[i]
		if daemonSet.Namespace != pod.Namespace {
			continue
		}
		selector, err = metav1.LabelSelectorAsSelector(daemonSet.Spec.Selector)
		if err != nil {
			// this should not happen if the DaemonSet passed validation
			return nil, err
		}

		// If a daemonSet with a nil or empty selector creeps in, it should match nothing, not everything.
		if selector.Empty() || !selector.Matches(labels.Set(pod.Labels)) {
			continue
		}
		daemonSets = append(daemonSets, daemonSet)
	}

	if len(daemonSets) == 0 {
		return nil, fmt.Errorf("could not find daemon set for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}

	return daemonSets, nil
}
