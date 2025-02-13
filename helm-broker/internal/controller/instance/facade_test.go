package instance_test

import (
	"testing"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/kyma-project/helm-broker/internal/controller/broker"
	"github.com/kyma-project/helm-broker/internal/controller/instance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	namespace         = "testing-ns"
	clusterBrokerName = "helmbroker"
)

func TestFacadeExistingHBInstance(t *testing.T) {
	// given
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme, fixHBServiceClass(), fixServiceClass(), fixHBServiceInstnace())
	facade := instance.New(cli, clusterBrokerName)

	// when / then
	exists, err := facade.AnyServiceInstanceExistsForNamespacedServiceBroker(namespace)

	require.NoError(t, err)
	assert.True(t, exists)
}

func TestFacadeNoInstances(t *testing.T) {
	// given
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme, fixHBServiceClass(), fixServiceClass())
	facade := instance.New(cli, clusterBrokerName)

	// when / then
	exists, err := facade.AnyServiceInstanceExistsForNamespacedServiceBroker(namespace)

	require.NoError(t, err)
	assert.False(t, exists)
}

func TestFacadeNoHBInstances(t *testing.T) {
	// given
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme, fixHBServiceClass(), fixServiceClass(), fixServiceInstnaceClusterServiceClass(), fixServiceInstance())
	facade := instance.New(cli, clusterBrokerName)

	// when / then
	exists, err := facade.AnyServiceInstanceExistsForNamespacedServiceBroker(namespace)

	require.NoError(t, err)
	assert.False(t, exists)
}

func fixHBServiceInstnace() *v1.ServiceInstance {
	return &v1.ServiceInstance{
		ObjectMeta: v1.ObjectMeta{
			Name:      "si-application01",
			Namespace: namespace,
		},
		Spec: v1.ServiceInstanceSpec{
			ServiceClassRef: &v1.LocalObjectReference{
				Name: "a-class",
			},
		}}
}

func fixServiceInstance() *v1.ServiceInstance {
	return &v1.ServiceInstance{
		ObjectMeta: v1.ObjectMeta{
			Name:      "si-application01",
			Namespace: namespace,
		},
		Spec: v1.ServiceInstanceSpec{
			ServiceClassRef: &v1.LocalObjectReference{
				Name: "sc-service",
			},
		}}
}

func fixServiceInstnaceClusterServiceClass() *v1.ServiceInstance {
	return &v1.ServiceInstance{
		ObjectMeta: v1.ObjectMeta{
			Name:      "si-clusterserviceclass",
			Namespace: namespace,
		},
		Spec: v1.ServiceInstanceSpec{
			ClusterServiceClassRef: &v1.ClusterObjectReference{
				Name: "some-class",
			},
		}}
}

func fixHBServiceClass() *v1.ServiceClass {
	return &v1.ServiceClass{
		ObjectMeta: v1.ObjectMeta{
			Name:      "a-class",
			Namespace: namespace,
		},
		Spec: v1.ServiceClassSpec{
			ServiceBrokerName: broker.NamespacedBrokerName,
		}}
}

func fixServiceClass() *v1.ServiceClass {
	return &v1.ServiceClass{
		ObjectMeta: v1.ObjectMeta{
			Name:      "sc-service",
			Namespace: namespace,
		},
		Spec: v1.ServiceClassSpec{
			ServiceBrokerName: "other-broker",
		}}
}
