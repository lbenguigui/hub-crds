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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	hubv1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	versioned "github.com/traefik/hub-crds/hub/v1alpha1/generated/clientset/versioned"
	internalinterfaces "github.com/traefik/hub-crds/hub/v1alpha1/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1/generated/listers/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// APIRateLimitInformer provides access to a shared informer and lister for
// APIRateLimits.
type APIRateLimitInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.APIRateLimitLister
}

type aPIRateLimitInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAPIRateLimitInformer constructs a new informer for APIRateLimit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIRateLimitInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIRateLimitInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAPIRateLimitInformer constructs a new informer for APIRateLimit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIRateLimitInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIRateLimits().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIRateLimits().Watch(context.TODO(), options)
			},
		},
		&hubv1alpha1.APIRateLimit{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIRateLimitInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIRateLimitInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aPIRateLimitInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hubv1alpha1.APIRateLimit{}, f.defaultInformer)
}

func (f *aPIRateLimitInformer) Lister() v1alpha1.APIRateLimitLister {
	return v1alpha1.NewAPIRateLimitLister(f.Informer().GetIndexer())
}
