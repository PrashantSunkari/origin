package fake

import (
	build_v1 "github.com/openshift/origin/pkg/build/apis/build/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilds implements BuildResourceInterface
type FakeBuilds struct {
	Fake *FakeBuildV1
	ns   string
}

var buildsResource = schema.GroupVersionResource{Group: "build.openshift.io", Version: "v1", Resource: "builds"}

var buildsKind = schema.GroupVersionKind{Group: "build.openshift.io", Version: "v1", Kind: "Build"}

// Get takes name of the buildResource, and returns the corresponding buildResource object, and an error if there is any.
func (c *FakeBuilds) Get(name string, options v1.GetOptions) (result *build_v1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildsResource, c.ns, name), &build_v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*build_v1.Build), err
}

// List takes label and field selectors, and returns the list of Builds that match those selectors.
func (c *FakeBuilds) List(opts v1.ListOptions) (result *build_v1.BuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildsResource, buildsKind, c.ns, opts), &build_v1.BuildList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &build_v1.BuildList{}
	for _, item := range obj.(*build_v1.BuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builds.
func (c *FakeBuilds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildsResource, c.ns, opts))

}

// Create takes the representation of a buildResource and creates it.  Returns the server's representation of the buildResource, and an error, if there is any.
func (c *FakeBuilds) Create(buildResource *build_v1.Build) (result *build_v1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildsResource, c.ns, buildResource), &build_v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*build_v1.Build), err
}

// Update takes the representation of a buildResource and updates it. Returns the server's representation of the buildResource, and an error, if there is any.
func (c *FakeBuilds) Update(buildResource *build_v1.Build) (result *build_v1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildsResource, c.ns, buildResource), &build_v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*build_v1.Build), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuilds) UpdateStatus(buildResource *build_v1.Build) (*build_v1.Build, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildsResource, "status", c.ns, buildResource), &build_v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*build_v1.Build), err
}

// Delete takes name of the buildResource and deletes it. Returns an error if one occurs.
func (c *FakeBuilds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildsResource, c.ns, name), &build_v1.Build{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuilds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &build_v1.BuildList{})
	return err
}

// Patch applies the patch and returns the patched buildResource.
func (c *FakeBuilds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *build_v1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildsResource, c.ns, name, data, subresources...), &build_v1.Build{})

	if obj == nil {
		return nil, err
	}
	return obj.(*build_v1.Build), err
}
