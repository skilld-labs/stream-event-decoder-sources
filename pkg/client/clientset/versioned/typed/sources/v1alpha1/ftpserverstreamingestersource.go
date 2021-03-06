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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/skilld-labs/stream-event-decoder-sources/pkg/apis/samples/v1alpha1"
	scheme "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/clientset/versioned/scheme"
)

// SampleSourcesGetter has a method to return a SampleSourceInterface.
// A group's client should implement this interface.
type SampleSourcesGetter interface {
	SampleSources(namespace string) SampleSourceInterface
}

// SampleSourceInterface has methods to work with SampleSource resources.
type SampleSourceInterface interface {
	Create(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.CreateOptions) (*v1alpha1.SampleSource, error)
	Update(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (*v1alpha1.SampleSource, error)
	UpdateStatus(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (*v1alpha1.SampleSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SampleSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SampleSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SampleSource, err error)
	SampleSourceExpansion
}

// sampleSources implements SampleSourceInterface
type sampleSources struct {
	client rest.Interface
	ns     string
}

// newSampleSources returns a SampleSources
func newSampleSources(c *SourcesV1alpha1Client, namespace string) *sampleSources {
	return &sampleSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sampleSource, and returns the corresponding sampleSource object, and an error if there is any.
func (c *sampleSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SampleSource, err error) {
	result = &v1alpha1.SampleSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samplesources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SampleSources that match those selectors.
func (c *sampleSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SampleSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SampleSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samplesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sampleSources.
func (c *sampleSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("samplesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sampleSource and creates it.  Returns the server's representation of the sampleSource, and an error, if there is any.
func (c *sampleSources) Create(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.CreateOptions) (result *v1alpha1.SampleSource, err error) {
	result = &v1alpha1.SampleSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("samplesources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sampleSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sampleSource and updates it. Returns the server's representation of the sampleSource, and an error, if there is any.
func (c *sampleSources) Update(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (result *v1alpha1.SampleSource, err error) {
	result = &v1alpha1.SampleSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("samplesources").
		Name(sampleSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sampleSource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sampleSources) UpdateStatus(ctx context.Context, sampleSource *v1alpha1.SampleSource, opts v1.UpdateOptions) (result *v1alpha1.SampleSource, err error) {
	result = &v1alpha1.SampleSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("samplesources").
		Name(sampleSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sampleSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sampleSource and deletes it. Returns an error if one occurs.
func (c *sampleSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samplesources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sampleSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samplesources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sampleSource.
func (c *sampleSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SampleSource, err error) {
	result = &v1alpha1.SampleSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("samplesources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
