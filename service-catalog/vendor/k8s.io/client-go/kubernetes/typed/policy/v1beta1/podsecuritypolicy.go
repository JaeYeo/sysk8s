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

package v1

import (
	"context"
	"time"

	v1 "k8s.io/api/policy/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// PodSecurityPoliciesGetter has a method to return a PodSecurityPolicyInterface.
// A group's client should implement this interface.
type PodSecurityPoliciesGetter interface {
	PodSecurityPolicies() PodSecurityPolicyInterface
}

// PodSecurityPolicyInterface has methods to work with PodSecurityPolicy resources.
type PodSecurityPolicyInterface interface {
	Create(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.CreateOptions) (*v1.PodSecurityPolicy, error)
	Update(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.UpdateOptions) (*v1.PodSecurityPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1.PodSecurityPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1.PodSecurityPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.PodSecurityPolicy, err error)
	PodSecurityPolicyExpansion
}

// podSecurityPolicies implements PodSecurityPolicyInterface
type podSecurityPolicies struct {
	client rest.Interface
}

// newPodSecurityPolicies returns a PodSecurityPolicies
func newPodSecurityPolicies(c *Policyv1Client) *podSecurityPolicies {
	return &podSecurityPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the podSecurityPolicy, and returns the corresponding podSecurityPolicy object, and an error if there is any.
func (c *podSecurityPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1.PodSecurityPolicy, err error) {
	result = &v1.PodSecurityPolicy{}
	err = c.client.Get().
		Resource("podsecuritypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodSecurityPolicies that match those selectors.
func (c *podSecurityPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1.PodSecurityPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PodSecurityPolicyList{}
	err = c.client.Get().
		Resource("podsecuritypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podSecurityPolicies.
func (c *podSecurityPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("podsecuritypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podSecurityPolicy and creates it.  Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *podSecurityPolicies) Create(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.CreateOptions) (result *v1.PodSecurityPolicy, err error) {
	result = &v1.PodSecurityPolicy{}
	err = c.client.Post().
		Resource("podsecuritypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podSecurityPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podSecurityPolicy and updates it. Returns the server's representation of the podSecurityPolicy, and an error, if there is any.
func (c *podSecurityPolicies) Update(ctx context.Context, podSecurityPolicy *v1.PodSecurityPolicy, opts v1.UpdateOptions) (result *v1.PodSecurityPolicy, err error) {
	result = &v1.PodSecurityPolicy{}
	err = c.client.Put().
		Resource("podsecuritypolicies").
		Name(podSecurityPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podSecurityPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podSecurityPolicy and deletes it. Returns an error if one occurs.
func (c *podSecurityPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("podsecuritypolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podSecurityPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("podsecuritypolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podSecurityPolicy.
func (c *podSecurityPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.PodSecurityPolicy, err error) {
	result = &v1.PodSecurityPolicy{}
	err = c.client.Patch(pt).
		Resource("podsecuritypolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
