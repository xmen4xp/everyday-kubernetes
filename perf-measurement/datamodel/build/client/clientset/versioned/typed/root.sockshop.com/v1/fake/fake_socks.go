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
	rootsockshopcomv1 "perfdm/build/apis/root.sockshop.com/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSockses implements SocksInterface
type FakeSockses struct {
	Fake *FakeRootSockshopV1
}

var socksesResource = schema.GroupVersionResource{Group: "root.sockshop.com", Version: "v1", Resource: "sockses"}

var socksesKind = schema.GroupVersionKind{Group: "root.sockshop.com", Version: "v1", Kind: "Socks"}

// Get takes name of the socks, and returns the corresponding socks object, and an error if there is any.
func (c *FakeSockses) Get(ctx context.Context, name string, options v1.GetOptions) (result *rootsockshopcomv1.Socks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(socksesResource, name), &rootsockshopcomv1.Socks{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rootsockshopcomv1.Socks), err
}

// List takes label and field selectors, and returns the list of Sockses that match those selectors.
func (c *FakeSockses) List(ctx context.Context, opts v1.ListOptions) (result *rootsockshopcomv1.SocksList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(socksesResource, socksesKind, opts), &rootsockshopcomv1.SocksList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rootsockshopcomv1.SocksList{ListMeta: obj.(*rootsockshopcomv1.SocksList).ListMeta}
	for _, item := range obj.(*rootsockshopcomv1.SocksList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sockses.
func (c *FakeSockses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(socksesResource, opts))
}

// Create takes the representation of a socks and creates it.  Returns the server's representation of the socks, and an error, if there is any.
func (c *FakeSockses) Create(ctx context.Context, socks *rootsockshopcomv1.Socks, opts v1.CreateOptions) (result *rootsockshopcomv1.Socks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(socksesResource, socks), &rootsockshopcomv1.Socks{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rootsockshopcomv1.Socks), err
}

// Update takes the representation of a socks and updates it. Returns the server's representation of the socks, and an error, if there is any.
func (c *FakeSockses) Update(ctx context.Context, socks *rootsockshopcomv1.Socks, opts v1.UpdateOptions) (result *rootsockshopcomv1.Socks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(socksesResource, socks), &rootsockshopcomv1.Socks{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rootsockshopcomv1.Socks), err
}

// Delete takes name of the socks and deletes it. Returns an error if one occurs.
func (c *FakeSockses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(socksesResource, name, opts), &rootsockshopcomv1.Socks{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSockses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(socksesResource, listOpts)

	_, err := c.Fake.Invokes(action, &rootsockshopcomv1.SocksList{})
	return err
}

// Patch applies the patch and returns the patched socks.
func (c *FakeSockses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rootsockshopcomv1.Socks, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(socksesResource, name, pt, data, subresources...), &rootsockshopcomv1.Socks{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rootsockshopcomv1.Socks), err
}