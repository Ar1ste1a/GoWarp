package warp

import "strings"

type UpdateType uint64

const (
	NEW_API_KEY UpdateType = iota
	CURRENT_USER
	ACTIVE_MACHINE
	MACHINES_LIST
	RETIRED_MACHINES_LIST
	ACTIVITY_LIST
	VPN_LIST
	VPN_STATUS
	BADGES_LIST
	PROLABS_PROGRESS
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
	Machine Machine
}

func (am ActiveMachine) GetUpdate() map[string]string {
	return map[string]string{"Machine": am.Machine.GetMapString()}
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

type ActivityList struct {
	Activity []UserActivity
}

func (al ActivityList) GetUpdate() map[string]string {
	var activityOut []string
	for _, activity := range al.Activity {
		activityOut = append(activityOut, activity.GetMapString())
	}
	activityList := strings.Join(activityOut, ",")
	return map[string]string{"Activity": activityList}
}

func (al ActivityList) GetType() UpdateType {
	return ACTIVITY_LIST
}

type BadgesList struct {
	Badges []Badge
}

func (bl BadgesList) GetUpdate() map[string]string {
	var badgesOut []string
	for _, badge := range bl.Badges {
		badgesOut = append(badgesOut, badge.GetMapString())
	}
	badgesList := strings.Join(badgesOut, ",")
	return map[string]string{"Badges": badgesList}
}

func (bl BadgesList) GetType() UpdateType {
	return BADGES_LIST
}

type ProlabsProgress struct {
	Progress []ProLabProgress
}

func (pp ProlabsProgress) GetUpdate() map[string]string {
	var progressOut []string
	for _, progress := range pp.Progress {
		progressOut = append(progressOut, progress.GetMapString())
	}
	progressList := strings.Join(progressOut, ",")
	return map[string]string{"Progress": progressList}
}

func (pp ProlabsProgress) GetType() UpdateType {
	return PROLABS_PROGRESS
}

type RetiredMachinesList struct {
	Machines []RetiredMachine
}

func (rml RetiredMachinesList) GetUpdate() map[string]string {
	var machinesOut []string
	for _, machine := range rml.Machines {
		machinesOut = append(machinesOut, machine.GetMapString())
	}
	machinesList := strings.Join(machinesOut, ",")
	return map[string]string{"Machines": machinesList}
}

func (rml RetiredMachinesList) GetType() UpdateType {
	return RETIRED_MACHINES_LIST
}
