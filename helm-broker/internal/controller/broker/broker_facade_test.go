package broker

import (
	"fmt"
	"testing"

	"context"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8s_errors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	k8s_testing "k8s.io/client-go/testing"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestServiceBrokerCreateHappyPath(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	svcURL := fmt.Sprintf("http://%s.%s.svc.cluster.local/ns/%s", fixService(), fixWorkingNs(), "stage")
	sut := NewBrokersFacade(cli, fixWorkingNs(), fixService(), logrus.New())
	// WHEN
	sut.SetNamespace(fixDestNs())
	err := sut.Create()

	// THEN
	require.NoError(t, err)

	actualBroker := &v1.ServiceBroker{}
	err = cli.Get(context.Background(), types.NamespacedName{Name: fixBrokerName(), Namespace: fixDestNs()}, actualBroker)
	require.NoError(t, err)
	assert.Equal(t, "true", actualBroker.Labels["namespaced-helm-broker"])
	assert.Equal(t, svcURL, actualBroker.Spec.URL)

	require.NoError(t, err)
}

func TestServiceBrokerDeleteHappyPath(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewBrokersFacade(cli, fixWorkingNs(), fixService(), logrus.New())
	// WHEN
	sut.SetNamespace(fixDestNs())
	err := sut.Delete()
	// THEN
	require.NoError(t, err)
}

func TestServiceBrokerDeleteNotFoundErrorsIgnored(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewBrokersFacade(cli, fixWorkingNs(), fixService(), logrus.New())
	// WHEN
	sut.SetNamespace(fixDestNs())
	err := sut.Delete()
	// THEN
	require.NoError(t, err)
}

func TestServiceBrokerDoesNotExist(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme)

	sut := NewBrokersFacade(cli, fixWorkingNs(), fixService(), logrus.New())
	// WHEN
	sut.SetNamespace(fixDestNs())
	ex, err := sut.Exist()
	// THEN
	require.NoError(t, err)
	assert.False(t, ex)
}

func TestServiceBrokerExist(t *testing.T) {
	// GIVEN
	require.NoError(t, v1.AddToScheme(scheme.Scheme))
	cli := fake.NewFakeClientWithScheme(scheme.Scheme, &v1.ServiceBroker{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      fixBrokerName(),
			Namespace: fixDestNs(),
		}})

	sut := NewBrokersFacade(cli, fixWorkingNs(), fixService(), logrus.New())
	// WHEN
	sut.SetNamespace(fixDestNs())
	ex, err := sut.Exist()
	// THEN
	require.NoError(t, err)
	assert.True(t, ex)
}

func fixDestNs() string {
	return "stage"
}

func fixService() string {
	return "service"
}

func fixWorkingNs() string {
	return "kyma-system"
}

func failingReactor(retErr error) k8s_testing.ReactionFunc {
	return func(action k8s_testing.Action) (handled bool, ret runtime.Object, err error) {
		return true, nil, retErr
	}
}

func fixAlreadyExistError() error {
	return k8s_errors.NewAlreadyExists(schema.GroupResource{}, "")
}

func assertPerformedAction(t *testing.T, verb, resource string, actions []k8s_testing.Action) {
	for _, action := range actions {
		if action.Matches(verb, resource) {
			return
		}
	}
	t.Errorf("Action %s %s not found", verb, resource)
}
