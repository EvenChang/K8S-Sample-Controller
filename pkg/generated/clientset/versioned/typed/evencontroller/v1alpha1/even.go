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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/evencontroller/v1alpha1"
	scheme "k8s.io/sample-controller/pkg/generated/clientset/versioned/scheme"
)

// EvensGetter has a method to return a EvenInterface.
// A group's client should implement this interface.
type EvensGetter interface {
	Evens(namespace string) EvenInterface
}

// EvenInterface has methods to work with Even resources.
type EvenInterface interface {
	Create(ctx context.Context, even *v1alpha1.Even, opts v1.CreateOptions) (*v1alpha1.Even, error)
	Update(ctx context.Context, even *v1alpha1.Even, opts v1.UpdateOptions) (*v1alpha1.Even, error)
	UpdateStatus(ctx context.Context, even *v1alpha1.Even, opts v1.UpdateOptions) (*v1alpha1.Even, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Even, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EvenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Even, err error)
	EvenExpansion
}

// evens implements EvenInterface
type evens struct {
	client rest.Interface
	ns     string
}

// newEvens returns a Evens
func newEvens(c *EvencontrollerV1alpha1Client, namespace string) *evens {
	return &evens{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the even, and returns the corresponding even object, and an error if there is any.
func (c *evens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Even, err error) {
	result = &v1alpha1.Even{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("evens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Evens that match those selectors.
func (c *evens) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EvenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EvenList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("evens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested evens.
func (c *evens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("evens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a even and creates it.  Returns the server's representation of the even, and an error, if there is any.
func (c *evens) Create(ctx context.Context, even *v1alpha1.Even, opts v1.CreateOptions) (result *v1alpha1.Even, err error) {
	result = &v1alpha1.Even{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("evens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(even).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a even and updates it. Returns the server's representation of the even, and an error, if there is any.
func (c *evens) Update(ctx context.Context, even *v1alpha1.Even, opts v1.UpdateOptions) (result *v1alpha1.Even, err error) {
	result = &v1alpha1.Even{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("evens").
		Name(even.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(even).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *evens) UpdateStatus(ctx context.Context, even *v1alpha1.Even, opts v1.UpdateOptions) (result *v1alpha1.Even, err error) {
	result = &v1alpha1.Even{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("evens").
		Name(even.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(even).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the even and deletes it. Returns an error if one occurs.
func (c *evens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("evens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *evens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("evens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched even.
func (c *evens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Even, err error) {
	result = &v1alpha1.Even{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("evens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
