// +k8s:deepcopy-gen=package
// +groupName=clusterops.io

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceCleaner is a custom resource for cleaning up namespaces based on schedule and selector
type NamespaceCleaner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceCleanerSpec   `json:"spec,omitempty"`
	Status NamespaceCleanerStatus `json:"status,omitempty"`
}

// DeepCopyObject implements runtime.Object.
func (n *NamespaceCleaner) DeepCopyObject() runtime.Object {
	panic("unimplemented")
}

// GetObjectKind implements runtime.Object.
// Subtle: this method shadows the method (TypeMeta).GetObjectKind of NamespaceCleaner.TypeMeta.
func (n *NamespaceCleaner) GetObjectKind() schema.ObjectKind {
	panic("unimplemented")
}

// NamespaceCleanerSpec defines the desired state of NamespaceCleaner
type NamespaceCleanerSpec struct {
	RetentionDays int `json:"retentionDays"`
}

// NamespaceCleanerStatus defines the observed state of NamespaceCleaner
type NamespaceCleanerStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceCleanerList contains a list of NamespaceCleaner
type NamespaceCleanerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceCleaner `json:"items"`
}

// DeepCopyObject implements runtime.Object.
func (n *NamespaceCleanerList) DeepCopyObject() runtime.Object {
	panic("unimplemented")
}

// GetObjectKind implements runtime.Object.
// Subtle: this method shadows the method (TypeMeta).GetObjectKind of NamespaceCleanerList.TypeMeta.
func (n *NamespaceCleanerList) GetObjectKind() schema.ObjectKind {
	panic("unimplemented")
}
