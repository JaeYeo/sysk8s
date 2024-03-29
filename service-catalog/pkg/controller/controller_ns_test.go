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

package controller

import (
	"testing"
	"time"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	testServiceClassGUID           = "scguid"
	testServicePlanGUID            = "spguid"
	testServiceBrokerName          = "test-servicebroker"
	testServiceClassName           = "test-serviceclass"
	testServicePlanName            = "test-serviceplan"
	testNonbindableServicePlanName = "test-unbindable-serviceplan"
	testNonbindableServicePlanGUID = "unbindable-serviceplan"
)

func getTestCommonServiceBrokerSpec() v1.CommonServiceBrokerSpec {
	return v1.CommonServiceBrokerSpec{
		URL:            "https://example.com",
		RelistBehavior: v1.ServiceBrokerRelistBehaviorDuration,
		RelistDuration: &metav1.Duration{Duration: 15 * time.Minute},
	}
}

func getTestServiceBroker() *v1.ServiceBroker {
	return &v1.ServiceBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testServiceBrokerName,
			Namespace: testNamespace,
		},
		Spec: v1.ServiceBrokerSpec{
			CommonServiceBrokerSpec: getTestCommonServiceBrokerSpec(),
		},
	}
}

func getTestServiceBrokerWithAuth(authInfo *v1.ServiceBrokerAuthInfo) *v1.ServiceBroker {
	broker := getTestServiceBroker()
	broker.Spec.AuthInfo = authInfo
	return broker
}

func getTestBrokerBasicAuthInfo() *v1.ServiceBrokerAuthInfo {
	return &v1.ServiceBrokerAuthInfo{
		Basic: &v1.BasicAuthConfig{
			SecretRef: &v1.LocalObjectReference{Name: "auth-secret"},
		},
	}
}

func getTestBrokerBearerAuthInfo() *v1.ServiceBrokerAuthInfo {
	return &v1.ServiceBrokerAuthInfo{
		Bearer: &v1.BearerTokenAuthConfig{
			SecretRef: &v1.LocalObjectReference{Name: "auth-secret"},
		},
	}
}

func assertServiceBrokerReadyFalse(t *testing.T, obj runtime.Object) {
	assertServiceBrokerCondition(t, obj, v1.ServiceBrokerConditionReady, v1.ConditionFalse)
}

func assertServiceBrokerReadyTrue(t *testing.T, obj runtime.Object) {
	assertServiceBrokerCondition(t, obj, v1.ServiceBrokerConditionReady, v1.ConditionTrue)
}

func assertServiceBrokerCondition(t *testing.T, obj runtime.Object, conditionType v1.ServiceBrokerConditionType, status v1.ConditionStatus) {
	broker, ok := obj.(*v1.ServiceBroker)
	if !ok {
		fatalf(t, "Couldn't convert object %+v into a *v1.ServiceBroker", obj)
	}

	for _, condition := range broker.Status.Conditions {
		if condition.Type == conditionType && condition.Status != status {
			fatalf(t, "%v condition had unexpected status; expected %v, got %v", conditionType, status, condition.Status)
		}
	}
}

func getTestCommonServiceClassSpec() v1.CommonServiceClassSpec {
	return v1.CommonServiceClassSpec{
		Description:  "a test service",
		ExternalName: testServiceClassName,
		ExternalID:   testServiceClassGUID,
		Bindable:     true,
	}
}

func getTestServiceClass() *v1.ServiceClass {
	return &v1.ServiceClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testServiceClassGUID,
			Namespace: testNamespace,
			Labels: map[string]string{
				v1.GroupName + "/" + v1.FilterSpecServiceClassRefName: util.GenerateSHA(testServiceClassGUID),
				v1.GroupName + "/" + v1.FilterSpecExternalName:        util.GenerateSHA(testServiceClassName),
				v1.GroupName + "/" + v1.FilterSpecServiceBrokerName:   util.GenerateSHA(testServiceBrokerName),
			},
		},
		Spec: v1.ServiceClassSpec{
			ServiceBrokerName:      testServiceBrokerName,
			CommonServiceClassSpec: getTestCommonServiceClassSpec(),
		},
	}
}

func getTestCommonServicePlanSpec() v1.CommonServicePlanSpec {
	return v1.CommonServicePlanSpec{
		ExternalID:   testServicePlanGUID,
		ExternalName: testServicePlanName,
		Bindable:     truePtr(),
	}
}

func getTestServicePlan() *v1.ServicePlan {
	return &v1.ServicePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testServicePlanGUID,
			Namespace: testNamespace,
			Labels: map[string]string{
				v1.GroupName + "/" + v1.FilterSpecServicePlanRefName:  util.GenerateSHA(testServicePlanGUID),
				v1.GroupName + "/" + v1.FilterSpecExternalName:        util.GenerateSHA(testServicePlanName),
				v1.GroupName + "/" + v1.FilterSpecServiceBrokerName:   util.GenerateSHA(testServiceBrokerName),
				v1.GroupName + "/" + v1.FilterSpecServiceClassRefName: util.GenerateSHA(testServiceClassGUID),
			},
		},
		Spec: v1.ServicePlanSpec{
			ServiceBrokerName:     testServiceBrokerName,
			CommonServicePlanSpec: getTestCommonServicePlanSpec(),
			ServiceClassRef: v1.LocalObjectReference{
				Name: testServiceClassGUID,
			},
		},
		Status: v1.ServicePlanStatus{},
	}
}

