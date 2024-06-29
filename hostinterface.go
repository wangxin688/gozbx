package gozbx 

type HostInterfaceEnum uint8

type HostInterfaceAvailableEnum uint8

type HostIpEnum uint8

const (
	AgentHostInterfaceType HostInterfaceEnum = 1
	SNMPHostInterfaceType HostInterfaceEnum= 2
	IPMIHostInterfaceType HostInterfaceEnum = 3
	JMXHostInterfaceType HostInterfaceEnum = 4
)

const (
	InterfaceUnknownType HostInterfaceAvailableEnum = 0
	InterfaceAvailableType HostInterfaceAvailableEnum = 1
	InterfaceUnAvailableType HostInterfaceAvailableEnum = 2
)

const (
	UseDNSType HostIpEnum = 0
	UseIPType HostIpEnum = 1
)


type InterfaceCreate struct {
	Type HostInterfaceEnum `json:"type"`
	Main bool `json:"main,omitempty"`
	UserIP HostIpEnum `json:"useip,omitempty"`
	DNS string `json:"dns,omitempty"`
	Port uint `json:"port,omitempty"`
	Details map[string]string `json:"details,omitempty"`

}