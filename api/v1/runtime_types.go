/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RuntimeConfigMap bulabula
type RuntimeConfigMap struct {
	Name  string `json:"name,omitempty"`
	Mount string `json:"mount,omitempty"`
}

// RuntimeSpec defines the desired state of Runtime
type RuntimeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The container image of runtime
	// +kubebuilder:validation:Required
	Image string `json:"image,omitempty"`

	// The container image of runtime
	// +kubebuilder:validation:Optional
	Command []string `json:"command,omitempty"`

	// Optional port for runtime.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=8000
	Port int32 `json:"port,omitempty"`

	// Optional port for runtime.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="http"
	PortName string `json:"portName,omitempty"`

	// Optional cluster IP spec of runtime
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=1
	Replicas *int32 `json:"replicas,omitempty"`

	// Optional cluster IP spec of runtime
	// +kubebuilder:validation:Optional
	ClusterIP string `json:"clusterIP,omitempty"`

	// Optional ready format spec of runtime
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="{AvailableReplicas}/{AvailableReplicas}"
	ReadyFormat string `json:"readyFormat,omitempty"`
}

// RuntimeStatus defines the observed state of Runtime
type RuntimeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Optional functions config maps of runtime
	// +kubebuilder:validation:Optional
	Functions map[string]RuntimeConfigMap `json:"functions,omitemty"`

	// Optional libraries config maps of runtime
	// +kubebuilder:validation:Optional
	Libraries map[string]RuntimeConfigMap `json:"libraries,omitempty"`

	// Optional ready string of runtime for show
	// +kubebuilder:validation:Optional
	Ready string `json:"ready,omitempty"`
}

// +kubebuilder:resource:categories="kess",shortName="rt"
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.selector
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.ready`,priority=0
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`,priority=0
// +kubebuilder:printcolumn:name="Command",type=string,JSONPath=`.spec.command`,priority=10
// +kubebuilder:object:root=true

// Runtime is the Schema for the runtimes API
type Runtime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RuntimeSpec   `json:"spec,omitempty"`
	Status RuntimeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuntimeList contains a list of Runtime
type RuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runtime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Runtime{}, &RuntimeList{})
}
