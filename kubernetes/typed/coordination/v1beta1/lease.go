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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "k8s.io/api/coordination/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// LeasesGetter has a method to return a LeaseInterface.
// A group's client should implement this interface.
type LeasesGetter interface {
	Leases(namespace string) LeaseInterface
}

// LeaseInterface has methods to work with Lease resources.
type LeaseInterface interface {
	Create(context.Context, *v1beta1.Lease) (*v1beta1.Lease, error)
	Update(context.Context, *v1beta1.Lease) (*v1beta1.Lease, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*v1beta1.Lease, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.LeaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Lease, err error)
	LeaseExpansion
}

// leases implements LeaseInterface
type leases struct {
	client rest.Interface
	ns     string
}

// newLeases returns a Leases
func newLeases(c *CoordinationV1beta1Client, namespace string) *leases {
	return &leases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lease, and returns the corresponding lease object, and an error if there is any.
func (c *leases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Lease, err error) {
	result = &v1beta1.Lease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Leases that match those selectors.
func (c *leases) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.LeaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.LeaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested leases.
func (c *leases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lease and creates it.  Returns the server's representation of the lease, and an error, if there is any.
func (c *leases) Create(ctx context.Context, lease *v1beta1.Lease) (result *v1beta1.Lease, err error) {
	result = &v1beta1.Lease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("leases").
		Body(lease).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lease and updates it. Returns the server's representation of the lease, and an error, if there is any.
func (c *leases) Update(ctx context.Context, lease *v1beta1.Lease) (result *v1beta1.Lease, err error) {
	result = &v1beta1.Lease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("leases").
		Name(lease.Name).
		Body(lease).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lease and deletes it. Returns an error if one occurs.
func (c *leases) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leases").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *leases) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("leases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lease.
func (c *leases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Lease, err error) {
	result = &v1beta1.Lease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("leases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
