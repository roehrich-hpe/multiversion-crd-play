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

// VehicleSpec defines the desired state of Vehicle
type VehicleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Vehicle. Edit vehicle_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	// Make requests a vehicle
	// +kubebuilder:validation:Enum:=LandRover;Jeep;Toyota
	Make string `json:"make"`
}

// VehicleStatus defines the observed state of Vehicle
type VehicleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Make is the vehicle being used
	// +kubebuilder:validation:Enum:=LandRover;Jeep;Toyota
	Make string `json:"make"`

	// Tires indicates the condition of the tires
	// +kubebuilder:validation:Enum:=New;Worn
	Tires string `json:"tires"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// Vehicle is the Schema for the vehicles API
type Vehicle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VehicleSpec   `json:"spec,omitempty"`
	Status VehicleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VehicleList contains a list of Vehicle
type VehicleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vehicle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Vehicle{}, &VehicleList{})
}
