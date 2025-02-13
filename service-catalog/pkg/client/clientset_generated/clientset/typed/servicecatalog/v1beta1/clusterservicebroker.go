/*
Copyright 2021 The Kubernetes Authors.

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

	v1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1"
	scheme "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterServiceBrokersGetter has a method to return a ClusterServiceBrokerInterface.
// A group's client should implement this interface.
type ClusterServiceBrokersGetter interface {
	ClusterServiceBrokers() ClusterServiceBrokerInterface
}

// ClusterServiceBrokerInterface has methods to work with ClusterServiceBroker resources.
type ClusterServiceBrokerInterface interface {
	Create(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.CreateOptions) (*v1.ClusterServiceBroker, error)
	Update(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.UpdateOptions) (*v1.ClusterServiceBroker, error)
	UpdateStatus(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.UpdateOptions) (*v1.ClusterServiceBroker, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1.ClusterServiceBroker, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1.ClusterServiceBrokerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.ClusterServiceBroker, err error)
	ClusterServiceBrokerExpansion
}

// clusterServiceBrokers implements ClusterServiceBrokerInterface
type clusterServiceBrokers struct {
	client rest.Interface
}

// newClusterServiceBrokers returns a ClusterServiceBrokers
func newClusterServiceBrokers(c *Servicecatalogv1Client) *clusterServiceBrokers {
	return &clusterServiceBrokers{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterServiceBroker, and returns the corresponding clusterServiceBroker object, and an error if there is any.
func (c *clusterServiceBrokers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1.ClusterServiceBroker, err error) {
	result = &v1.ClusterServiceBroker{}
	err = c.client.Get().
		Resource("clusterservicebrokers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterServiceBrokers that match those selectors.
func (c *clusterServiceBrokers) List(ctx context.Context, opts v1.ListOptions) (result *v1.ClusterServiceBrokerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterServiceBrokerList{}
	err = c.client.Get().
		Resource("clusterservicebrokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterServiceBrokers.
func (c *clusterServiceBrokers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterservicebrokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterServiceBroker and creates it.  Returns the server's representation of the clusterServiceBroker, and an error, if there is any.
func (c *clusterServiceBrokers) Create(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.CreateOptions) (result *v1.ClusterServiceBroker, err error) {
	result = &v1.ClusterServiceBroker{}
	err = c.client.Post().
		Resource("clusterservicebrokers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServiceBroker).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterServiceBroker and updates it. Returns the server's representation of the clusterServiceBroker, and an error, if there is any.
func (c *clusterServiceBrokers) Update(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.UpdateOptions) (result *v1.ClusterServiceBroker, err error) {
	result = &v1.ClusterServiceBroker{}
	err = c.client.Put().
		Resource("clusterservicebrokers").
		Name(clusterServiceBroker.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServiceBroker).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterServiceBrokers) UpdateStatus(ctx context.Context, clusterServiceBroker *v1.ClusterServiceBroker, opts v1.UpdateOptions) (result *v1.ClusterServiceBroker, err error) {
	result = &v1.ClusterServiceBroker{}
	err = c.client.Put().
		Resource("clusterservicebrokers").
		Name(clusterServiceBroker.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterServiceBroker).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterServiceBroker and deletes it. Returns an error if one occurs.
func (c *clusterServiceBrokers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterservicebrokers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterServiceBrokers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterservicebrokers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterServiceBroker.
func (c *clusterServiceBrokers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1.ClusterServiceBroker, err error) {
	result = &v1.ClusterServiceBroker{}
	err = c.client.Patch(pt).
		Resource("clusterservicebrokers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
