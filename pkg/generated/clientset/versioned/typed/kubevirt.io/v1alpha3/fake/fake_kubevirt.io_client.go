/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1alpha3 "github.com/rancher/harvester/pkg/generated/clientset/versioned/typed/kubevirt.io/v1alpha3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubevirtV1alpha3 struct {
	*testing.Fake
}

func (c *FakeKubevirtV1alpha3) KubeVirts(namespace string) v1alpha3.KubeVirtInterface {
	return &FakeKubeVirts{c, namespace}
}

func (c *FakeKubevirtV1alpha3) VirtualMachines(namespace string) v1alpha3.VirtualMachineInterface {
	return &FakeVirtualMachines{c, namespace}
}

func (c *FakeKubevirtV1alpha3) VirtualMachineInstances(namespace string) v1alpha3.VirtualMachineInstanceInterface {
	return &FakeVirtualMachineInstances{c, namespace}
}

func (c *FakeKubevirtV1alpha3) VirtualMachineInstanceMigrations(namespace string) v1alpha3.VirtualMachineInstanceMigrationInterface {
	return &FakeVirtualMachineInstanceMigrations{c, namespace}
}

func (c *FakeKubevirtV1alpha3) VirtualMachineInstancePresets(namespace string) v1alpha3.VirtualMachineInstancePresetInterface {
	return &FakeVirtualMachineInstancePresets{c, namespace}
}

func (c *FakeKubevirtV1alpha3) VirtualMachineInstanceReplicaSets(namespace string) v1alpha3.VirtualMachineInstanceReplicaSetInterface {
	return &FakeVirtualMachineInstanceReplicaSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubevirtV1alpha3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
