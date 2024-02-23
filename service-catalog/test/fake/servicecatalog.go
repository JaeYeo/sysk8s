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
	rest "k8s.io/client-go/rest"

	servicecatalogv1 "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1"
	v1 "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1"
)

// Servicecatalogv1 is a wrapper around the generated fake service catalog
// that clones the ServiceInstance and ServiceBinding objects being
// passed to UpdateStatus. This is a workaround until the generated fake clientset
// does its own copying.
type Servicecatalogv1 struct {
	servicecatalogv1.Servicecatalogv1Interface
}

var _ servicecatalogv1.Servicecatalogv1Interface = &Servicecatalogv1{}

func (c *Servicecatalogv1) ClusterServiceBrokers() v1.ClusterServiceBrokerInterface {
	return c.Servicecatalogv1Interface.ClusterServiceBrokers()
}

func (c *Servicecatalogv1) ClusterServiceClasses() v1.ClusterServiceClassInterface {
	return c.Servicecatalogv1Interface.ClusterServiceClasses()
}

func (c *Servicecatalogv1) ServiceInstances(namespace string) v1.ServiceInstanceInterface {
	serviceInstances := c.Servicecatalogv1Interface.ServiceInstances(namespace)
	return &ServiceInstances{serviceInstances}
}

func (c *Servicecatalogv1) ServiceBindings(namespace string) v1.ServiceBindingInterface {
	serviceBindings := c.Servicecatalogv1Interface.ServiceBindings(namespace)
	return &ServiceBindings{serviceBindings}
}

func (c *Servicecatalogv1) ClusterServicePlans() v1.ClusterServicePlanInterface {
	return c.Servicecatalogv1Interface.ClusterServicePlans()
}

func (c *Servicecatalogv1) RESTClient() rest.Interface {
	return c.Servicecatalogv1Interface.RESTClient()
}
