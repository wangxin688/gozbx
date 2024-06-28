package gozbx

import "github.com/google/uuid"

type HostGroupModel struct {
	GroupId string `json:"groupid"`
	Name    string `json:"name"`
	Flags   int    `json:"flags"` // 0: a plain hostgroup 4: discovered hostgroup
	Uuid    string `json:"uuid"`
}

type HostGroupGet struct {
	CommonGetParams
	GroupIds       []string `json:"groupids,omitempty"`       // Return only host groups with the given host group IDs.
	HostIds        []string `json:"hostids,omitempty"`        // Return only host groups that contain the given hosts.
	MaintenanceIds []string `json:"maintenanceids,omitempty"` // Return only host groups with the given maintenance IDs.
	TriggerIds     []string `json:"triggerids,omitempty"`     // Return only host groups that contain the given triggers.
}

type HostGroupCreate struct {
	Host string    `json:"name"`
	Uuid uuid.UUID `json:"uuid,omitempty"`
}

type HostGroupUpdate struct {
	GroupId string `json:"groupid"`
	Name    string `json:"name,omitempty"`
	Flags   int    `json:"flags,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

// This method allows to simultaneously add multiple related objects to all the given host groups.
type HostGroupMassAdd struct {
	Groups []GroupId `json:"groups"`
	Host   []HostId  `json:"hosts"`
}

// This method allows to remove related objects from multiple host groups.
type HostGroupMassRemove struct {
	HostGroupMassAdd
}

// This method allows to replace hosts and templates with the specified ones in multiple host groups.
type HostGroupMassUpdate struct {
	HostGroupMassAdd
}

// type HostGroupPropagate struct {
// 	groups []GroupId `json:"groups"`
// 	permission bool `json:"permission,omitempty"`
// 	tagFilters []TagFilter `json:"tag_filters,omitempty"`
// }

type HosGroupImpl struct {
	z *ZbxAPI
}

func (hg *HosGroupImpl) Get(params HostGroupGet) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.get",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) Create(params HostGroupCreate) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.create",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) Update(params HostGroupUpdate) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.update",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) Delete(hostgroupIds []string) (*Response, error) {

	request := &Request{
		Params: hostgroupIds,
		Method: "hostgroup.delete",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) MassAdd(params HostGroupMassAdd) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massadd",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) MassRemove(params HostGroupMassRemove) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massremove",
	}
	return hg.z.rpc(request)
}

func (hg *HosGroupImpl) MassUpdate(params HostGroupMassUpdate) (*Response, error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massupdate",
	}
	return hg.z.rpc(request)
}
