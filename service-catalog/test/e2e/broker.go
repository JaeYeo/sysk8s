/*
Copyright 2016 The Kubernetes Authors.

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
	"context"
	v1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kubernetes-sigs/service-catalog/test/e2e/framework"
	"github.com/kubernetes-sigs/service-catalog/test/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTestBroker(name, url string) *v1.ClusterServiceBroker {
	return &v1.ClusterServiceBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.ClusterServiceBrokerSpec{
			CommonServiceBrokerSpec: v1.CommonServiceBrokerSpec{
				URL: url,
			},
		},
	}
}

func newNamespacedTestBroker(name, namespace, url string) *v1.ServiceBroker {
	return &v1.ServiceBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.ServiceBrokerSpec{
			CommonServiceBrokerSpec: v1.CommonServiceBrokerSpec{
				URL: url,
			},
		},
	}
}

var _ = framework.ServiceCatalogDescribe("Brokers", func() {
	f := framework.NewDefaultFramework("create-service-broker")

	brokerName := "test-broker"

	BeforeEach(func() {
		By("Creating a user broker pod")
		pod, err := f.KubeClientSet.CoreV1().Pods(f.Namespace.Name).Create(context.Background(), NewUPSBrokerPod(brokerName), metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred())
		By("Waiting for pod to be running")
		err = framework.WaitForPodRunningInNamespace(f.KubeClientSet, pod)
		Expect(err).NotTo(HaveOccurred())
		By("Creating a user broker service")
		_, err = f.KubeClientSet.CoreV1().Services(f.Namespace.Name).Create(context.Background(), NewUPSBrokerService(brokerName), metav1.CreateOptions{})
		Expect(err).NotTo(HaveOccurred())
		By("Waiting for service endpoint")
		err = framework.WaitForEndpoint(f.KubeClientSet, f.Namespace.Name, brokerName)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		By("Deleting the user broker pod")
		err := f.KubeClientSet.CoreV1().Pods(f.Namespace.Name).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
		By("Deleting the user broker service")
		err = f.KubeClientSet.CoreV1().Services(f.Namespace.Name).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("ClusterServiceBroker", func() {
		It("should become ready", func() {
			By("Making sure the ClusterServiceBroker does not exist before creating it")
			if _, err := f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Get(context.Background(), brokerName, metav1.GetOptions{}); err == nil {
				By("deleting the ClusterServiceBroker if it does exist")
				err = f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Delete(context.Background(), brokerName, metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

				By("Waiting for the ClusterServiceBroker to not exist after deleting it")
				err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName)
				Expect(err).NotTo(HaveOccurred())
			}

			By("Creating a ClusterBroker")
			url := "http://" + brokerName + "." + f.Namespace.Name + ".svc.cluster.local"

			broker, err := f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Create(context.Background(), newTestBroker(brokerName, url), metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			By("Waiting for ClusterServiceBroker to be ready")
			err = util.WaitForBrokerCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
				broker.Name,
				v1.ServiceBrokerCondition{
					Type:   v1.ServiceBrokerConditionReady,
					Status: v1.ConditionTrue,
				})
			Expect(err).NotTo(HaveOccurred())

			By("Deleting the ClusterServiceBroker")
			err = f.ServiceCatalogClientSet.Servicecatalogv1().ClusterServiceBrokers().Delete(context.Background(), brokerName, metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred())

			By("Waiting for ClusterServiceBroker to not exist")
			err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Namespaced ServiceBroker", func() {
		It("should become ready", func() {
			By("Making sure the ServiceBroker does not exist before creating it")
			if _, err := f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(f.Namespace.Name).Get(context.Background(), brokerName, metav1.GetOptions{}); err == nil {
				By("deleting the ServiceBroker if it does exist")
				err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(f.Namespace.Name).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
				Expect(err).NotTo(HaveOccurred(), "failed to delete the broker")

				By("Waiting for the ServiceBroker to not exist after deleting it")
				err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName)
				Expect(err).NotTo(HaveOccurred())
			}

			By("Creating a ServiceBroker")
			url := "http://" + brokerName + "." + f.Namespace.Name + ".svc.cluster.local"
			broker, err := f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(f.Namespace.Name).Create(context.Background(), newNamespacedTestBroker(brokerName, f.Namespace.Name, url), metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			By("Waiting for Broker to be ready")
			err = util.WaitForBrokerCondition(f.ServiceCatalogClientSet.Servicecatalogv1(),
				broker.Name,
				v1.ServiceBrokerCondition{
					Type:   v1.ServiceBrokerConditionReady,
					Status: v1.ConditionTrue,
				},
				broker.Namespace,
			)
			Expect(err).NotTo(HaveOccurred())

			By("Deleting the ServiceBroker")
			err = f.ServiceCatalogClientSet.Servicecatalogv1().ServiceBrokers(broker.Namespace).Delete(context.Background(), brokerName, metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred())

			By("Waiting for ServiceBroker to not exist")
			err = util.WaitForBrokerToNotExist(f.ServiceCatalogClientSet.Servicecatalogv1(), brokerName, broker.Namespace)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
