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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

// PodSchedulingContextSpecApplyConfiguration represents an declarative configuration of the PodSchedulingContextSpec type for use
// with apply.
type PodSchedulingContextSpecApplyConfiguration struct {
	SelectedNode   *string  `json:"selectedNode,omitempty"`
	PotentialNodes []string `json:"potentialNodes,omitempty"`
}

// PodSchedulingContextSpecApplyConfiguration constructs an declarative configuration of the PodSchedulingContextSpec type for use with
// apply.
func PodSchedulingContextSpec() *PodSchedulingContextSpecApplyConfiguration {
	return &PodSchedulingContextSpecApplyConfiguration{}
}

// WithSelectedNode sets the SelectedNode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelectedNode field is set to the value of the last call.
func (b *PodSchedulingContextSpecApplyConfiguration) WithSelectedNode(value string) *PodSchedulingContextSpecApplyConfiguration {
	b.SelectedNode = &value
	return b
}

// WithPotentialNodes adds the given value to the PotentialNodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PotentialNodes field.
func (b *PodSchedulingContextSpecApplyConfiguration) WithPotentialNodes(values ...string) *PodSchedulingContextSpecApplyConfiguration {
	for i := range values {
		b.PotentialNodes = append(b.PotentialNodes, values[i])
	}
	return b
}
