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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EdgeIngressLister helps list EdgeIngresses.
// All objects returned here must be treated as read-only.
type EdgeIngressLister interface {
	// List lists all EdgeIngresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error)
	// EdgeIngresses returns an object that can list and get EdgeIngresses.
	EdgeIngresses(namespace string) EdgeIngressNamespaceLister
	EdgeIngressListerExpansion
}

// edgeIngressLister implements the EdgeIngressLister interface.
type edgeIngressLister struct {
	indexer cache.Indexer
}

// NewEdgeIngressLister returns a new EdgeIngressLister.
func NewEdgeIngressLister(indexer cache.Indexer) EdgeIngressLister {
	return &edgeIngressLister{indexer: indexer}
}

// List lists all EdgeIngresses in the indexer.
func (s *edgeIngressLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeIngress))
	})
	return ret, err
}

// EdgeIngresses returns an object that can list and get EdgeIngresses.
func (s *edgeIngressLister) EdgeIngresses(namespace string) EdgeIngressNamespaceLister {
	return edgeIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EdgeIngressNamespaceLister helps list and get EdgeIngresses.
// All objects returned here must be treated as read-only.
type EdgeIngressNamespaceLister interface {
	// List lists all EdgeIngresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error)
	// Get retrieves the EdgeIngress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EdgeIngress, error)
	EdgeIngressNamespaceListerExpansion
}

// edgeIngressNamespaceLister implements the EdgeIngressNamespaceLister
// interface.
type edgeIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EdgeIngresses in the indexer for a given namespace.
func (s edgeIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeIngress))
	})
	return ret, err
}

// Get retrieves the EdgeIngress from the indexer for a given namespace and name.
func (s edgeIngressNamespaceLister) Get(name string) (*v1alpha1.EdgeIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("edgeingress"), name)
	}
	return obj.(*v1alpha1.EdgeIngress), nil
}
