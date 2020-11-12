/*
Copyright 2018 The Kubernetes Authors.

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

package ovirt

import (
	"github.com/openshift/machine-api-operator/pkg/generated/clientset/versioned/typed/machine/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	OvirtIdAnnotationKey = "VmId"
	ProviderIDPrefix     = "ovirt://"
)

// ActuatorParams holds parameter information for Actuator
type ActuatorParams struct {
	Namespace      string
	Client         client.Client
	KubeClient     *kubernetes.Clientset
	Scheme         *runtime.Scheme
	MachinesClient v1beta1.MachineV1beta1Interface
	EventRecorder  record.EventRecorder
}
