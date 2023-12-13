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
	v1 "perfdm/build/apis/root.sockshop.com/v1"
	scheme "perfdm/build/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SockShopsGetter has a method to return a SockShopInterface.
// A group's client should implement this interface.
type SockShopsGetter interface {
	SockShops() SockShopInterface
}

// SockShopInterface has methods to work with SockShop resources.
type SockShopInterface interface {
	Create(ctx context.Context, sockShop *v1.SockShop, opts metav1.CreateOptions) (*v1.SockShop, error)
	Update(ctx context.Context, sockShop *v1.SockShop, opts metav1.UpdateOptions) (*v1.SockShop, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SockShop, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SockShopList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SockShop, err error)
	SockShopExpansion
}

// sockShops implements SockShopInterface
type sockShops struct {
	client rest.Interface
}

// newSockShops returns a SockShops
func newSockShops(c *RootSockshopV1Client) *sockShops {
	return &sockShops{
		client: c.RESTClient(),
	}
}

// Get takes name of the sockShop, and returns the corresponding sockShop object, and an error if there is any.
func (c *sockShops) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SockShop, err error) {
	result = &v1.SockShop{}
	err = c.client.Get().
		Resource("sockshops").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SockShops that match those selectors.
func (c *sockShops) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SockShopList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SockShopList{}
	err = c.client.Get().
		Resource("sockshops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sockShops.
func (c *sockShops) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sockshops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sockShop and creates it.  Returns the server's representation of the sockShop, and an error, if there is any.
func (c *sockShops) Create(ctx context.Context, sockShop *v1.SockShop, opts metav1.CreateOptions) (result *v1.SockShop, err error) {
	result = &v1.SockShop{}
	err = c.client.Post().
		Resource("sockshops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sockShop).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sockShop and updates it. Returns the server's representation of the sockShop, and an error, if there is any.
func (c *sockShops) Update(ctx context.Context, sockShop *v1.SockShop, opts metav1.UpdateOptions) (result *v1.SockShop, err error) {
	result = &v1.SockShop{}
	err = c.client.Put().
		Resource("sockshops").
		Name(sockShop.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sockShop).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sockShop and deletes it. Returns an error if one occurs.
func (c *sockShops) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sockshops").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sockShops) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sockshops").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sockShop.
func (c *sockShops) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SockShop, err error) {
	result = &v1.SockShop{}
	err = c.client.Patch(pt).
		Resource("sockshops").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
