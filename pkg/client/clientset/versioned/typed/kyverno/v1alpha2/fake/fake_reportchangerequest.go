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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "github.com/kyverno/kyverno/pkg/api/kyverno/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReportChangeRequests implements ReportChangeRequestInterface
type FakeReportChangeRequests struct {
	Fake *FakeKyvernoV1alpha2
	ns   string
}

var reportchangerequestsResource = schema.GroupVersionResource{Group: "kyverno.io", Version: "v1alpha2", Resource: "reportchangerequests"}

var reportchangerequestsKind = schema.GroupVersionKind{Group: "kyverno.io", Version: "v1alpha2", Kind: "ReportChangeRequest"}

// Get takes name of the reportChangeRequest, and returns the corresponding reportChangeRequest object, and an error if there is any.
func (c *FakeReportChangeRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportchangerequestsResource, c.ns, name), &v1alpha2.ReportChangeRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReportChangeRequest), err
}

// List takes label and field selectors, and returns the list of ReportChangeRequests that match those selectors.
func (c *FakeReportChangeRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ReportChangeRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportchangerequestsResource, reportchangerequestsKind, c.ns, opts), &v1alpha2.ReportChangeRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ReportChangeRequestList{ListMeta: obj.(*v1alpha2.ReportChangeRequestList).ListMeta}
	for _, item := range obj.(*v1alpha2.ReportChangeRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reportChangeRequests.
func (c *FakeReportChangeRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportchangerequestsResource, c.ns, opts))

}

// Create takes the representation of a reportChangeRequest and creates it.  Returns the server's representation of the reportChangeRequest, and an error, if there is any.
func (c *FakeReportChangeRequests) Create(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.CreateOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportchangerequestsResource, c.ns, reportChangeRequest), &v1alpha2.ReportChangeRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReportChangeRequest), err
}

// Update takes the representation of a reportChangeRequest and updates it. Returns the server's representation of the reportChangeRequest, and an error, if there is any.
func (c *FakeReportChangeRequests) Update(ctx context.Context, reportChangeRequest *v1alpha2.ReportChangeRequest, opts v1.UpdateOptions) (result *v1alpha2.ReportChangeRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportchangerequestsResource, c.ns, reportChangeRequest), &v1alpha2.ReportChangeRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReportChangeRequest), err
}

// Delete takes name of the reportChangeRequest and deletes it. Returns an error if one occurs.
func (c *FakeReportChangeRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportchangerequestsResource, c.ns, name), &v1alpha2.ReportChangeRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReportChangeRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportchangerequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ReportChangeRequestList{})
	return err
}

// Patch applies the patch and returns the patched reportChangeRequest.
func (c *FakeReportChangeRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ReportChangeRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportchangerequestsResource, c.ns, name, pt, data, subresources...), &v1alpha2.ReportChangeRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ReportChangeRequest), err
}
