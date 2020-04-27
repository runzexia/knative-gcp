/*
Copyright 2020 Google LLC

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
	"time"

	v1alpha1 "github.com/google/knative-gcp/pkg/apis/intevents/v1alpha1"
	scheme "github.com/google/knative-gcp/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BrokerCellsGetter has a method to return a BrokerCellInterface.
// A group's client should implement this interface.
type BrokerCellsGetter interface {
	BrokerCells(namespace string) BrokerCellInterface
}

// BrokerCellInterface has methods to work with BrokerCell resources.
type BrokerCellInterface interface {
	Create(*v1alpha1.BrokerCell) (*v1alpha1.BrokerCell, error)
	Update(*v1alpha1.BrokerCell) (*v1alpha1.BrokerCell, error)
	UpdateStatus(*v1alpha1.BrokerCell) (*v1alpha1.BrokerCell, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BrokerCell, error)
	List(opts v1.ListOptions) (*v1alpha1.BrokerCellList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BrokerCell, err error)
	BrokerCellExpansion
}

// brokerCells implements BrokerCellInterface
type brokerCells struct {
	client rest.Interface
	ns     string
}

// newBrokerCells returns a BrokerCells
func newBrokerCells(c *InternalV1alpha1Client, namespace string) *brokerCells {
	return &brokerCells{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the brokerCell, and returns the corresponding brokerCell object, and an error if there is any.
func (c *brokerCells) Get(name string, options v1.GetOptions) (result *v1alpha1.BrokerCell, err error) {
	result = &v1alpha1.BrokerCell{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("brokercells").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BrokerCells that match those selectors.
func (c *brokerCells) List(opts v1.ListOptions) (result *v1alpha1.BrokerCellList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BrokerCellList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("brokercells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested brokerCells.
func (c *brokerCells) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("brokercells").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a brokerCell and creates it.  Returns the server's representation of the brokerCell, and an error, if there is any.
func (c *brokerCells) Create(brokerCell *v1alpha1.BrokerCell) (result *v1alpha1.BrokerCell, err error) {
	result = &v1alpha1.BrokerCell{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("brokercells").
		Body(brokerCell).
		Do().
		Into(result)
	return
}

// Update takes the representation of a brokerCell and updates it. Returns the server's representation of the brokerCell, and an error, if there is any.
func (c *brokerCells) Update(brokerCell *v1alpha1.BrokerCell) (result *v1alpha1.BrokerCell, err error) {
	result = &v1alpha1.BrokerCell{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("brokercells").
		Name(brokerCell.Name).
		Body(brokerCell).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *brokerCells) UpdateStatus(brokerCell *v1alpha1.BrokerCell) (result *v1alpha1.BrokerCell, err error) {
	result = &v1alpha1.BrokerCell{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("brokercells").
		Name(brokerCell.Name).
		SubResource("status").
		Body(brokerCell).
		Do().
		Into(result)
	return
}

// Delete takes name of the brokerCell and deletes it. Returns an error if one occurs.
func (c *brokerCells) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("brokercells").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *brokerCells) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("brokercells").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched brokerCell.
func (c *brokerCells) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BrokerCell, err error) {
	result = &v1alpha1.BrokerCell{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("brokercells").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}