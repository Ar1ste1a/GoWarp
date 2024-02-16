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

func (warp *Warp) disconnect(vpn htb.VPN) {
	filepath := warp.Vpn[vpn.String()]
	// Disconnect from the VPN
	fmt.Printf("[ %s ] Disconnecting from VPN: ", filepath)
	// Log output of VPN
}

func (warp *Warp) killVPNs() {
	// sudo
	// killall
	// openvpn
	// Get active machine
	fmt.Println("Getting active machine")
}
