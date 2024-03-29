/*
Copyright 2019 The Kubernetes Authors.

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

package clusterbroker

import (
	"context"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	scClientset "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
)

type creator struct {
	common
	sc        scClientset.Servicecatalogv1Interface
	namespace string
}

func newCreator(cli ClientGetter, ns string) *creator {
	return &creator{
		sc:        cli.ServiceCatalogClient().Servicecatalogv1(),
		namespace: ns,
		common: common{
			sc:        cli.ServiceCatalogClient().Servicecatalogv1(),
			namespace: ns,
		},
	}
}

func (c *creator) execute() error {
	klog.Info("Start prepare resources for ClusterServiceBroker test")
	for _, fn := range []func() error{
		c.registerClusterServiceBroker,
		c.checkClusterServiceClass,
		c.checkClusterServicePlan,
		c.createServiceInstance,
		c.createServiceBinding,
	} {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *creator) registerClusterServiceBroker() error {
	klog.Infof("Create ClusterServiceBroker %q", clusterServiceBrokerName)
	_, err := c.sc.ClusterServiceBrokers().Create(context.Background(),&v1.ClusterServiceBroker{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterServiceBrokerName,
			Namespace: c.namespace,
		},
		Spec: v1.ClusterServiceBrokerSpec{
			CommonServiceBrokerSpec: v1.CommonServiceBrokerSpec{
				URL: "http://test-broker-test-broker.test-broker.svc.cluster.local",
			},
		},
	},metav1.CreateOptions{})

	if err != nil {
		return errors.Wrap(err, "failed during creating ClusterServiceBroker")
	}

	return nil
}

func (c *creator) createServiceInstance() error {
	klog.Info("Create ServiceInstance")
	if err := c.createDefaultServiceInstance(); err != nil {
		return errors.Wrap(err, "failed during creating ServiceInstance")
	}

	klog.Info("Check ServiceInstance is ready")
	if err := c.assertServiceInstanceIsReady(); err != nil {
		return errors.Wrap(err, "failed during checking ServiceInstance conditions")
	}

	return nil
}

func (c *creator) createDefaultServiceInstance() error {
	_, err := c.sc.ServiceInstances(c.namespace).Create(context.Background(), &v1.ServiceInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceInstanceName,
			Namespace: c.namespace,
		},
		Spec: v1.ServiceInstanceSpec{
			PlanReference: v1.PlanReference{
				ClusterServiceClassExternalName: "test-service-multiple-plans",
				ClusterServicePlanExternalName:  "default",
			},
			Parameters: &runtime.RawExtension{
				Raw: []byte(`{ "param-1":"value-1", "param-2":"value-2" }`),
			},
		},
	}, metav1.CreateOptions{})

	return err
}

func (c *creator) createServiceBinding() error {
	klog.Info("Create ServiceBinding")
	if err := c.createDefaultServiceBinding(); err != nil {
		return errors.Wrap(err, "failed during creating ServiceBinding")
	}

	klog.Info("Check ServiceBinding is ready")
	if err := c.assertServiceBindingIsReady(); err != nil {
		return errors.Wrap(err, "failed during checking ServiceBinding conditions")
	}

	return nil
}

func (c *creator) createDefaultServiceBinding() error {
	_, err := c.sc.ServiceBindings(c.namespace).Create(context.Background(), &v1.ServiceBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceBindingName,
			Namespace: c.namespace,
		},
		Spec: v1.ServiceBindingSpec{
			InstanceRef: v1.LocalObjectReference{
				Name: serviceInstanceName,
			},
		},
	}, metav1.CreateOptions{})

	return err
}
