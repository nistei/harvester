/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodePools implements NodePoolInterface
type FakeNodePools struct {
	Fake *FakeManagementV3
	ns   string
}

var nodepoolsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "nodepools"}

var nodepoolsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "NodePool"}

// Get takes name of the nodePool, and returns the corresponding nodePool object, and an error if there is any.
func (c *FakeNodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodepoolsResource, c.ns, name), &v3.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodePool), err
}

// List takes label and field selectors, and returns the list of NodePools that match those selectors.
func (c *FakeNodePools) List(ctx context.Context, opts v1.ListOptions) (result *v3.NodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodepoolsResource, nodepoolsKind, c.ns, opts), &v3.NodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.NodePoolList{ListMeta: obj.(*v3.NodePoolList).ListMeta}
	for _, item := range obj.(*v3.NodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodePools.
func (c *FakeNodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodepoolsResource, c.ns, opts))

}

// Create takes the representation of a nodePool and creates it.  Returns the server's representation of the nodePool, and an error, if there is any.
func (c *FakeNodePools) Create(ctx context.Context, nodePool *v3.NodePool, opts v1.CreateOptions) (result *v3.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodepoolsResource, c.ns, nodePool), &v3.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodePool), err
}

// Update takes the representation of a nodePool and updates it. Returns the server's representation of the nodePool, and an error, if there is any.
func (c *FakeNodePools) Update(ctx context.Context, nodePool *v3.NodePool, opts v1.UpdateOptions) (result *v3.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodepoolsResource, c.ns, nodePool), &v3.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodePools) UpdateStatus(ctx context.Context, nodePool *v3.NodePool, opts v1.UpdateOptions) (*v3.NodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodepoolsResource, "status", c.ns, nodePool), &v3.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodePool), err
}

// Delete takes name of the nodePool and deletes it. Returns an error if one occurs.
func (c *FakeNodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nodepoolsResource, c.ns, name, opts), &v3.NodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodepoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.NodePoolList{})
	return err
}

// Patch applies the patch and returns the patched nodePool.
func (c *FakeNodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodepoolsResource, c.ns, name, pt, data, subresources...), &v3.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.NodePool), err
}
