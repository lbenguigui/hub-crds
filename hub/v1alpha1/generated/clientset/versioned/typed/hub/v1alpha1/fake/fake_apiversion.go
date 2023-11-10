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

// FakeAPIVersions implements APIVersionInterface
type FakeAPIVersions struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var apiversionsResource = schema.GroupVersionResource{Group: "hub.traefik.io", Version: "v1alpha1", Resource: "apiversions"}

var apiversionsKind = schema.GroupVersionKind{Group: "hub.traefik.io", Version: "v1alpha1", Kind: "APIVersion"}

// Get takes name of the aPIVersion, and returns the corresponding aPIVersion object, and an error if there is any.
func (c *FakeAPIVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiversionsResource, c.ns, name), &v1alpha1.APIVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIVersion), err
}

// List takes label and field selectors, and returns the list of APIVersions that match those selectors.
func (c *FakeAPIVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiversionsResource, apiversionsKind, c.ns, opts), &v1alpha1.APIVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIVersionList{ListMeta: obj.(*v1alpha1.APIVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIVersions.
func (c *FakeAPIVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiversionsResource, c.ns, opts))

}

// Create takes the representation of a aPIVersion and creates it.  Returns the server's representation of the aPIVersion, and an error, if there is any.
func (c *FakeAPIVersions) Create(ctx context.Context, aPIVersion *v1alpha1.APIVersion, opts v1.CreateOptions) (result *v1alpha1.APIVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiversionsResource, c.ns, aPIVersion), &v1alpha1.APIVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIVersion), err
}

// Update takes the representation of a aPIVersion and updates it. Returns the server's representation of the aPIVersion, and an error, if there is any.
func (c *FakeAPIVersions) Update(ctx context.Context, aPIVersion *v1alpha1.APIVersion, opts v1.UpdateOptions) (result *v1alpha1.APIVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiversionsResource, c.ns, aPIVersion), &v1alpha1.APIVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIVersions) UpdateStatus(ctx context.Context, aPIVersion *v1alpha1.APIVersion, opts v1.UpdateOptions) (*v1alpha1.APIVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiversionsResource, "status", c.ns, aPIVersion), &v1alpha1.APIVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIVersion), err
}

// Delete takes name of the aPIVersion and deletes it. Returns an error if one occurs.
func (c *FakeAPIVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apiversionsResource, c.ns, name), &v1alpha1.APIVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIVersionList{})
	return err
}

// Patch applies the patch and returns the patched aPIVersion.
func (c *FakeAPIVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.APIVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIVersion), err
}
