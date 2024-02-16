package warp

import "strings"

type UpdateType uint64

const (
	NEW_API_KEY UpdateType = iota
	CURRENT_USER
	ACTIVE_MACHINE
	MACHINES_LIST
	ACTIVITY_LIST
	VPN_LIST
	VPN_STATUS
)

type Update interface {
	GetUpdate() map[string]string
	GetType() UpdateType
}

type NewApiKey struct {
	ApiKey string
}

func (nak NewApiKey) GetUpdate() map[string]string {
	return map[string]string{"ApiKey": nak.ApiKey}
}

func (nak NewApiKey) GetType() UpdateType {
	return NEW_API_KEY
}

type CurrentUser struct {
	Username string
	User     UserBrief
}

func (cu CurrentUser) GetUpdate() map[string]string {
	return map[string]string{"Username": cu.Username, "User": cu.User.GetMapString()}
}

func (cu CurrentUser) GetType() UpdateType {
	return CURRENT_USER
}

type ActiveMachine struct {
	Name string
	IP   string
}

func (am ActiveMachine) GetUpdate() map[string]string {
	return map[string]string{"Name": am.Name, "IP": am.IP}
}

func (am ActiveMachine) GetType() UpdateType {
	return ACTIVE_MACHINE
}

type MachinesList struct {
	Machines []Machine
}

func (ml MachinesList) GetUpdate() map[string]string {
	var machinesOut []string
	for _, machine := range ml.Machines {
		machinesOut = append(machinesOut, machine.GetMapString())
	}
	machinesList := strings.Join(machinesOut, ",")
	return map[string]string{"Machines": machinesList}
}

func (ml MachinesList) GetType() UpdateType {
	return MACHINES_LIST
}
