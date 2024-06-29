package gozbx

type ActiveAvailableEnum uint8
type HostStatusEnum uint8
type MonitoredByEnum uint8


const (
	Unknown ActiveAvailableEnum = 0
	Available ActiveAvailableEnum = 1
	Unavailable ActiveAvailableEnum = 2
)

const (
	Monitered HostStatusEnum = 0
	Unmonitered HostStatusEnum = 1
)

const (
	ZabbixServer MonitoredByEnum = 0
	ZabbixProxy MonitoredByEnum = 1
	ZabbixProxyGroup MonitoredByEnum = 2
)



type HostModel struct {
	HostID    string        `json:"hostid,omitempty"`
	Host      string        `json:"host"`
	Description string        `json:"description,omitempty"`
	MonitoredBy []string      `json:"monitored_by,omitempty"`
	ProxyId    string        `json:"proxyid,omitempty"`
	ProxyGroupId string       `json:"proxy_groupid,omitempty"`
	Status     HostStatusEnum    `json:"status,omitempty"`
	ActiveAvailable ActiveAvailableEnum`json:"active_available,omitempty"`
	AssginedProxyID string `json:"assigned_proxyid,omitempty"` // ID of the proxy assigned by Zabbix server, if the host is monitored by a proxy group.
}

type HostGet struct {
	CommonGetParams
	HostIds        []string `json:"hostids,omitempty"`
	MaintenanceIds []string `json:"maintenanceids,omitempty"`
	TriggerIds     []string `json:"triggerids,omitempty"`
}


type HostCreate struct {
	Groups []GroupId `json:"groups"`
	Host string `json:"host"`
	Name string `json:"name,omitempty"`
	Interfaces []InterfaceCreate `json:"interfaces"`
	Tags []Tag `json:"tags,omitempty"`
	Templates []TemplatedId `json:"templates"`
	Macros []Macro `json:"macros,omitempty"`
	Inventory map[string]string `json:"inventory,omitempty"`
	TLSAccepted uint8 `json:"tls_accepted,omitempty"`
	TLSConnect  uint8 `json:"tls_connect,omitempty"`
	TLSPSKIdentity string `json:"tls_psk_identity,omitempty"`
	TSLPSK string `json:"tls_psk,omitempty"`
	MoniteredBy MonitoredByEnum `json:"monitored_by,omitempty"`
	ProxyId string `json:"proxyid,omitempty"`
	ProxyGroupId string `json:"proxy_groupid,omitempty"`
}
