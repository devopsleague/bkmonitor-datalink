// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/apis/monitoring/v1beta1"
	scheme "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataIDsGetter has a method to return a DataIDInterface.
// A group's client should implement this interface.
type DataIDsGetter interface {
	DataIDs(namespace string) DataIDInterface
}

// DataIDInterface has methods to work with DataID resources.
type DataIDInterface interface {
	Create(ctx context.Context, dataID *v1beta1.DataID, opts v1.CreateOptions) (*v1beta1.DataID, error)
	Update(ctx context.Context, dataID *v1beta1.DataID, opts v1.UpdateOptions) (*v1beta1.DataID, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DataID, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DataIDList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataID, err error)
	DataIDExpansion
}

// dataIDs implements DataIDInterface
type dataIDs struct {
	client rest.Interface
	ns     string
}

// newDataIDs returns a DataIDs
func newDataIDs(c *MonitoringV1beta1Client, namespace string) *dataIDs {
	return &dataIDs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataID, and returns the corresponding dataID object, and an error if there is any.
func (c *dataIDs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataID, err error) {
	result = &v1beta1.DataID{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataids").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataIDs that match those selectors.
func (c *dataIDs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataIDList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DataIDList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dataids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataIDs.
func (c *dataIDs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dataids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dataID and creates it.  Returns the server's representation of the dataID, and an error, if there is any.
func (c *dataIDs) Create(ctx context.Context, dataID *v1beta1.DataID, opts v1.CreateOptions) (result *v1beta1.DataID, err error) {
	result = &v1beta1.DataID{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dataids").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataID).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dataID and updates it. Returns the server's representation of the dataID, and an error, if there is any.
func (c *dataIDs) Update(ctx context.Context, dataID *v1beta1.DataID, opts v1.UpdateOptions) (result *v1beta1.DataID, err error) {
	result = &v1beta1.DataID{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dataids").
		Name(dataID.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataID).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dataID and deletes it. Returns an error if one occurs.
func (c *dataIDs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataids").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataIDs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dataids").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dataID.
func (c *dataIDs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataID, err error) {
	result = &v1beta1.DataID{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dataids").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
