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

package servicecatalog

import (
	"time"

	apiv1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1"
	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// SvcatClient is an interface containing the various actions in the svcat pkg lib
// This interface is then faked with Counterfeiter for the cmd/svcat unit tests
type SvcatClient interface {
	Bind(string, string, string, string, string, interface{}, map[string]string) (*apiv1.ServiceBinding, error)
	BindingParentHierarchy(*apiv1.ServiceBinding) (*apiv1.ServiceInstance, *apiv1.ClusterServiceClass, *apiv1.ClusterServicePlan, *apiv1.ClusterServiceBroker, error)
	DeleteBinding(string, string) error
	DeleteBindings([]types.NamespacedName) ([]types.NamespacedName, error)
	IsBindingFailed(*apiv1.ServiceBinding) bool
	IsBindingReady(*apiv1.ServiceBinding) bool
	RetrieveBinding(string, string) (*apiv1.ServiceBinding, error)
	RetrieveBindings(string) (*apiv1.ServiceBindingList, error)
	RetrieveBindingsByInstance(*apiv1.ServiceInstance) ([]apiv1.ServiceBinding, error)
	Unbind(string, string) ([]types.NamespacedName, error)
	WaitForBinding(string, string, time.Duration, *time.Duration) (*apiv1.ServiceBinding, error)
	RemoveBindingFinalizerByInstance(string, string) ([]types.NamespacedName, error)
	RemoveFinalizerForBindings([]types.NamespacedName) ([]types.NamespacedName, error)
	RemoveFinalizerForBinding(types.NamespacedName) error
	RemoveFinalizerForInstance(string, string) error

	Deregister(string, *ScopeOptions) error
	RetrieveBrokers(opts ScopeOptions) ([]Broker, error)
	RetrieveBrokerByID(string, ScopeOptions) (Broker, error)
	RetrieveBrokerByClass(*apiv1.ClusterServiceClass) (*apiv1.ClusterServiceBroker, error)
	Register(string, string, *RegisterOptions, *ScopeOptions) (Broker, error)
	Sync(string, ScopeOptions, int) error
	WaitForBroker(string, *ScopeOptions, time.Duration, *time.Duration) (Broker, error)

	RetrieveClasses(ScopeOptions, string) ([]Class, error)
	RetrieveClassByName(string, ScopeOptions) (Class, error)
	RetrieveClassByID(string, ScopeOptions) (Class, error)
	RetrieveClassByPlan(Plan) (Class, error)
	CreateClassFrom(CreateClassFromOptions) (Class, error)

	Deprovision(string, string) error
	InstanceParentHierarchy(*apiv1.ServiceInstance) (*apiv1.ClusterServiceClass, *apiv1.ClusterServicePlan, *apiv1.ClusterServiceBroker, error)
	InstanceToServiceClassAndPlan(*apiv1.ServiceInstance) (*apiv1.ClusterServiceClass, *apiv1.ClusterServicePlan, error)
	IsInstanceFailed(*apiv1.ServiceInstance) bool
	IsInstanceReady(*apiv1.ServiceInstance) bool
	Provision(string, string, string, bool, *ProvisionOptions) (*apiv1.ServiceInstance, error)
	RetrieveInstance(string, string) (*apiv1.ServiceInstance, error)
	RetrieveInstanceByBinding(*apiv1.ServiceBinding) (*apiv1.ServiceInstance, error)
	RetrieveInstances(string, string, string) (*apiv1.ServiceInstanceList, error)
	RetrieveInstancesByPlan(Plan) ([]apiv1.ServiceInstance, error)
	TouchInstance(string, string, int) error
	WaitForInstance(string, string, time.Duration, *time.Duration) (*apiv1.ServiceInstance, error)
	WaitForInstanceToNotExist(string, string, time.Duration, *time.Duration) (*apiv1.ServiceInstance, error)

	RetrievePlans(string, ScopeOptions) ([]Plan, error)
	RetrievePlanByName(string, ScopeOptions) (Plan, error)
	RetrievePlanByClassAndName(string, string, ScopeOptions) (Plan, error)
	RetrievePlanByClassIDAndName(string, string, ScopeOptions) (Plan, error)
	RetrievePlanByID(string, ScopeOptions) (Plan, error)

	RetrieveSecretByBinding(*apiv1.ServiceBinding) (*apicorev1.Secret, error)

	ServerVersion() (*version.Info, error)
}

// SDK wrapper around the generated Go client for the Kubernetes Service Catalog
type SDK struct {
	K8sClient            kubernetes.Interface
	ServiceCatalogClient clientset.Interface
}

// ServiceCatalog is the underlying generated Service Catalog versioned interface
// It should be used instead of accessing the client directly.
func (sdk *SDK) ServiceCatalog() v1.Servicecatalogv1Interface {
	return sdk.ServiceCatalogClient.Servicecatalogv1()
}

// Core is the underlying generated Core API versioned interface
// It should be used instead of accessing the client directly.
func (sdk *SDK) Core() corev1.CoreV1Interface {
	return sdk.K8sClient.CoreV1()
}
