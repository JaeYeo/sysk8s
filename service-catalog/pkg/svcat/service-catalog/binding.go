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
	"context"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	"github.com/pkg/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
)

// RetrieveBindings lists all bindings in a namespace.
func (sdk *SDK) RetrieveBindings(ns string) (*v1.ServiceBindingList, error) {
	bindings, err := sdk.ServiceCatalog().ServiceBindings(ns).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to list bindings in %s", ns)
	}

	return bindings, nil
}

// RetrieveBinding gets a binding by its name.
func (sdk *SDK) RetrieveBinding(ns, name string) (*v1.ServiceBinding, error) {
	binding, err := sdk.ServiceCatalog().ServiceBindings(ns).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get binding '%s.%s'", ns, name)
	}
	return binding, nil
}

// RetrieveBindingsByInstance gets all child bindings for an instance.
func (sdk *SDK) RetrieveBindingsByInstance(instance *v1.ServiceInstance,
) ([]v1.ServiceBinding, error) {
	// Not using a filtered list operation because it's not supported yet.
	results, err := sdk.ServiceCatalog().ServiceBindings(instance.Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "unable to search bindings")
	}

	var bindings []v1.ServiceBinding
	for _, binding := range results.Items {
		if binding.Spec.InstanceRef.Name == instance.Name {
			bindings = append(bindings, binding)
		}
	}

	return bindings, nil
}

// Bind an instance to a secret.
func (sdk *SDK) Bind(namespace, bindingName, externalID, instanceName, secretName string,
	params interface{}, secrets map[string]string) (*v1.ServiceBinding, error) {

	// Manually defaulting the name of the binding
	// I'm not doing the same for the secret since the API handles defaulting that value.
	if bindingName == "" {
		bindingName = instanceName
	}

	request := &v1.ServiceBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:      bindingName,
			Namespace: namespace,
		},
		Spec: v1.ServiceBindingSpec{
			ExternalID: externalID,
			InstanceRef: v1.LocalObjectReference{
				Name: instanceName,
			},
			SecretName:     secretName,
			Parameters:     BuildParameters(params),
			ParametersFrom: BuildParametersFrom(secrets),
		},
	}

	result, err := sdk.ServiceCatalog().ServiceBindings(namespace).Create(context.Background(), request, v1.CreateOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "bind request failed")
	}

	return result, nil
}

// Unbind deletes all bindings associated to an instance.
func (sdk *SDK) Unbind(ns, instanceName string) ([]types.NamespacedName, error) {
	instance, err := sdk.RetrieveInstance(ns, instanceName)
	if err != nil {
		return nil, err
	}
	bindings, err := sdk.RetrieveBindingsByInstance(instance)
	if err != nil {
		return nil, err
	}

	namespacedNames := []types.NamespacedName{}
	for _, b := range bindings {
		namespacedNames = append(namespacedNames, types.NamespacedName{Namespace: b.Namespace, Name: b.Name})
	}
	return sdk.DeleteBindings(namespacedNames)
}

// DeleteBindings deletes bindings by name.
func (sdk *SDK) DeleteBindings(bindings []types.NamespacedName) ([]types.NamespacedName, error) {
	var g sync.WaitGroup
	errs := make(chan error, len(bindings))
	deletedBindings := make(chan types.NamespacedName, len(bindings))
	for _, binding := range bindings {
		g.Add(1)
		go func(binding types.NamespacedName) {
			defer g.Done()
			err := sdk.DeleteBinding(binding.Namespace, binding.Name)
			if err == nil {
				deletedBindings <- binding
			}
			errs <- err
		}(binding)
	}

	g.Wait()
	close(errs)
	close(deletedBindings)

	// Collect any errors that occurred into a single formatted error
	bindErr := &multierror.Error{
		ErrorFormat: func(errors []error) string {
			return joinErrors("error:", errors, "\n  ")
		},
	}
	for err := range errs {
		bindErr = multierror.Append(bindErr, err)
	}

	//Range over the deleted bindings to build a slice to return
	deleted := []types.NamespacedName(nil)
	for b := range deletedBindings {
		deleted = append(deleted, b)
	}
	return deleted, bindErr.ErrorOrNil()
}

// DeleteBinding by name.
func (sdk *SDK) DeleteBinding(ns, bindingName string) error {
	err := sdk.ServiceCatalog().ServiceBindings(ns).Delete(context.Background(), bindingName, v1.DeleteOptions{})
	if err != nil {
		return errors.Wrapf(err, "remove binding %s/%s failed", ns, bindingName)
	}
	return nil
}

func joinErrors(groupMsg string, errors []error, sep string, a ...interface{}) string {
	if len(errors) == 0 {
		return ""
	}

	msgs := make([]string, 0, len(errors)+1)
	msgs = append(msgs, fmt.Sprintf(groupMsg, a...))
	for _, err := range errors {
		msgs = append(msgs, err.Error())
	}

	return strings.Join(msgs, sep)
}

