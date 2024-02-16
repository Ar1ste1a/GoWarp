package warp

import (
	"GoWarp/htb"
	"fmt"
)

func (warp *Warp) connect(vpn htb.VPN) {
	filepath := warp.Vpn[vpn.String()]
	// Connect to the VPN
	fmt.Printf("[ %s ] Connecting to VPN: ", filepath)
	// Log output of VPN
}
