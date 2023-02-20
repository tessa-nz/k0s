/*
Copyright k0s authors

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

package v1beta2

import (
	"context"

	v1beta2 "github.com/k0sproject/k0s/pkg/apis/autopilot.k0sproject.io/v1beta2"
	scheme "github.com/k0sproject/k0s/pkg/apis/autopilot.k0sproject.io/v1beta2/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// ControlNodeListsGetter has a method to return a ControlNodeListInterface.
// A group's client should implement this interface.
type ControlNodeListsGetter interface {
	ControlNodeLists() ControlNodeListInterface
}

// ControlNodeListInterface has methods to work with ControlNodeList resources.
type ControlNodeListInterface interface {
	Create(ctx context.Context, controlNodeList *v1beta2.ControlNodeList, opts v1.CreateOptions) (*v1beta2.ControlNodeList, error)
	ControlNodeListExpansion
}

// controlNodeLists implements ControlNodeListInterface
type controlNodeLists struct {
	client rest.Interface
}

// newControlNodeLists returns a ControlNodeLists
func newControlNodeLists(c *AutopilotV1beta2Client) *controlNodeLists {
	return &controlNodeLists{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a controlNodeList and creates it.  Returns the server's representation of the controlNodeList, and an error, if there is any.
func (c *controlNodeLists) Create(ctx context.Context, controlNodeList *v1beta2.ControlNodeList, opts v1.CreateOptions) (result *v1beta2.ControlNodeList, err error) {
	result = &v1beta2.ControlNodeList{}
	err = c.client.Post().
		Resource("controlnodelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(controlNodeList).
		Do(ctx).
		Into(result)
	return
}
