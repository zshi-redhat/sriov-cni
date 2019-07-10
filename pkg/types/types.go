package types

import (
	"github.com/containernetworking/cni/pkg/types"
	"net"
)

// NetConf extends types.NetConf for sriov-cni
type NetConf struct {
	types.NetConf
	DPDKMode    bool
	Master      string
	MAC         string
	OriginalMAC string `json:"originalMAC"`
	Vlan        int    `json:"vlan"`
	DeviceID    string `json:"deviceID"` // PCI address of a VF in valid sysfs format
	VFID        int
	HostIFNames string           // VF netdevice name(s)
	ContIFNames string           // VF names after in the container; used during deletion
	MaxTxRate   *int             `json:"max_tx_rate"`         // Mbps, 0 = disable rate limiting
	SpoofChk    string           `json:"spoofchk,omitempty"`  // on|off
	Trust       string           `json:"trust,omitempty"`     // on|off
	NodeGUID    net.HardwareAddr `json:"node_guid,omitempty"` // node GUID for Infiniband
	PortGUID    net.HardwareAddr `json:"port_guid,omitempty"` // port GUID for Infiniband
	/* The following 3 flags are not yet supported in netlink lib.
	MinTxRate   *int   `json:"min_tx_rate"`         // Mbps, 0 = disable rate limiting
	QueryRss    string `json:"query_rss,omitempty"` // on|off
	State       string `json:"state,omitempty"`     // auto, enable, disable */
}
