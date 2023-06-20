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

	v1alpha1 "github.com/linkerd/linkerd2/controller/gen/apis/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthorizationPolicies implements AuthorizationPolicyInterface
type FakeAuthorizationPolicies struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var authorizationpoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("authorizationpolicies")

var authorizationpoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("AuthorizationPolicy")

// Get takes name of the authorizationPolicy, and returns the corresponding authorizationPolicy object, and an error if there is any.
func (c *FakeAuthorizationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authorizationpoliciesResource, c.ns, name), &v1alpha1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of AuthorizationPolicies that match those selectors.
func (c *FakeAuthorizationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authorizationpoliciesResource, authorizationpoliciesKind, c.ns, opts), &v1alpha1.AuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuthorizationPolicyList{ListMeta: obj.(*v1alpha1.AuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authorizationPolicies.
func (c *FakeAuthorizationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a authorizationPolicy and creates it.  Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Create(ctx context.Context, authorizationPolicy *v1alpha1.AuthorizationPolicy, opts v1.CreateOptions) (result *v1alpha1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1alpha1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationPolicy), err
}

// Update takes the representation of a authorizationPolicy and updates it. Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Update(ctx context.Context, authorizationPolicy *v1alpha1.AuthorizationPolicy, opts v1.UpdateOptions) (result *v1alpha1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1alpha1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationPolicy), err
}

// Delete takes name of the authorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAuthorizationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(authorizationpoliciesResource, c.ns, name, opts), &v1alpha1.AuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthorizationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authorizationpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched authorizationPolicy.
func (c *FakeAuthorizationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuthorizationPolicy), err
}
