/*
Copyright 2021 The Clusternet Authors.

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

	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	scheme "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BasesGetter has a method to return a BaseInterface.
// A group's client should implement this interface.
type BasesGetter interface {
	Bases(namespace string) BaseInterface
}

// BaseInterface has methods to work with Base resources.
type BaseInterface interface {
	Create(ctx context.Context, base *v1alpha1.Base, opts v1.CreateOptions) (*v1alpha1.Base, error)
	Update(ctx context.Context, base *v1alpha1.Base, opts v1.UpdateOptions) (*v1alpha1.Base, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Base, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Base, err error)
	BaseExpansion
}

// bases implements BaseInterface
type bases struct {
	client rest.Interface
	ns     string
}

// newBases returns a Bases
func newBases(c *AppsV1alpha1Client, namespace string) *bases {
	return &bases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the base, and returns the corresponding base object, and an error if there is any.
func (c *bases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Base, err error) {
	result = &v1alpha1.Base{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Bases that match those selectors.
func (c *bases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bases.
func (c *bases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a base and creates it.  Returns the server's representation of the base, and an error, if there is any.
func (c *bases) Create(ctx context.Context, base *v1alpha1.Base, opts v1.CreateOptions) (result *v1alpha1.Base, err error) {
	result = &v1alpha1.Base{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(base).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a base and updates it. Returns the server's representation of the base, and an error, if there is any.
func (c *bases) Update(ctx context.Context, base *v1alpha1.Base, opts v1.UpdateOptions) (result *v1alpha1.Base, err error) {
	result = &v1alpha1.Base{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bases").
		Name(base.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(base).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the base and deletes it. Returns an error if one occurs.
func (c *bases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched base.
func (c *bases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Base, err error) {
	result = &v1alpha1.Base{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
