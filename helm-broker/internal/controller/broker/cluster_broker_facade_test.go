package broker

import (
	"fmt"
	"testing"

	"context"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestClusterServiceBrokerCreateHappyPath(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	svcURL := fmt.Sprintf("http://%s.%s.svc.cluster.local/cluster", fixService(), fixWorkingNs())
	sut := NewClusterBrokersFacade(cli, fixWorkingNs(), fixService(), fixBrokerName(), logrus.New())
	// WHEN
	err := sut.Create()

	// THEN
	require.NoError(t, err)

	sb := &v1.ClusterServiceBroker{}
	err = cli.Get(context.Background(), types.NamespacedName{Name: fixBrokerName()}, sb)
	require.NoError(t, err)
	assert.Equal(t, svcURL, sb.Spec.URL)

	require.NoError(t, err)
}

func TestClusterServiceBrokerDeleteHappyPath(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewClusterBrokersFacade(cli, fixWorkingNs(), fixService(), fixBrokerName(), logrus.New())
	// WHEN
	err := sut.Delete()
	// THEN
	require.NoError(t, err)
}

func TestClusterServiceBrokerDeleteNotFoundErrorsIgnored(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewClusterBrokersFacade(cli, fixWorkingNs(), fixService(), fixBrokerName(), logrus.New())
	// WHEN
	err := sut.Delete()
	// THEN
	require.NoError(t, err)
}

func TestClusterServiceBrokerDoesNotExist(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewClusterBrokersFacade(cli, fixWorkingNs(), fixService(), fixBrokerName(), logrus.New())
	// WHEN
	ex, err := sut.Exist()
	// THEN
	require.NoError(t, err)
	assert.False(t, ex)
}

func TestClusterServiceBrokerExist(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme, &v1.ClusterServiceBroker{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: fixBrokerName(),
		}})

	sut := NewClusterBrokersFacade(cli, fixWorkingNs(), fixService(), fixBrokerName(), logrus.New())
	// WHEN
	ex, err := sut.Exist()
	// THEN
	require.NoError(t, err)
	assert.True(t, ex)
}

func fixBrokerName() string {
	return "helm-broker"
}
