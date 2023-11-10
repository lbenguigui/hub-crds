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

	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPICollections implements APICollectionInterface
type FakeAPICollections struct {
	Fake *FakeHubV1alpha1
}

var apicollectionsResource = schema.GroupVersionResource{Group: "hub.traefik.io", Version: "v1alpha1", Resource: "apicollections"}

var apicollectionsKind = schema.GroupVersionKind{Group: "hub.traefik.io", Version: "v1alpha1", Kind: "APICollection"}

// Get takes name of the aPICollection, and returns the corresponding aPICollection object, and an error if there is any.
func (c *FakeAPICollections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APICollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apicollectionsResource, name), &v1alpha1.APICollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APICollection), err
}

// List takes label and field selectors, and returns the list of APICollections that match those selectors.
func (c *FakeAPICollections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APICollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apicollectionsResource, apicollectionsKind, opts), &v1alpha1.APICollectionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APICollectionList{ListMeta: obj.(*v1alpha1.APICollectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.APICollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPICollections.
func (c *FakeAPICollections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apicollectionsResource, opts))
}

// Create takes the representation of a aPICollection and creates it.  Returns the server's representation of the aPICollection, and an error, if there is any.
func (c *FakeAPICollections) Create(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.CreateOptions) (result *v1alpha1.APICollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apicollectionsResource, aPICollection), &v1alpha1.APICollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APICollection), err
}

// Update takes the representation of a aPICollection and updates it. Returns the server's representation of the aPICollection, and an error, if there is any.
func (c *FakeAPICollections) Update(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (result *v1alpha1.APICollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apicollectionsResource, aPICollection), &v1alpha1.APICollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APICollection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPICollections) UpdateStatus(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (*v1alpha1.APICollection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apicollectionsResource, "status", aPICollection), &v1alpha1.APICollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APICollection), err
}

// Delete takes name of the aPICollection and deletes it. Returns an error if one occurs.
func (c *FakeAPICollections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(apicollectionsResource, name), &v1alpha1.APICollection{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPICollections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apicollectionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APICollectionList{})
	return err
}

// Patch applies the patch and returns the patched aPICollection.
func (c *FakeAPICollections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APICollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apicollectionsResource, name, pt, data, subresources...), &v1alpha1.APICollection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APICollection), err
}
