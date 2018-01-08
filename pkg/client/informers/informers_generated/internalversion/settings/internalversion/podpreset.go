/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	settings "k8s.io/kubernetes/pkg/apis/settings"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/settings/internalversion"
	time "time"
)

// PodPresetInformer provides access to a shared informer and lister for
// PodPresets.
type PodPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PodPresetLister
}

type podPresetInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewPodPresetInformer constructs a new informer for PodPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodPresetInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Settings().PodPresets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Settings().PodPresets(namespace).Watch(options)
			},
		},
		&settings.PodPreset{},
		resyncPeriod,
		indexers,
	)
}

func defaultPodPresetInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewPodPresetInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *podPresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&settings.PodPreset{}, defaultPodPresetInformer)
}

func (f *podPresetInformer) Lister() internalversion.PodPresetLister {
	return internalversion.NewPodPresetLister(f.Informer().GetIndexer())
}
