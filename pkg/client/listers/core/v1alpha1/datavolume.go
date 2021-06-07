/*
Copyright 2018 The CDI Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

// DataVolumeLister helps list DataVolumes.
// All objects returned here must be treated as read-only.
type DataVolumeLister interface {
	// List lists all DataVolumes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataVolume, err error)
	// DataVolumes returns an object that can list and get DataVolumes.
	DataVolumes(namespace string) DataVolumeNamespaceLister
	DataVolumeListerExpansion
}

// dataVolumeLister implements the DataVolumeLister interface.
type dataVolumeLister struct {
	indexer cache.Indexer
}

// NewDataVolumeLister returns a new DataVolumeLister.
func NewDataVolumeLister(indexer cache.Indexer) DataVolumeLister {
	return &dataVolumeLister{indexer: indexer}
}

// List lists all DataVolumes in the indexer.
func (s *dataVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.DataVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataVolume))
	})
	return ret, err
}

// DataVolumes returns an object that can list and get DataVolumes.
func (s *dataVolumeLister) DataVolumes(namespace string) DataVolumeNamespaceLister {
	return dataVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataVolumeNamespaceLister helps list and get DataVolumes.
// All objects returned here must be treated as read-only.
type DataVolumeNamespaceLister interface {
	// List lists all DataVolumes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataVolume, err error)
	// Get retrieves the DataVolume from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataVolume, error)
	DataVolumeNamespaceListerExpansion
}

// dataVolumeNamespaceLister implements the DataVolumeNamespaceLister
// interface.
type dataVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataVolumes in the indexer for a given namespace.
func (s dataVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataVolume))
	})
	return ret, err
}

// Get retrieves the DataVolume from the indexer for a given namespace and name.
func (s dataVolumeNamespaceLister) Get(name string) (*v1alpha1.DataVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datavolume"), name)
	}
	return obj.(*v1alpha1.DataVolume), nil
}
