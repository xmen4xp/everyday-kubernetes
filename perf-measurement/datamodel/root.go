package root

import (
	"github.com/vmware-tanzu/graph-framework-for-microservices/nexus/nexus"
)

type SockShop struct {
	nexus.SingletonNode

	OrgName  string
	Location string
	Website  string

	Inventory      Socks    `nexus:"children"`
	PO             Orders   `nexus:"children"`
	ShippingLedger Shipping `nexus:"children"`
}

var SocksRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/sock/{root.Socks}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
		{
			Uri:     "/socks",
			Methods: nexus.HTTPListResponse,
		},
	},
}

// nexus-rest-api-gen:SocksRestAPISpec
type Socks struct {
	nexus.Node

	Brand string
	Color string
	Size  int
}

var OrderRestAPISpec = nexus.RestAPISpec{
	Uris: []nexus.RestURIs{
		{
			Uri:     "/order/{root.Orders}",
			Methods: nexus.DefaultHTTPMethodsResponses,
		},
	},
}

// nexus-rest-api-gen:OrderRestAPISpec
type Orders struct {
	nexus.Node

	SockName string
	Address  string

	Cart     Socks    `nexus:"link"`
	Shipping Shipping `nexus:"link"`
}

type Shipping struct {
	nexus.Node

	TrackingId int
}
