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

package pretty

import (
	"fmt"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
)

// Name prints in the form `<Kind> (K8S: <K8S-Name> ExternalName: <External-Name>)`
// kind is required. k8sName and externalName are optional
func Name(kind Kind, k8sName, externalName string) string {
	s := fmt.Sprintf("%s", kind)
	if k8sName != "" && externalName != "" {
		s += fmt.Sprintf(" (K8S: %q ExternalName: %q)", k8sName, externalName)
	} else if k8sName != "" {
		s += fmt.Sprintf(" (K8S: %q)", k8sName)
	} else if externalName != "" {
		s += fmt.Sprintf(" (ExternalName: %q)", externalName)
	}
	return s
}

// ServiceInstanceName returns a string with the type, namespace and name of an instance.
func ServiceInstanceName(instance *v1.ServiceInstance) string {
	return fmt.Sprintf(`%s "%s/%s"`, ServiceInstance, instance.Namespace, instance.Name)
}

// ServiceBindingName returns a string with the type, namespace and name of a binding.
func ServiceBindingName(binding *v1.ServiceBinding) string {
	return fmt.Sprintf(`%s "%s/%s"`, ServiceBinding, binding.Namespace, binding.Name)
}

// ClusterServiceBrokerName returns a string with the type and name of a broker
func ClusterServiceBrokerName(clusterServiceBrokerName string) string {
	return fmt.Sprintf(`%s %q`, ClusterServiceBroker, clusterServiceBrokerName)
}

// ServiceBrokerName returns a string with the type and name of a broker
func ServiceBrokerName(serviceBrokerName string) string {
	return fmt.Sprintf(`%s %q`, ServiceBroker, serviceBrokerName)
}

// ClusterServiceClassName returns a string with the k8s name and external name if available.
func ClusterServiceClassName(serviceClass *v1.ClusterServiceClass) string {
	if serviceClass != nil {
		return Name(ClusterServiceClass, serviceClass.Name, serviceClass.Spec.ExternalName)
	}
	return Name(ClusterServiceClass, "", "")
}

// ServiceClassName returns a string with the k8s name and external name if available.
func ServiceClassName(serviceClass *v1.ServiceClass) string {
	if serviceClass != nil {
		return Name(ServiceClass, fmt.Sprintf("%s/%s", serviceClass.Namespace, serviceClass.Name), serviceClass.Spec.ExternalName)
	}
	return Name(ServiceClass, "", "")
}

// ClusterServicePlanName returns a string with the k8s name and external name if available.
func ClusterServicePlanName(servicePlan *v1.ClusterServicePlan) string {
	if servicePlan != nil {
		return Name(ClusterServicePlan, servicePlan.Name, servicePlan.Spec.ExternalName)
	}
	return Name(ClusterServicePlan, "", "")
}

// ServicePlanName returns a string with the k8s name and external name if available.
func ServicePlanName(servicePlan *v1.ServicePlan) string {
	if servicePlan != nil {
		return Name(ServicePlan, fmt.Sprintf("%s/%s", servicePlan.Namespace, servicePlan.Name), servicePlan.Spec.ExternalName)
	}
	return Name(ServicePlan, "", "")
}

// FromServiceInstanceOfClusterServiceClassAtBrokerName returns a string in the form of "%s of %s at %s" to help in logging the full context.
func FromServiceInstanceOfClusterServiceClassAtBrokerName(instance *v1.ServiceInstance, serviceClass *v1.ClusterServiceClass, brokerName string) string {
	return fmt.Sprintf(
		"%s of %s at %s",
		ServiceInstanceName(instance), ClusterServiceClassName(serviceClass), ClusterServiceBrokerName(brokerName),
	)
}

// FromServiceInstanceOfServiceClassAtBrokerName returns a string in the form of "%s of %s at %s" to help in logging the full context.
func FromServiceInstanceOfServiceClassAtBrokerName(instance *v1.ServiceInstance, serviceClass *v1.ServiceClass, brokerName string) string {
	return fmt.Sprintf(
		"%s of %s at %s",
		ServiceInstanceName(instance), ServiceClassName(serviceClass), ServiceBrokerName(brokerName),
	)
}
