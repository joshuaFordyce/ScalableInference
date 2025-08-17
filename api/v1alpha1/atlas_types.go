/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k9s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AtlasSpec defines the desired state of Atlas



type AtlasSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// foo is an example field of Atlas. Edit atlas_types.go to remove/update
	// +optional
	Deployment appsv1.DeploymentSpec `json:"deployment,omitempty"`
	Service corev1.ServiceSpec `json:"service,omitempty"`
	HPA autoscalingv2.HorizontalPodAutoscalerSpec `json:"HPA,omitempty"`

}

// AtlasStatus defines the observed state of Atlas.
type AtlasStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	ReadyReplicas int32 `json:"readyreplicas, omitempty"`

	HPAStatus string `json:"hpaStatus,omitempty"`
	ServiceStatus string `json:"serviceStatus,omitempty`

}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Atlas is the Schema for the atlas API
type Atlas struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of Atlas
	// +required
	Spec AtlasSpec `json:"spec"`

	// status defines the observed state of Atlas
	// +optional
	Status AtlasStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// AtlasList contains a list of Atlas
type AtlasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Atlas `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Atlas{}, &AtlasList{})
}
