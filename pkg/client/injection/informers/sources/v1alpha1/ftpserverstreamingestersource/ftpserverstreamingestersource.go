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

// Code generated by injection-gen. DO NOT EDIT.

package samplesource

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
	apissamplesv1alpha1 "github.com/skilld-labs/stream-event-decoder-sources/pkg/apis/samples/v1alpha1"
	versioned "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/clientset/versioned"
	v1alpha1 "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/informers/externalversions/samples/v1alpha1"
	client "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/injection/client"
	factory "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/injection/informers/factory"
	samplesv1alpha1 "github.com/skilld-labs/stream-event-decoder-sources/pkg/client/listers/samples/v1alpha1"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Samples().V1alpha1().SampleSources()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

func withDynamicInformer(ctx context.Context) context.Context {
	inf := &wrapper{client: client.Get(ctx)}
	return context.WithValue(ctx, Key{}, inf)
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1alpha1.SampleSourceInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/skilld-labs/stream-event-decoder-sources/pkg/client/informers/externalversions/samples/v1alpha1.SampleSourceInformer from context.")
	}
	return untyped.(v1alpha1.SampleSourceInformer)
}

type wrapper struct {
	client versioned.Interface

	namespace string
}

var _ v1alpha1.SampleSourceInformer = (*wrapper)(nil)
var _ samplesv1alpha1.SampleSourceLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apissamplesv1alpha1.SampleSource{}, 0, nil)
}

func (w *wrapper) Lister() samplesv1alpha1.SampleSourceLister {
	return w
}

func (w *wrapper) SampleSources(namespace string) samplesv1alpha1.SampleSourceNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apissamplesv1alpha1.SampleSource, err error) {
	lo, err := w.client.SamplesV1alpha1().SampleSources(w.namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: selector.String(),
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apissamplesv1alpha1.SampleSource, error) {
	return w.client.SamplesV1alpha1().SampleSources(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}
