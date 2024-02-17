package main

import (
	"GoWarp/warp"
	"fmt"
)

func main() {
	wc, _ := warp.GetWarpClient()
	//userInfo, err := wc.GetUserInfo()
	//if err != nil {
	//	fmt.Printf("User Info Error: %v\n", err)
	//}
	//ua, err := wc.GetUserActivity()
	//if err != nil {
	//	fmt.Printf("User Activity Error: %v\n", err)
	//}
	//uqr, err := wc.GetUserProfileOverview(1313425)

	//badges, err := wc.GetUserBadges()
	//if err != nil {
	//	fmt.Printf("User Badges Error: %v\n", err)
	//}
	//wc.GetUserProgressProlabs()
	//activeMachine, err := wc.GetActiveMachine()
	//if err != nil {
	//	fmt.Printf("Active Machine Error: %v\n", err)
	//}
	retired, err := wc.ListRetiredMachines()
	if err != nil {
		fmt.Printf("Retired Machines Error: %v\n", err)
	}
	active, err := wc.GetActiveMachine()
	if err != nil {
		fmt.Printf("Active Machine Error: %v\n", err)
	}

	//fmt.Sprintf("%v", ua)
	//fmt.Sprintf("%v", badges)
	//fmt.Sprintf("%v", activeMachine)
	fmt.Sprintf("%v", retired)
	fmt.Sprintf("%v", active)
	//fmt.Sprintf("%v", uqr)
	//fmt.Sprintf("%v", userInfo)
}
