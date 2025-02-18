/*


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
)

// RouteMonitorSpec defines the desired state of RouteMonitor
type RouteMonitorSpec struct {
	Route RouteMonitorRouteSpec `json:"route,omitempty"`
	Slo   SloSpec               `json:"slo,omitempty"`

	// +kubebuilder:default:false
	// +kubebuilder:validation:Optional

	// SkipPrometheusRule instructs the controller to skip the creation of PrometheusRule CRs.
	// One common use-case for is for alerts that are defined separately, such as for hosted clusters.
	SkipPrometheusRule bool `json:"skipPrometheusRule"`

	// +kubebuilder:default:false
	// +kubebuilder:validation:Optional

	// InsecureSkipTLSVerify indicates that the blackbox exporter module used to probe this route
	// should *not* use https
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify"`
}

// RouteMonitorRouteSpec references the observed Route resource
type RouteMonitorRouteSpec struct {
	// Name is the name of the Route
	Name string `json:"name,omitempty"`
	// Namespace is the namespace of the Route
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum:=1

	// Port optionally defines the port we should use while probing
	Port int64 `json:"port,omitempty"`

	// +kubebuilder:validation:Optional

	// Suffix optionally defines the path we should probe (/livez /readyz etc)
	Suffix string `json:"suffix,omitempty"`
}

// RouteMonitorStatus defines the observed state of RouteMonitor
type RouteMonitorStatus struct {
	// RouteURL is the url extracted from the Route resource
	RouteURL          string         `json:"routeURL,omitempty"`
	ServiceMonitorRef NamespacedName `json:"serviceMonitorRef,omitempty"`
	PrometheusRuleRef NamespacedName `json:"prometheusRuleRef,omitempty"`
	ErrorStatus       string         `json:"errorStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RouteMonitor is the Schema for the routemonitors API
type RouteMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteMonitorSpec   `json:"spec,omitempty"`
	Status RouteMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteMonitorList contains a list of RouteMonitor
type RouteMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RouteMonitor{}, &RouteMonitorList{})
}
