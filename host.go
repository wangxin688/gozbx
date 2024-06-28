package gozbx

type HostModel struct {

}

type HostGet struct {
	CommonGetParams
	HostIds        []string `json:"hostids,omitempty"`
	MaintenanceIds []string `json:"maintenanceids,omitempty"`
	TriggerIds     []string `json:"triggerids,omitempty"`
}


type HostCreate struct {
	Host string `json:"host"`
}