// BindingParentHierarchy retrieves all ancestor resources of a binding.
func (sdk *SDK) BindingParentHierarchy(binding *v1.ServiceBinding,
) (*v1.ServiceInstance, *v1.ClusterServiceClass, *v1.ClusterServicePlan, *v1.ClusterServiceBroker, error) {
	instance, err := sdk.RetrieveInstanceByBinding(binding)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	class, plan, err := sdk.InstanceToServiceClassAndPlan(instance)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	broker, err := sdk.RetrieveBrokerByClass(class)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return instance, class, plan, broker, nil
}

// GetBindingStatusCondition returns the last condition on a binding status.
// When no conditions exist, an empty condition is returned.
func GetBindingStatusCondition(status v1.ServiceBindingStatus) v1.ServiceBindingCondition {
	if len(status.Conditions) > 0 {
		return status.Conditions[len(status.Conditions)-1]
	}
	return v1.ServiceBindingCondition{}
}

// WaitForBinding waits for the instance to complete the current operation (or fail).
func (sdk *SDK) WaitForBinding(ns, name string, interval time.Duration, timeout *time.Duration) (binding *v1.ServiceBinding, err error) {
	if timeout == nil {
		notimeout := time.Duration(math.MaxInt64)
		timeout = &notimeout
	}

	err = wait.PollImmediate(interval, *timeout,
		func() (bool, error) {
			binding, err = sdk.RetrieveBinding(ns, name)
			if err != nil {
				return true, err
			}

			if len(binding.Status.Conditions) == 0 {
				return false, nil
			}

			isDone := (sdk.IsBindingReady(binding) || sdk.IsBindingFailed(binding)) && !binding.Status.AsyncOpInProgress
			return isDone, nil
		},
	)

	return binding, err
}

// IsBindingReady returns true if the instance is in the Ready status.
func (sdk *SDK) IsBindingReady(binding *v1.ServiceBinding) bool {
	return sdk.bindingHasStatus(binding, v1.ServiceBindingConditionReady)
}

// IsBindingFailed returns true if the instance is in the Failed status.
func (sdk *SDK) IsBindingFailed(binding *v1.ServiceBinding) bool {
	return sdk.bindingHasStatus(binding, v1.ServiceBindingConditionFailed)
}

// BindingHasStatus returns if the instance is in the specified status.
func (sdk *SDK) bindingHasStatus(binding *v1.ServiceBinding, status v1.ServiceBindingConditionType) bool {
	if binding == nil {
		return false
	}

	for _, cond := range binding.Status.Conditions {
		if cond.Type == status &&
			cond.Status == v1.ConditionTrue {
			return true
		}
	}

	return false
}

// RemoveFinalizerForBinding removes the finalizer for a single binding
func (sdk *SDK) RemoveFinalizerForBinding(namespacedName types.NamespacedName) error {
	// Get binding object from namespacedName
	binding, err := sdk.RetrieveBinding(namespacedName.Namespace, namespacedName.Name)
	if err != nil {
		return err
	}
	finalizers := sets.NewString(binding.Finalizers...)
	finalizers.Delete(v1.FinalizerServiceCatalog)
	binding.Finalizers = finalizers.List()
	_, err = sdk.ServiceCatalog().ServiceBindings(binding.Namespace).Update(context.Background(), binding, v1.UpdateOptions{})
	return err
}

// RemoveFinalizerForBindings removes the finalizer for the provided list of bindings
func (sdk *SDK) RemoveFinalizerForBindings(bindings []types.NamespacedName) ([]types.NamespacedName, error) {
	var g sync.WaitGroup
	deletedBindings := make(chan types.NamespacedName, len(bindings))
	errs := make(chan error, len(bindings))
	for _, binding := range bindings {
		g.Add(1)
		go func(binding types.NamespacedName) {
			defer g.Done()
			err := sdk.RemoveFinalizerForBinding(binding)
			if err == nil {
				deletedBindings <- binding
			}
			errs <- err
		}(binding)
	}

	g.Wait()
	close(errs)
	close(deletedBindings)

	// Collect any errors that occurred into a single formatted error
	bindErr := &multierror.Error{
		ErrorFormat: func(errors []error) string {
			return joinErrors("error:", errors, "\n  ")
		},
	}
	for err := range errs {
		bindErr = multierror.Append(bindErr, err)
	}

	//Range over the deleted bindings to build a slice to return
	deleted := []types.NamespacedName(nil)
	for b := range deletedBindings {
		deleted = append(deleted, b)
	}
	return deleted, bindErr.ErrorOrNil()
}

// RemoveBindingFinalizerByInstance removes v1.FinalizerServiceCatalog from all bindings for the specified instance.
func (sdk *SDK) RemoveBindingFinalizerByInstance(ns, name string) ([]types.NamespacedName, error) {
	instance, err := sdk.RetrieveInstance(ns, name)
	if err != nil {
		return nil, err
	}

	instanceBindings, err := sdk.RetrieveBindingsByInstance(instance)
	if err != nil {
		return nil, err
	}
	namespacedNames := []types.NamespacedName{}
	for _, b := range instanceBindings {
		namespacedNames = append(namespacedNames, types.NamespacedName{Namespace: b.Namespace, Name: b.Name})
	}
	return sdk.RemoveFinalizerForBindings(namespacedNames)
}
