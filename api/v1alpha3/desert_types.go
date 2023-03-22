/*
Copyright 2023.

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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DesertSpec defines the desired state of Desert
type DesertSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Desert. Edit desert_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	// https://www.worldatlas.com/articles/what-are-the-four-types-of-deserts.html
	// +kubebuilder:validation:Enum:=HotAndDry;Semiarid;Coastal;Cold
	Type string `json:"type"`

	// +kubebuilder:validation:Enum:=Arriving;Landed;Surviving;Critical
	// +kubebuilder:default:=Arriving
	Traveler string `json:"traveler,omitempty"`

	// Set length of stay.
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:default:=5
	Days int `json:"days,omitempty"`
}

// DesertStatus defines the observed state of Desert
type DesertStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Enum:=Arriving;Landed;Surviving;Critical
	Traveler string `json:"traveler,omitempty"`

	// +kubebuilder:validation:Minimum:=0
	WaterLevel int `json:"waterLevel"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// Desert is the Schema for the deserts API
type Desert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DesertSpec   `json:"spec,omitempty"`
	Status DesertStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DesertList contains a list of Desert
type DesertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Desert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Desert{}, &DesertList{})
}
