// Code generated by nexus. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"perfdm/build/common"
)

// +k8s:openapi-gen=true
type Child struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type Link struct {
	Group string `json:"group" yaml:"group"`
	Kind  string `json:"kind" yaml:"kind"`
	Name  string `json:"name" yaml:"name"`
}

// +k8s:openapi-gen=true
type SyncerStatus struct {
	EtcdVersion    int64 `json:"etcdVersion, omitempty" yaml:"etcdVersion, omitempty"`
	CRGenerationId int64 `json:"cRGenerationId, omitempty" yaml:"cRGenerationId, omitempty"`
}

// +k8s:openapi-gen=true
type NexusStatus struct {
	SourceGeneration int64        `json:"sourceGeneration, omitempty" yaml:"sourceGeneration, omitempty"`
	RemoteGeneration int64        `json:"remoteGeneration, omitempty" yaml:"remoteGeneration, omitempty"`
	SyncerStatus     SyncerStatus `json:"syncerStatus, omitempty" yaml:"syncerStatus, omitempty"`
}

/* ------------------- CRDs definitions ------------------- */

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type SockShop struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              SockShopSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            SockShopNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type SockShopNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *SockShop) CRDName() string {
	return "sockshops.root.sockshop.com"
}

func (c *SockShop) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type SockShopSpec struct {
	OrgName           string           `json:"orgName" yaml:"orgName"`
	Location          string           `json:"location" yaml:"location"`
	Website           string           `json:"website" yaml:"website"`
	InventoryGvk      map[string]Child `json:"inventoryGvk,omitempty" yaml:"inventoryGvk,omitempty" nexus:"children"`
	POGvk             map[string]Child `json:"pOGvk,omitempty" yaml:"pOGvk,omitempty" nexus:"children"`
	ShippingLedgerGvk map[string]Child `json:"shippingLedgerGvk,omitempty" yaml:"shippingLedgerGvk,omitempty" nexus:"children"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SockShopList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []SockShop `json:"items" yaml:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Socks struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              SocksSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            SocksNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type SocksNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Socks) CRDName() string {
	return "sockses.root.sockshop.com"
}

func (c *Socks) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type SocksSpec struct {
	Brand string `json:"brand" yaml:"brand"`
	Color string `json:"color" yaml:"color"`
	Size  int    `json:"size" yaml:"size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SocksList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Socks `json:"items" yaml:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Orders struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              OrdersSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            OrdersNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type OrdersNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Orders) CRDName() string {
	return "orderses.root.sockshop.com"
}

func (c *Orders) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type OrdersSpec struct {
	SockName    string `json:"sockName" yaml:"sockName"`
	Address     string `json:"address" yaml:"address"`
	CartGvk     *Link  `json:"cartGvk,omitempty" yaml:"cartGvk,omitempty" nexus:"link"`
	ShippingGvk *Link  `json:"shippingGvk,omitempty" yaml:"shippingGvk,omitempty" nexus:"link"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type OrdersList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Orders `json:"items" yaml:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Shipping struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ShippingSpec        `json:"spec,omitempty" yaml:"spec,omitempty"`
	Status            ShippingNexusStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// +k8s:openapi-gen=true
type ShippingNexusStatus struct {
	Nexus NexusStatus `json:"nexus,omitempty" yaml:"nexus,omitempty"`
}

func (c *Shipping) CRDName() string {
	return "shippings.root.sockshop.com"
}

func (c *Shipping) DisplayName() string {
	if c.GetLabels() != nil {
		return c.GetLabels()[common.DISPLAY_NAME_LABEL]
	}
	return ""
}

// +k8s:openapi-gen=true
type ShippingSpec struct {
	TrackingId int `json:"trackingId" yaml:"trackingId"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ShippingList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Shipping `json:"items" yaml:"items"`
}