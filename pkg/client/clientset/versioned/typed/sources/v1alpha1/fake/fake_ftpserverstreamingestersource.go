/*
Copyright 2020 The Knative Authors

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "github.com/skilld-labs/stream-event-decoder-sources/pkg/apis/samples/v1alpha1"
)

// FakeSampleSources implements SampleSourceInterface
type FakeSampleSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var samplesourcesResource = schema.GroupVersionResource{Group: "samples.knative.dev", Version: "v1alpha1", Resource: "samplesources"}

var samplesourcesKind = schema.GroupVersionKind{Group: "samples.knative.dev", Version: "v1alpha1", Kind: "SampleSource"}

// Get takes name of the sampleSource, and returns the corresponding sampleSource object, and an error if there is any.
func (c *FakeSampleSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SampleSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(samplesourcesResource, c.ns, name), &v1alpha1.SampleSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleSource), err
}

// List takes label and field selectors, and returns the list of SampleSources that match those selectors.
func (c *FakeSampleSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SampleSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(samplesourcesResource, samplesourcesKind, c.ns, opts), &v1alpha1.SampleSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SampleSourceList{ListMeta: obj.(*v1alpha1.SampleSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SampleSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sampleSources.
func (c *FakeSampleSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(samplesourcesResource, c.ns, opts))

}

// Create takes the representation of a sampleSource and creates it.  Returns the server's representation of the sampleSource, and an error, if there is any.
func (c *FakeSampleSources) Create(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.CreateOptions) (result *v1alpha1.SampleSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(samplesourcesResource, c.ns, sampleSource), &v1alpha1.SampleSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleSource), err
}

// Update takes the representation of a sampleSource and updates it. Returns the server's representation of the sampleSource, and an error, if there is any.
func (c *FakeSampleSources) Update(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (result *v1alpha1.SampleSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(samplesourcesResource, c.ns, sampleSource), &v1alpha1.SampleSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSampleSources) UpdateStatus(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (*v1alpha1.SampleSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(samplesourcesResource, "status", c.ns, sampleSource), &v1alpha1.SampleSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleSource), err
}

// Delete takes name of the sampleSource and deletes it. Returns an error if one occurs.
func (c *FakeSampleSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(samplesourcesResource, c.ns, name), &v1alpha1.SampleSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSampleSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(samplesourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SampleSourceList{})
	return err
}

// Patch applies the patch and returns the patched sampleSource.
func (c *FakeSampleSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SampleSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(samplesourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SampleSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleSource), err
}
