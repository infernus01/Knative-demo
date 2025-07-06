package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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

type NamespaceCleanerSpec struct {
	Schedule string `json:"schedule,omitempty"`

	Selector metav1.LabelSelector `json:"selector,omitempty"`
}
type NamespaceCleanerStatus struct {
	LastCleanup metav1.Time        `json:"lastCleanup,omitempty"`
	NextCleanup metav1.Time        `json:"nextCleanup,omitempty"`
	Conditions  []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
