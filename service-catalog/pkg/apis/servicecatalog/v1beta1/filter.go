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

package v1

import (
	"strconv"

	"github.com/kubernetes-sigs/service-catalog/pkg/filter"
	"k8s.io/apimachinery/pkg/labels"
)

// These are functions to support filtering. This is where we can add more fields
// to the labels.Set to support other kinds of catalog filtering.

// ConvertServiceClassToProperties takes a Service Class and pulls out the
// properties we support for filtering, converting them into a map in the
// expected format.
func ConvertServiceClassToProperties(serviceClass *ServiceClass) filter.Properties {
	if serviceClass == nil {
		return labels.Set{}
	}
	return labels.Set{
		FilterName:             serviceClass.Name,
		FilterSpecExternalName: serviceClass.Spec.ExternalName,
		FilterSpecExternalID:   serviceClass.Spec.ExternalID,
	}
}

// IsValidServiceClassProperty returns true if the specified property
// is a valid filterable property of ServiceClasses
func IsValidServiceClassProperty(p string) bool {
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID
}

// ConvertServicePlanToProperties takes a Service Plan and pulls out the
// properties we support for filtering, converting them into a map in the
// expected format.
func ConvertServicePlanToProperties(servicePlan *ServicePlan) filter.Properties {
	if servicePlan == nil {
		return labels.Set{}
	}
	return labels.Set{
		FilterName:                 servicePlan.Name,
		FilterSpecExternalName:     servicePlan.Spec.ExternalName,
		FilterSpecExternalID:       servicePlan.Spec.ExternalID,
		FilterSpecServiceClassName: servicePlan.Spec.ServiceClassRef.Name,
		FilterSpecFree:             strconv.FormatBool(servicePlan.Spec.Free),
	}
}

// IsValidServicePlanProperty returns true if the specified property
// is a valid filterable property of ServicePlans
func IsValidServicePlanProperty(p string) bool {
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID || p == FilterSpecServiceClassName || p == FilterSpecFree
}

// ConvertClusterServiceClassToProperties takes a Service Class and pulls out the
// properties we support for filtering, converting them into a map in the
// expected format.
func ConvertClusterServiceClassToProperties(serviceClass *ClusterServiceClass) filter.Properties {
	if serviceClass == nil {
		return labels.Set{}
	}
	return labels.Set{
		FilterName:             serviceClass.Name,
		FilterSpecExternalName: serviceClass.Spec.ExternalName,
		FilterSpecExternalID:   serviceClass.Spec.ExternalID,
	}
}

// IsValidClusterServiceClassProperty returns true if the specified property
// is a valid filterable property of ClusterServiceClasses
func IsValidClusterServiceClassProperty(p string) bool {
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID
}

// ConvertClusterServicePlanToProperties takes a Service Plan and pulls out the
// properties we support for filtering, converting them into a map in the
// expected format.
func ConvertClusterServicePlanToProperties(servicePlan *ClusterServicePlan) filter.Properties {
	if servicePlan == nil {
		return labels.Set{}
	}
	return labels.Set{
		FilterName:                        servicePlan.Name,
		FilterSpecExternalName:            servicePlan.Spec.ExternalName,
		FilterSpecExternalID:              servicePlan.Spec.ExternalID,
		FilterSpecClusterServiceClassName: servicePlan.Spec.ClusterServiceClassRef.Name,
		FilterSpecFree:                    strconv.FormatBool(servicePlan.Spec.Free),
	}
}

// IsValidClusterServicePlanProperty returns true if the specified property
// is a valid filterable property of ServicePlans
func IsValidClusterServicePlanProperty(p string) bool {
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID || p == FilterSpecClusterServiceClassName || p == FilterSpecFree
}
