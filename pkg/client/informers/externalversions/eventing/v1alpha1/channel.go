/*
Copyright 2018 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	eventingv1alpha1 "github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
	versioned "github.com/knative/eventing/pkg/client/clientset/versioned"
	internalinterfaces "github.com/knative/eventing/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/knative/eventing/pkg/client/listers/eventing/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ChannelInformer provides access to a shared informer and lister for
// Channels.
type ChannelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ChannelLister
}

type channelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChannelInformer constructs a new informer for Channel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChannelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredChannelInformer constructs a new informer for Channel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1alpha1().Channels(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventingV1alpha1().Channels(namespace).Watch(options)
			},
		},
		&eventingv1alpha1.Channel{},
		resyncPeriod,
		indexers,
	)
}

func (f *channelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredChannelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *channelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventingv1alpha1.Channel{}, f.defaultInformer)
}

func (f *channelInformer) Lister() v1alpha1.ChannelLister {
	return v1alpha1.NewChannelLister(f.Informer().GetIndexer())
}
