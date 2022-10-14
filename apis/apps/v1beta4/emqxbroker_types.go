/*
Copyright 2021.

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

package v1beta4

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EmqxBrokerSpec defines the desired state of EmqxBroker
type EmqxBrokerSpec struct {
	//+kubebuilder:default:=3
	Replicas *int32 `json:"replicas,omitempty"`

	// VolumeClaimTemplates describes the common attributes of storage devices
	VolumeClaimTemplates []corev1.PersistentVolumeClaim `json:"volumeClaimTemplates,omitempty"`

	Template EmqxTemplate `json:"template,omitempty"`

	// ServiceTemplate defines a logical set of ports and a policy by which to access them
	ServiceTemplate ServiceTemplate `json:"serviceTemplate,omitempty"`
}

// EmqxBrokerStatus defines the observed state of EmqxBroker
type EmqxBrokerStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas
//+kubebuilder:storageversion

// EmqxBroker is the Schema for the emqxbrokers API
type EmqxBroker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmqxBrokerSpec `json:"spec,omitempty"`
	Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EmqxBrokerList contains a list of EmqxBroker
type EmqxBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmqxBroker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmqxBroker{}, &EmqxBrokerList{})
}

func (emqx *EmqxBroker) GetReplicas() *int32 {
	return emqx.Spec.Replicas
}

func (emqx *EmqxBroker) SetReplicas(replicas int32) {
	emqx.Spec.Replicas = &replicas
}

func (emqx *EmqxBroker) GetVolumeClaimTemplates() []corev1.PersistentVolumeClaim {
	return emqx.Spec.VolumeClaimTemplates
}

func (emqx *EmqxBroker) SetVolumeClaimTemplates(volumeClaimTemplates []corev1.PersistentVolumeClaim) {
	emqx.Spec.VolumeClaimTemplates = volumeClaimTemplates
}

func (emqx *EmqxBroker) GetTemplate() EmqxTemplate {
	return emqx.Spec.Template
}

func (emqx *EmqxBroker) SetTemplate(template EmqxTemplate) {
	emqx.Spec.Template = template
}

func (emqx *EmqxBroker) GetServiceTemplate() ServiceTemplate {
	return emqx.Spec.ServiceTemplate
}
func (emqx *EmqxBroker) SetServiceTemplate(serviceTemplate ServiceTemplate) {
	emqx.Spec.ServiceTemplate = serviceTemplate
}

func (emqx *EmqxBroker) GetStatus() Status {
	return emqx.Status
}
func (emqx *EmqxBroker) SetStatus(status Status) {
	emqx.Status = status
}
