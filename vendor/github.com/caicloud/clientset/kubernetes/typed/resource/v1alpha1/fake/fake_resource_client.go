/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeResourceV1alpha1 struct {
	*testing.Fake
}

func (c *FakeResourceV1alpha1) StorageServices() v1alpha1.StorageServiceInterface {
	return &FakeStorageServices{c}
}

func (c *FakeResourceV1alpha1) StorageTypes() v1alpha1.StorageTypeInterface {
	return &FakeStorageTypes{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeResourceV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
