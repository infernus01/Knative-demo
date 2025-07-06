package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// NamespaceCleanerSpec defines the desired state of NamespaceCleaner
type NamespaceCleanerSpec struct {
	// Schedule defines when to run the cleanup (cron format)
	Schedule string `json:"schedule,omitempty"`

	// Selector defines which namespaces to clean up
	Selector metav1.LabelSelector `json:"selector,omitempty"`

	// MaxAge defines how old namespaces should be before cleanup (e.g., "7d", "24h", "30m")
	MaxAge string `json:"maxAge,omitempty"`
}

// NamespaceCleanerStatus defines the observed state of NamespaceCleaner
type NamespaceCleanerStatus struct {
	// LastCleanup is the timestamp of the last cleanup
	LastCleanup metav1.Time `json:"lastCleanup,omitempty"`

	// NextCleanup is the timestamp of the next scheduled cleanup
	NextCleanup metav1.Time `json:"nextCleanup,omitempty"`

	// Conditions represent the current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceCleanerList contains a list of NamespaceCleaner
type NamespaceCleanerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceCleaner `json:"items"`
}
