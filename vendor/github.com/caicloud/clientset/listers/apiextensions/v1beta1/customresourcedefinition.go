/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/caicloud/clientset/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionLister helps list CustomResourceDefinitions.
type CustomResourceDefinitionLister interface {
	// List lists all CustomResourceDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error)
	// Get retrieves the CustomResourceDefinition from the index for a given name.
	Get(name string) (*v1beta1.CustomResourceDefinition, error)
	CustomResourceDefinitionListerExpansion
}

// customResourceDefinitionLister implements the CustomResourceDefinitionLister interface.
type customResourceDefinitionLister struct {
	indexer cache.Indexer
}

// NewCustomResourceDefinitionLister returns a new CustomResourceDefinitionLister.
func NewCustomResourceDefinitionLister(indexer cache.Indexer) CustomResourceDefinitionLister {
	return &customResourceDefinitionLister{indexer: indexer}
}

// List lists all CustomResourceDefinitions in the indexer.
func (s *customResourceDefinitionLister) List(selector labels.Selector) (ret []*v1beta1.CustomResourceDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CustomResourceDefinition))
	})
	return ret, err
}

// Get retrieves the CustomResourceDefinition from the index for a given name.
func (s *customResourceDefinitionLister) Get(name string) (*v1beta1.CustomResourceDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("customresourcedefinition"), name)
	}
	return obj.(*v1beta1.CustomResourceDefinition), nil
}
