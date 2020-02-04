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

	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// RolesGetter has a method to return a RoleInterface.
// A group's client should implement this interface.
type RolesGetter interface {
	Roles(namespace string) RoleInterface
}

// RoleInterface has methods to work with Role resources.
type RoleInterface interface {
	Create(context.Context, *v1.Role) (*v1.Role, error)
	Update(context.Context, *v1.Role) (*v1.Role, error)
	Delete(ctx context.Context, name string, options *metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Role, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RoleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Role, err error)
	RoleExpansion
}

// roles implements RoleInterface
type roles struct {
	client rest.Interface
	ns     string
}

// newRoles returns a Roles
func newRoles(c *RbacV1Client, namespace string) *roles {
	return &roles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the role, and returns the corresponding role object, and an error if there is any.
func (c *roles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Role, err error) {
	result = &v1.Role{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("roles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *roles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("roles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested roles.
func (c *roles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("roles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a role and creates it.  Returns the server's representation of the role, and an error, if there is any.
func (c *roles) Create(ctx context.Context, role *v1.Role) (result *v1.Role, err error) {
	result = &v1.Role{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("roles").
		Body(role).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a role and updates it. Returns the server's representation of the role, and an error, if there is any.
func (c *roles) Update(ctx context.Context, role *v1.Role) (result *v1.Role, err error) {
	result = &v1.Role{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("roles").
		Name(role.Name).
		Body(role).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the role and deletes it. Returns an error if one occurs.
func (c *roles) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("roles").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *roles) DeleteCollection(ctx context.Context, options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("roles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched role.
func (c *roles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Role, err error) {
	result = &v1.Role{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("roles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
