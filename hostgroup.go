package gozbx

import (
	"github.com/google/uuid"
)

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

type HostGroupIdResponse struct {
	GroupIds []string `json:"groupids"`
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

func (hg *HosGroupImpl) Get(params HostGroupGet) (res []HostGroupModel, err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.get",
	}
	resp, err := hg.z.Rpc(request)
	if err != nil {
		return nil, err
	}
	resp.GetResult(&res)
	return

}

func (hg *HosGroupImpl) Create(params HostGroupCreate) (rsp string,err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.create",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return "", err
	}
	hgr := HostGroupIdResponse{}
	resp.GetResult(&hgr)
	return hgr.GroupIds[0], nil
}

func (hg *HosGroupImpl) Update(params HostGroupUpdate) (rsp string, err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.update",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return "nil", err
	}
	hgr := HostGroupIdResponse{}
	resp.GetResult(&hgr)
	return hgr.GroupIds[0], nil
}

func (hg *HosGroupImpl) Delete(hostgroupIds []string) (rsp []string , err error) {

	request := &Request{
		Params: hostgroupIds,
		Method: "hostgroup.delete",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return nil, err
	}
	resp.GetResult(&rsp)
	return
}

func (hg *HosGroupImpl) MassAdd(params HostGroupMassAdd) (rsp *HostGroupIdResponse, err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massadd",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return nil, err
	}
	resp.GetResult(&rsp)
	return
}

func (hg *HosGroupImpl) MassRemove(params HostGroupMassRemove) (rsp *HostGroupIdResponse, err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massremove",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return nil, err
	}
	resp.GetResult(&rsp)
	return
}

func (hg *HosGroupImpl) MassUpdate(params HostGroupMassUpdate) (rsp *HostGroupIdResponse, err error) {
	request := &Request{
		Params: params,
		Method: "hostgroup.massupdate",
	}
	resp, err:= hg.z.Rpc(request)
	if err != nil {
		return nil, err
	}
	resp.GetResult(&rsp)
	return
}