func getTestServiceInstanceWithNamespacedPlanReference() *v1.ServiceInstance {
	return &v1.ServiceInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:       testServiceInstanceName,
			Namespace:  testNamespace,
			Generation: 1,
		},
		Spec: v1.ServiceInstanceSpec{
			PlanReference: v1.PlanReference{
				ServiceClassExternalName: testServiceClassName,
				ServicePlanExternalName:  testServicePlanName,
			},
			ExternalID: testServiceInstanceGUID,
		},
		Status: v1.ServiceInstanceStatus{
			DeprovisionStatus: v1.ServiceInstanceDeprovisionStatusRequired,
		},
	}
}

func getTestServiceInstanceAsyncProvisioningWithNamespacedRefs(operation string) *v1.ServiceInstance {
	instance := getTestServiceInstanceWithNamespacedRefs()

	operationStartTime := metav1.NewTime(time.Now().Add(-1 * time.Hour))
	instance.Status = v1.ServiceInstanceStatus{
		Conditions: []v1.ServiceInstanceCondition{{
			Type:               v1.ServiceInstanceConditionReady,
			Status:             v1.ConditionFalse,
			Message:            "Provisioning",
			LastTransitionTime: metav1.NewTime(time.Now().Add(-5 * time.Minute)),
		}},
		AsyncOpInProgress:  true,
		OperationStartTime: &operationStartTime,
		CurrentOperation:   v1.ServiceInstanceOperationProvision,
		InProgressProperties: &v1.ServiceInstancePropertiesState{
			ServicePlanExternalName: testServicePlanName,
			ServicePlanExternalID:   testServicePlanGUID,
		},
		ObservedGeneration: instance.Generation,
		DeprovisionStatus:  v1.ServiceInstanceDeprovisionStatusRequired,
	}
	if operation != "" {
		instance.Status.LastOperation = &operation
	}

	return instance
}

func getTestServiceInstanceAsyncDeprovisioningWithNamespacedRefs(operation string) *v1.ServiceInstance {
	instance := getTestServiceInstanceWithNamespacedRefs()
	instance.Generation = 2

	operationStartTime := metav1.NewTime(time.Now().Add(-1 * time.Hour))
	instance.Status = v1.ServiceInstanceStatus{
		Conditions: []v1.ServiceInstanceCondition{{
			Type:               v1.ServiceInstanceConditionReady,
			Status:             v1.ConditionFalse,
			Message:            "Deprovisioning",
			LastTransitionTime: metav1.NewTime(time.Now().Add(-5 * time.Minute)),
		}},
		AsyncOpInProgress:  true,
		OperationStartTime: &operationStartTime,
		CurrentOperation:   v1.ServiceInstanceOperationDeprovision,
		InProgressProperties: &v1.ServiceInstancePropertiesState{
			ServicePlanExternalName: testServicePlanName,
			ServicePlanExternalID:   testServicePlanGUID,
		},

		ReconciledGeneration: 1,
		ObservedGeneration:   2,
		ExternalProperties: &v1.ServiceInstancePropertiesState{
			ServicePlanExternalName: testServicePlanName,
			ServicePlanExternalID:   testServicePlanGUID,
		},
		ProvisionStatus:   v1.ServiceInstanceProvisionStatusProvisioned,
		DeprovisionStatus: v1.ServiceInstanceDeprovisionStatusRequired,
	}
	if operation != "" {
		instance.Status.LastOperation = &operation
	}

	// Set the deleted timestamp to simulate deletion
	ts := metav1.NewTime(time.Now().Add(-5 * time.Minute))
	instance.DeletionTimestamp = &ts
	return instance
}

func getTestServiceInstanceWithNamespacedRefsAndStatus(status v1.ConditionStatus) *v1.ServiceInstance {
	instance := getTestServiceInstanceWithNamespacedRefsAndExternalProperties()
	instance.Status.Conditions = []v1.ServiceInstanceCondition{{
		Type:               v1.ServiceInstanceConditionReady,
		Status:             status,
		LastTransitionTime: metav1.NewTime(time.Now().Add(-5 * time.Minute)),
	}}
	return instance
}

func getTestServiceInstanceWithNamespacedRefsAndExternalProperties() *v1.ServiceInstance {
	sc := getTestServiceInstanceWithNamespacedRefs()
	sc.Status.ExternalProperties = &v1.ServiceInstancePropertiesState{
		ServicePlanExternalID:   testServicePlanGUID,
		ServicePlanExternalName: testServicePlanName,
	}
	return sc
}
