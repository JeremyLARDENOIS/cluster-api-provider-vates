package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// XOClusterTemplateResource describes the data needed to create a XOCluster from a template.
type XOClusterTemplateResource struct {
	// Spec is the specification of the desired behavior of the cluster.
	Spec XOClusterSpec `json:"spec"`
}

// XOClusterTemplateSpec defines the desired state of XOClusterTemplate.
type XOClusterTemplateSpec struct {
	Template XOClusterTemplateResource `json:"template"`
}

// XOClusterTemplateStatus defines the observed state of XOClusterTemplate.
type XOClusterTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=xoclustertemplates,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion

// XOClusterTemplate is the Schema for the xoclustertemplates API.
type XOClusterTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitzero"`

	Spec   XOClusterTemplateSpec   `json:"spec,omitempty"`
	Status XOClusterTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XOClusterTemplateList contains a list of XOClusterTemplate.
type XOClusterTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []XOClusterTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&XOClusterTemplate{}, &XOClusterTemplateList{})
}
