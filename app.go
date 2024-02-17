package main

import (
	"GoWarp/warp"
)

func init() {
	//cobra.AddTemplateFunc("getActiveMachine", )
}

func main() {
	warpClient, err := warp.GetWarpClient()
	if err != nil {
		panic(err)
	}
	warpClient.Start()
}
