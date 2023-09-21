// apis/cronjob/v1alpha1/types.go
// +k8s:deepcopy-gen=package
// +groupName=jack.jack.operator.test

// Package v1alpha1 is the v1alpha1 version of the API.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CronJobSpec `json:"spec"`
}

type CronJobSpec struct {
	Foo string `json:"foo"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CronJobList is a list of Foo resources
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CronJob `json:"items"`
}
