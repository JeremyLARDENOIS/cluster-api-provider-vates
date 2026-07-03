package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// XOMachineTemplateResource describes the data needed to create a XOMachine from a template.
type XOMachineTemplateResource struct {
	// Spec is the specification of the desired behavior of the machine.
	Spec XOMachineSpec `json:"spec"`
}

// XOMachineTemplateSpec defines the desired state of XOMachineTemplate.
type XOMachineTemplateSpec struct {
	Template XOMachineTemplateResource `json:"template"`
}

// XOMachineTemplateStatus defines the observed state of XOMachineTemplate.
type XOMachineTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=xomachinetemplates,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion

// XOMachineTemplate is the Schema for the xomachinetemplates API.
type XOMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitzero"`

	Spec   XOMachineTemplateSpec   `json:"spec,omitempty"`
	Status XOMachineTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XOMachineTemplateList contains a list of XOMachineTemplate.
type XOMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []XOMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&XOMachineTemplate{}, &XOMachineTemplateList{})
}
