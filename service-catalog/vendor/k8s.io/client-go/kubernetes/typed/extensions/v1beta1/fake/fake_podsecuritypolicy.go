/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/api/extensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodSecurityPolicies implements PodSecurityPolicyInterface
type FakePodSecurityPolicies struct {
	Fake *FakeExtensionsv1
}

var podsecuritypoliciesResource = schema.GroupVersionResource{Group: "extensions", Version: "v1", Resource: "podsecuritypolicies"}

var podsecuritypoliciesKind = schema.GroupVersionKind{Group: "extensions", Version: "v1", Kind: "PodSecurityPolicy"}

// Get takes name of the podSecurityPolicy, and returns the corresponding podSecurityPolicy object, and an error if there is any.
func (c *FakePodSecurityPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(podsecuritypoliciesResource, name), &v1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodSecurityPolicy), err
}

// List takes label and field selectors, and returns the list of PodSecurityPolicies that match those selectors.
func (c *FakePodSecurityPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1.PodSecurityPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(podsecuritypoliciesResource, podsecuritypoliciesKind, opts), &v1.PodSecurityPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PodSecurityPolicyList{ListMeta: obj.(*v1.PodSecurityPolicyList).ListMeta}
	for _, item := range obj.(*v1.PodSecurityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podSecurityPolicies.
func (c *FakePodSecurityPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(podsecuritypoliciesResource, opts))
}

// Create takes the representation of a podSecurityPolicy and creates it.  Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *FakePodSecurityPolicies) Create(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.CreateOptions) (result *v1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(podsecuritypoliciesResource, podSecurityPolicy), &v1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodSecurityPolicy), err
}

// Update takes the representation of a podSecurityPolicy and updates it. Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *FakePodSecurityPolicies) Update(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.UpdateOptions) (result *v1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(podsecuritypoliciesResource, podSecurityPolicy), &v1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodSecurityPolicy), err
}

// Delete takes name of the podSecurityPolicy and deletes it. Returns an error if one occurs.
func (c *FakePodSecurityPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(podsecuritypoliciesResource, name), &v1.PodSecurityPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodSecurityPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(podsecuritypoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PodSecurityPolicyList{})
	return err
}

// Patch applies the patch and returns the patched podSecurityPolicy.
func (c *FakePodSecurityPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.PodSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(podsecuritypoliciesResource, name, pt, data, subresources...), &v1.PodSecurityPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodSecurityPolicy), err
}
