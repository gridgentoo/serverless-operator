// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperatorHubs implements OperatorHubInterface
type FakeOperatorHubs struct {
	Fake *FakeConfigV1
}

var operatorhubsResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "operatorhubs"}

var operatorhubsKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "OperatorHub"}

// Get takes name of the operatorHub, and returns the corresponding operatorHub object, and an error if there is any.
func (c *FakeOperatorHubs) Get(name string, options v1.GetOptions) (result *configv1.OperatorHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(operatorhubsResource, name), &configv1.OperatorHub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OperatorHub), err
}

// List takes label and field selectors, and returns the list of OperatorHubs that match those selectors.
func (c *FakeOperatorHubs) List(opts v1.ListOptions) (result *configv1.OperatorHubList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(operatorhubsResource, operatorhubsKind, opts), &configv1.OperatorHubList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.OperatorHubList{ListMeta: obj.(*configv1.OperatorHubList).ListMeta}
	for _, item := range obj.(*configv1.OperatorHubList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operatorHubs.
func (c *FakeOperatorHubs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(operatorhubsResource, opts))
}

// Create takes the representation of a operatorHub and creates it.  Returns the server's representation of the operatorHub, and an error, if there is any.
func (c *FakeOperatorHubs) Create(operatorHub *configv1.OperatorHub) (result *configv1.OperatorHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(operatorhubsResource, operatorHub), &configv1.OperatorHub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OperatorHub), err
}

// Update takes the representation of a operatorHub and updates it. Returns the server's representation of the operatorHub, and an error, if there is any.
func (c *FakeOperatorHubs) Update(operatorHub *configv1.OperatorHub) (result *configv1.OperatorHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(operatorhubsResource, operatorHub), &configv1.OperatorHub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OperatorHub), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperatorHubs) UpdateStatus(operatorHub *configv1.OperatorHub) (*configv1.OperatorHub, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(operatorhubsResource, "status", operatorHub), &configv1.OperatorHub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OperatorHub), err
}

// Delete takes name of the operatorHub and deletes it. Returns an error if one occurs.
func (c *FakeOperatorHubs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(operatorhubsResource, name), &configv1.OperatorHub{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperatorHubs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(operatorhubsResource, listOptions)

	_, err := c.Fake.Invokes(action, &configv1.OperatorHubList{})
	return err
}

// Patch applies the patch and returns the patched operatorHub.
func (c *FakeOperatorHubs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configv1.OperatorHub, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(operatorhubsResource, name, pt, data, subresources...), &configv1.OperatorHub{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.OperatorHub), err
}