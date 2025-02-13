/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicecatalogv1 struct {
	*testing.Fake
}

func (c *FakeServicecatalogv1) ClusterServiceBrokers() v1.ClusterServiceBrokerInterface {
	return &FakeClusterServiceBrokers{c}
}

func (c *FakeServicecatalogv1) ClusterServiceClasses() v1.ClusterServiceClassInterface {
	return &FakeClusterServiceClasses{c}
}

func (c *FakeServicecatalogv1) ClusterServicePlans() v1.ClusterServicePlanInterface {
	return &FakeClusterServicePlans{c}
}

func (c *FakeServicecatalogv1) ServiceBindings(namespace string) v1.ServiceBindingInterface {
	return &FakeServiceBindings{c, namespace}
}

func (c *FakeServicecatalogv1) ServiceBrokers(namespace string) v1.ServiceBrokerInterface {
	return &FakeServiceBrokers{c, namespace}
}

func (c *FakeServicecatalogv1) ServiceClasses(namespace string) v1.ServiceClassInterface {
	return &FakeServiceClasses{c, namespace}
}

func (c *FakeServicecatalogv1) ServiceInstances(namespace string) v1.ServiceInstanceInterface {
	return &FakeServiceInstances{c, namespace}
}

func (c *FakeServicecatalogv1) ServicePlans(namespace string) v1.ServicePlanInterface {
	return &FakeServicePlans{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicecatalogv1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
