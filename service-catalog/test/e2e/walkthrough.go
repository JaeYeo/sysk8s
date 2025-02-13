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

package e2e

import (
	"bytes"
	"context"

	v1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/test/e2e/framework"
	"github.com/kubernetes-sigs/service-catalog/test/util"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = framework.ServiceCatalogDescribe("walkthrough", func() {
	f := framework.NewDefaultFramework("walkthrough-example")

	var (
		upsbrokername                  = "ups-broker"
		brokerName                     = upsbrokername
		serviceclassName               = "user-provided-service"
		serviceclassID                 = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
		serviceplanID                  = "86064792-7ea2-467b-af93-ac9694d96d52"
		serviceclassNameWithSinglePlan = "user-provided-service-single-plan"
		serviceclassIDWithSinglePlan   = "5f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
		testns                         = "test-ns"
		instanceName                   = "ups-instance"
		bindingName                    = "ups-binding"
		instanceNameDef                = "ups-instance-def"
		instanceNameK8sNames           = "ups-instance-k8s-names"
		instanceNameK8sNamesDef        = "ups-instance-k8s-names-def"
	)

	BeforeEach(func() {
		// Deploy the ups-broker
		By("Creating a ups-broker pod")
		pod, err := f.KubeClientSet.CoreV1().Pods(f.Namespace.Name).Create(context.Background(), NewUPSBrokerPod(upsbrokername), metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create upsbroker pod")

		By("Waiting for ups-broker pod to be running")
		err = framework.WaitForPodRunningInNamespace(f.KubeClientSet, pod)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ups-broker service")
		_, err = f.KubeClientSet.CoreV1().Services(f.Namespace.Name).Create(context.Background(), NewUPSBrokerService(upsbrokername), metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create upsbroker service")

		By("Waiting for service endpoint")
		err = framework.WaitForEndpoint(f.KubeClientSet, f.Namespace.Name, upsbrokername)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		rc, err := f.KubeClientSet.CoreV1().Pods(f.Namespace.Name).GetLogs(upsbrokername, &v1.PodLogOptions{
			Container: upsbrokername,
		}).Stream(context.Background())
		if err != nil {
			framework.Logf("Error getting logs for pod %s: %v", upsbrokername, err)
		} else {
			defer rc.Close()
			buf := new(bytes.Buffer)
			buf.ReadFrom(rc)
			framework.Logf("Pod %s has the following logs:\n%sEnd %s logs", upsbrokername, buf.String(), upsbrokername)
		}

		// Delete ups-broker pod and service
		By("Deleting the ups-broker pod")
		err = f.KubeClientSet.CoreV1().Pods(f.Namespace.Name).Delete(context.Background(), upsbrokername, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())

		By("Deleting the ups-broker service")
		err = f.KubeClientSet.CoreV1().Services(f.Namespace.Name).Delete(context.Background(), upsbrokername, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
	})

	It("Runs through the walkthrough with cluster resources", func() {
		// Broker and ClusterServiceClass should become ready
		By("Make sure the named ClusterServiceBroker does not exist before create")
		if _, err := f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Get(context.Background(), brokerName, metav1.GetOptions{}); err == nil {
			err = f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Delete(context.Background(), brokerName, metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

			By("Waiting for ClusterServiceBroker to not exist")
			err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName)
			Expect(err).NotTo(HaveOccurred())
		}

		By("Creating a ClusterServiceBroker")
		url := "http://" + upsbrokername + "." + f.Namespace.Name + ".svc.cluster.local"
		broker := &v1.ClusterServiceBroker{
			ObjectMeta: metav1.ObjectMeta{
				Name: brokerName,
			},
			Spec: v1.ClusterServiceBrokerSpec{
				CommonServiceBrokerSpec: v1.CommonServiceBrokerSpec{
					URL: url,
				},
			},
		}
		broker, err := f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Create(context.Background(), broker, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create ClusterServiceBroker")

		By("Waiting for ClusterServiceBroker to be ready")
		err = util.WaitForBrokerCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			broker.Name,
			v1.ServiceBrokerCondition{
				Type:   v1.ServiceBrokerConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for ClusterServiceBroker to be ready")

		By("Waiting for ClusterServiceClass to be ready")
		err = util.WaitForServiceClassToExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceclassID)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for ClusterServiceclass to be ready")

		By("Waiting for ClusterServicePlan to be ready")
		err = util.WaitForServicePlanToExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceplanID)
		Ω(err).ShouldNot(HaveOccurred(), "serviceplan never became ready")

		// Provisioning a ServiceInstance and binding to it
		By("Creating a namespace")
		testnamespace, err := framework.CreateKubeNamespace(testns, f.KubeClientSet)
		Expect(err).NotTo(HaveOccurred(), "failed to create kube namespace")

		By("Creating a ServiceInstance")
		instance := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceName,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ClusterServiceClassExternalName: serviceclassName,
					ClusterServicePlanExternalName:  "default",
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instance, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance")
		Expect(instance).NotTo(BeNil())

		By("Waiting for ServiceInstance to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceName,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for instance to be ready")

		// Make sure references have been resolved
		By("References should have been resolved before ServiceInstance is ready ")
		sc, err := f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Get(context.Background(), instanceName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to get ServiceInstance after binding")
		Expect(sc.Spec.ClusterServiceClassRef).NotTo(BeNil())
		Expect(sc.Spec.ClusterServicePlanRef).NotTo(BeNil())
		Expect(sc.Spec.ClusterServiceClassRef.Name).To(Equal(serviceclassID))
		Expect(sc.Spec.ClusterServicePlanRef.Name).To(Equal(serviceplanID))

		// Binding to the ServiceInstance
		By("Creating a ServiceBinding")
		binding := &v1.ServiceBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      bindingName,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceBindingSpec{
				InstanceRef: v1.LocalObjectReference{
					Name: instanceName,
				},
				SecretName: "my-secret",
			},
		}
		binding, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBindings(testnamespace.Name).Create(context.Background(), binding, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create binding")
		Expect(binding).NotTo(BeNil())

		By("Waiting for ServiceBinding to be ready")
		_, err = util.WaitForBindingCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			bindingName,
			v1.ServiceBindingCondition{
				Type:   v1.ServiceBindingConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for binding to be ready")

		By("Secret should have been created after binding")
		_, err = f.KubeClientSet.CoreV1().Secrets(testnamespace.Name).Get(context.Background(), "my-secret", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create secret after binding")

		// Unbinding from the ServiceInstance
		By("Deleting the ServiceBinding")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBindings(testnamespace.Name).Delete(context.Background(), bindingName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the binding")

		By("Waiting for ServiceBinding to not exist")
		err = util.WaitForBindingToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, bindingName)
		Expect(err).NotTo(HaveOccurred())

		By("Secret should been deleted after delete the binding")
		_, err = f.KubeClientSet.CoreV1().Secrets(testnamespace.Name).Get(context.Background(), "my-secret", metav1.GetOptions{})
		Expect(err).To(HaveOccurred())

		// Deprovisioning the ServiceInstance
		By("Deleting the ServiceInstance")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance")

		By("Waiting for ServiceInstance to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceName)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using a default plan")
		instanceDef := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameDef,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ClusterServiceClassExternalName: serviceclassNameWithSinglePlan,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceDef, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with default plan")
		Expect(instanceDef).NotTo(BeNil())

		By("Waiting for ServiceInstance to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameDef,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for instance with default plan to be ready")

		// Deprovisioning the ServiceInstance with default plan
		By("Deleting the ServiceInstance with default plan")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameDef, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with default plan")

		By("Waiting for ServiceInstance with default plan to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameDef)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using k8s names plan")
		instanceK8SNames := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameK8sNames,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ClusterServiceClassName: serviceclassID,
					ClusterServicePlanName:  serviceplanID,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceK8SNames, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with K8S names")
		Expect(instanceK8SNames).NotTo(BeNil())

		By("Waiting for ServiceInstance with k8s names to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameK8sNames,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for instance with k8s names to be ready")

		// Deprovisioning the ServiceInstance with k8s names
		By("Deleting the ServiceInstance with k8s names")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameK8sNames, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with k8s names")

		By("Waiting for ServiceInstance with k8s names to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameK8sNames)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using k8s name and default plan")
		instanceK8SNamesDef := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameK8sNamesDef,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ClusterServiceClassName: serviceclassIDWithSinglePlan,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceK8SNamesDef, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with K8S name and default plan")
		Expect(instanceK8SNamesDef).NotTo(BeNil())

		By("Waiting for ServiceInstance with k8s name and default plan to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameK8sNamesDef,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for instance with k8s name and default plan to be ready")

		// Deprovisioning the ServiceInstance with k8s name and default plan
		By("Deleting the ServiceInstance with k8s name and default plan")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameK8sNamesDef, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with k8s name and default plan")

		By("Waiting for ServiceInstance with k8s name and default plan to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameK8sNamesDef)
		Expect(err).NotTo(HaveOccurred())

		By("Deleting the test namespace")
		err = framework.DeleteKubeNamespace(f.KubeClientSet, testnamespace.Name)
		Expect(err).NotTo(HaveOccurred())

		// Deleting ClusterServiceBroker and ClusterServiceClass
		By("Deleting the ClusterServiceBroker")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Delete(context.Background(), brokerName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

		By("Waiting for ClusterServiceBroker to not exist")
		err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName)
		Expect(err).NotTo(HaveOccurred())

		By("Waiting for ClusterServiceClass to not exist")
		err = util.WaitForServiceClassToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceclassID)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Runs through the walkthrough with namespaced resources", func() {
		// Everything has to run in a namespace
		By("Creating a namespace")
		testnamespace, err := framework.CreateKubeNamespace(testns, f.KubeClientSet)
		Expect(err).NotTo(HaveOccurred(), "failed to create kube namespace")

		// Broker and ServiceClass should become ready
		By("Make sure the named ServiceBroker does not exist before create")
		if _, err := f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(testnamespace.Name).Get(context.Background(), brokerName, metav1.GetOptions{}); err == nil {
			err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(testnamespace.Name).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

			By("Waiting for ServiceBroker to not exist")
			err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName, testnamespace.Name)
			Expect(err).NotTo(HaveOccurred())
		}

		By("Creating a ServiceBroker")
		url := "http://" + upsbrokername + "." + f.Namespace.Name + ".svc.cluster.local"
		broker := &v1.ServiceBroker{
			ObjectMeta: metav1.ObjectMeta{
				Name:      brokerName,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceBrokerSpec{
				CommonServiceBrokerSpec: v1.CommonServiceBrokerSpec{
					URL: url,
				},
			},
		}
		broker, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(testnamespace.Name).Create(context.Background(), broker, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create ServiceBroker")

		By("Waiting for ServiceBroker to be ready")
		err = util.WaitForBrokerCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			broker.Name,
			v1.ServiceBrokerCondition{
				Type:   v1.ServiceBrokerConditionReady,
				Status: v1.ConditionTrue,
			},
			testnamespace.Name,
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for ServiceBroker to be ready")

		By("Waiting for ServiceClass to be ready")
		err = util.WaitForServiceClassToExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceclassID, testnamespace.Name)
		Expect(err).NotTo(HaveOccurred(), "failed to wait for ServiceClass to be ready")

		By("Waiting for ServicePlan to be ready")
		err = util.WaitForServicePlanToExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceplanID, testnamespace.Name)
		Ω(err).ShouldNot(HaveOccurred(), "serviceplan never became ready")

		// Provisioning a ServiceInstance and binding to it
		By("Creating a ServiceInstance")
		instance := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceName,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ServiceClassExternalName: serviceclassName,
					ServicePlanExternalName:  "default",
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instance, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance")
		Expect(instance).NotTo(BeNil())

		By("Waiting for ServiceInstance to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceName,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait instance to be ready")

		// Make sure references have been resolved
		By("References should have been resolved before ServiceInstance is ready ")
		sc, err := f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Get(context.Background(), instanceName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to get ServiceInstance after binding")
		Expect(sc.Spec.ServiceClassRef).NotTo(BeNil())
		Expect(sc.Spec.ServicePlanRef).NotTo(BeNil())
		Expect(sc.Spec.ServiceClassRef.Name).To(Equal(serviceclassID))
		Expect(sc.Spec.ServicePlanRef.Name).To(Equal(serviceplanID))

		// Binding to the ServiceInstance
		By("Creating a ServiceBinding")
		binding := &v1.ServiceBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      bindingName,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceBindingSpec{
				InstanceRef: v1.LocalObjectReference{
					Name: instanceName,
				},
				SecretName: "my-secret",
			},
		}
		binding, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBindings(testnamespace.Name).Create(context.Background(), binding, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create binding")
		Expect(binding).NotTo(BeNil())

		By("Waiting for ServiceBinding to be ready")
		_, err = util.WaitForBindingCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			bindingName,
			v1.ServiceBindingCondition{
				Type:   v1.ServiceBindingConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait binding to be ready")

		By("Secret should have been created after binding")
		_, err = f.KubeClientSet.CoreV1().Secrets(testnamespace.Name).Get(context.Background(), "my-secret", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create secret after binding")

		// Unbinding from the ServiceInstance
		By("Deleting the ServiceBinding")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBindings(testnamespace.Name).Delete(context.Background(), bindingName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the binding")

		By("Waiting for ServiceBinding to not exist")
		err = util.WaitForBindingToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, bindingName)
		Expect(err).NotTo(HaveOccurred())

		By("Secret should been deleted after delete the binding")
		_, err = f.KubeClientSet.CoreV1().Secrets(testnamespace.Name).Get(context.Background(), "my-secret", metav1.GetOptions{})
		Expect(err).To(HaveOccurred())

		// Deprovisioning the ServiceInstance
		By("Deleting the ServiceInstance")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance")

		By("Waiting for ServiceInstance to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceName)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using a default plan")
		instanceDef := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameDef,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ServiceClassExternalName: serviceclassNameWithSinglePlan,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceDef, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with default plan")
		Expect(instanceDef).NotTo(BeNil())

		By("Waiting for ServiceInstance to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameDef,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait instance with default plan to be ready")

		// Deprovisioning the ServiceInstance with default plan
		By("Deleting the ServiceInstance with default plan")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameDef, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with default plan")

		By("Waiting for ServiceInstance with default plan to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameDef)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using k8s names plan")
		instanceK8SNames := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameK8sNames,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ServiceClassName: serviceclassID,
					ServicePlanName:  serviceplanID,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceK8SNames, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with K8S names")
		Expect(instanceK8SNames).NotTo(BeNil())

		By("Waiting for ServiceInstance with k8s names to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameK8sNames,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait instance with k8s names to be ready")

		// Deprovisioning the ServiceInstance with k8s names
		By("Deleting the ServiceInstance with k8s names")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameK8sNames, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with k8s names")

		By("Waiting for ServiceInstance with k8s names to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameK8sNames)
		Expect(err).NotTo(HaveOccurred())

		By("Creating a ServiceInstance using k8s name and default plan")
		instanceK8SNamesDef := &v1.ServiceInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instanceNameK8sNamesDef,
				Namespace: testnamespace.Name,
			},
			Spec: v1.ServiceInstanceSpec{
				PlanReference: v1.PlanReference{
					ServiceClassName: serviceclassIDWithSinglePlan,
				},
			},
		}
		instance, err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Create(context.Background(), instanceK8SNamesDef, metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to create instance with K8S name and default plan")
		Expect(instanceK8SNamesDef).NotTo(BeNil())

		By("Waiting for ServiceInstance with k8s name and default plan to be ready")
		err = util.WaitForInstanceCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
			testnamespace.Name,
			instanceNameK8sNamesDef,
			v1.ServiceInstanceCondition{
				Type:   v1.ServiceInstanceConditionReady,
				Status: v1.ConditionTrue,
			},
		)
		Expect(err).NotTo(HaveOccurred(), "failed to wait instance with k8s name and default plan to be ready")

		// Deprovisioning the ServiceInstance with k8s name and default plan
		By("Deleting the ServiceInstance with k8s name and default plan")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceInstances(testnamespace.Name).Delete(context.Background(), instanceNameK8sNamesDef, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the instance with k8s name and default plan")

		By("Waiting for ServiceInstance with k8s name and default plan to not exist")
		err = util.WaitForInstanceToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), testnamespace.Name, instanceNameK8sNamesDef)
		Expect(err).NotTo(HaveOccurred())

		By("Deleting the test namespace")
		err = framework.DeleteKubeNamespace(f.KubeClientSet, testnamespace.Name)
		Expect(err).NotTo(HaveOccurred())

		// Deleting ServiceBroker and ServiceClass
		By("Deleting the ServiceBroker")
		err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(testnamespace.Name).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

		By("Waiting for ServiceBroker to not exist")
		err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName, testnamespace.Name)
		Expect(err).NotTo(HaveOccurred())

		By("Waiting for ServiceClass to not exist")
		err = util.WaitForServiceClassToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), serviceclassID, testnamespace.Name)
		Expect(err).NotTo(HaveOccurred())
	})
})
