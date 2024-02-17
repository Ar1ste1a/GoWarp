package warp

import (
	"GoWarp/htb"
	"GoWarp/vpn"
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"net"
	"strings"
	"sync"
)

var (
	htbGreen    = tcell.NewRGBColor(159, 239, 0)
	black       = tcell.NewRGBColor(0, 0, 0)
	white       = tcell.NewRGBColor(255, 255, 255)
	buttonStyle = tcell.StyleDefault.Background(black).Foreground(htbGreen).Underline(true).Bold(true)
)

type CLI struct {
	Username         string
	user             UserBrief
	userPane         *tview.Flex
	setKeyButton     *tview.Button
	apiKeyInput      *tview.InputField
	userBox          *tview.Box
	vpnBox           *tview.Table
	machineTable     *tview.Table
	activeMachineBox *tview.Box
	rankingBox       *tview.Box
	app              *tview.Application
	apiSet           bool
	activeMachine    Machine
	ifaces           map[string][]string
	restart          chan bool
	targetIP         string
	lock             *sync.Mutex
	retiredMachines  []RetiredMachine
	activeMachines   []Machine
	activities       []UserActivity
	badges           []Badge
	prolabs          []ProLabProgress
}

func GetCLI(username string, apiSet bool, activeMachine Machine) *CLI {
	cli := &CLI{
		Username:      username,
		apiSet:        apiSet,
		activeMachine: activeMachine,
		restart:       make(chan bool),
	}
	cli.setIfaces()
	cli.setNewApp()
	go cli.handleUpdates()
	return cli
}

func (cli *CLI) getLock() *sync.Mutex {
	return cli.lock
}

func (cli *CLI) Start() {
	go cli.startApplication()
	for {
		select {
		case <-cli.restart:
			cli.app.Stop()
			cli.setNewApp()
			go cli.startApplication()
		}
	}
}

func (cli *CLI) getApplicationPanes() *tview.Flex {
	leftPane := cli.getLeftPane()
	rightPane := cli.getRightPane()

	// Get the top pane
	top := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, false).
		AddItem(rightPane, 0, 3, false)

	// Get the bottom pane
	bottom := cli.getBadgesPane()
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(top, 0, 10, false).
		AddItem(bottom, 0, 1, false)

	return flex
}

func (cli *CLI) getBadgesPane() *tview.Box {
	badgesPane := tview.NewBox()
	badgesPane.SetBorder(true).
		SetTitle("───────[ [::bl][purple]Badges[-] [::-]]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(htbGreen).
		SetTitleColor(htbGreen).
		SetBackgroundColor(black)
	return badgesPane
}

func (cli *CLI) getLeftPane() *tview.Flex {
	if cli.apiSet {
		return cli.getAuthenticatedLeftPane()
	} else {
		return cli.getUnauthenticatedLeftPane()
	}
}

func (cli *CLI) getRightPane() *tview.Flex {
	if cli.apiSet {
		return cli.getAuthenticatedRightPane()
	} else {
		return cli.getUnauthenticatedRightPane()
	}
}

// //////////////////////////////////////////////////////// LEFT PANE //////////////////////////////////////////////////////////
func (cli *CLI) getAuthenticatedLeftPane() *tview.Flex {
	apiStatus := "────────────[ [::b][green]API Key \u2713 [-][::-]]"
	userPaneTitle := "───────[ [::bl][purple]GoWarp[-][::-] [::-]]" + apiStatus
	vpnPaneTitle := "───────[ [::bl][purple]VPN[-] [::-]]"

	// Create User Box
	userBox := tview.NewBox()
	userBox.SetBorder(true).SetTitle(userPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.userBox = userBox

	// Create VPN Box
	vpnBox := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	vpnBox.SetCell(0, 0, tview.NewTableCell("VPN").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 1, tview.NewTableCell(" Available ").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 2, tview.NewTableCell("Status").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 3, tview.NewTableCell("Connect").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetBorder(true).SetTitle(vpnPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.vpnBox = vpnBox
	cli.fillVPNPane()

	cli.apiKeyInput = nil
	cli.setKeyButton = nil

	// Create Flex
	leftPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(userBox, 0, 2, false).
		AddItem(vpnBox, 0, 2, false)

	return leftPane
}

func (cli *CLI) getUnauthenticatedLeftPane() *tview.Flex {
	apiStatus := "────────────[ [::bl][red]API Key \u2717 [-][::-]]"
	userPaneTitle := "───────[ [::bl][purple]GoWarp[-][::-] ]" + apiStatus
	vpnPaneTitle := "───────[ [::bl][purple]VPN[-] [::-]]"

	// Create User Box
	userBox := tview.NewBox()
	userBox.SetBorder(true).SetTitle(userPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.userBox = userBox

	// Create VPN Box
	vpnBox := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	vpnBox.SetCell(0, 0, tview.NewTableCell("VPN").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 1, tview.NewTableCell(" Available ").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 2, tview.NewTableCell("Status").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 3, tview.NewTableCell("Connect").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetBorder(true).SetTitle(vpnPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.vpnBox = vpnBox
	cli.fillVPNPane()

	// Create API Key Input
	apiKeyInput := tview.NewInputField()
	apiKeyInput.SetBackgroundColor(black)
	apiKeyInput.SetFieldBackgroundColor(htbGreen)
	apiKeyInput.SetFieldTextColor(black)
	apiKeyInput.SetLabel("──[ API Key ] ")
	apiKeyInput.SetLabelColor(htbGreen)
	cli.apiKeyInput = apiKeyInput

	// Create Set API Key Button
	setKeyButton := tview.NewButton("[ [::bl][red]Set Key[-] [::-]]")
	setKeyButton.SetStyle(buttonStyle)
	setKeyButton.SetDisabled(false)
	setKeyButton.SetBorderColor(htbGreen)
	setKeyButton.SetBorder(true)
	setKeyButton.SetActivatedStyle(buttonStyle)
	setKeyButton.SetSelectedFunc(func() {
		cli.parseAPIKeyInput(apiKeyInput)
	})
	cli.setKeyButton = setKeyButton

	// Create Flex
	leftPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(apiKeyInput, 1, 1, false).
		AddItem(setKeyButton, 0, 2, false).
		AddItem(userBox, 0, 4, false).
		AddItem(vpnBox, 1, 5, false)

	return leftPane
}

// //////////////////////////////////////////////////////// LEFT PANE //////////////////////////////////////////////////////////
func (cli *CLI) getAuthenticatedRightPane() *tview.Flex {
	activeMachineTitle := "───────[ [::bl][purple]Active Machine[-] [::-]]───────"
	activityTitle := "───────[ [::bl][purple]Recent Activity[-] [::-]]───────"
	interfaceLine := cli.getIfacesLine()

	// ActiveMachine Box
	activeMachineBox := tview.NewBox()
	activeMachineBox.SetBorder(true).SetTitle(activeMachineTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activeMachineBox = activeMachineBox

	// Ranking Box
	activityBox := tview.NewBox()
	activityBox.SetBorder(true).SetTitle(activityTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.rankingBox = activityBox

	// Machines Table
	machineTable := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	machineTable.SetCell(0, 0, tview.NewTableCell("Avatar").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 1, tview.NewTableCell("Name").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 2, tview.NewTableCell("OS").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 3, tview.NewTableCell("Points").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 4, tview.NewTableCell("Stars").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 5, tview.NewTableCell("Free").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 6, tview.NewTableCell("Difficulty").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 7, tview.NewTableCell("Retired").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 8, tview.NewTableCell("Release").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	//machineTable.SetCell(0, 9, tview.NewTableCell("Root Own").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetBorder(true).SetTitle(interfaceLine).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.machineTable = machineTable

	// Create Flex
	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(machineTable, 0, 3, false).
		AddItem(activeMachineBox, 5, 2, false).
		AddItem(activityBox, 0, 1, false)

	return rightPane
}

func (cli *CLI) getUnauthenticatedRightPane() *tview.Flex {
	activeMachineTitle := "───────[ [::bl][purple]Active Machine[-] [::-]]───────"
	activityTitle := "───────[ [::bl][purple]Recent Activity[-] [::-]]───────"
	interfaceLine := cli.getIfacesLine()

	// ActiveMachine Box
	activeMachineBox := tview.NewBox()
	activeMachineBox.SetBorder(true).SetTitle(activeMachineTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activeMachineBox = activeMachineBox

	// Ranking Box
	activityBox := tview.NewBox()
	activityBox.SetBorder(true).SetTitle(activityTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.rankingBox = activityBox

	// Machines Table
	machineTable := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	machineTable.SetCell(0, 0, tview.NewTableCell("Avatar").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 1, tview.NewTableCell("Name").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 2, tview.NewTableCell("OS").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 3, tview.NewTableCell("Points").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 4, tview.NewTableCell("Stars").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 5, tview.NewTableCell("Free").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 6, tview.NewTableCell("Difficulty").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 7, tview.NewTableCell("Retired").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 8, tview.NewTableCell("Release").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	//machineTable.SetCell(0, 9, tview.NewTableCell("Root Owns").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetBorder(true).SetTitle(interfaceLine).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.machineTable = machineTable

	// Create Flex
	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(machineTable, 0, 3, false).
		AddItem(activeMachineBox, 5, 2, false).
		AddItem(activityBox, 0, 1, false)

	return rightPane
}

func (cli *CLI) getIfacesLine() string {
	interfaceLine := ""
	for iface, addrs := range cli.ifaces {
		for _, addr := range addrs {
			interfaceLine += fmt.Sprintf("─────[[::bl][purple] %s: %s [-][::-]]", iface, addr)
		}
	}
	interfaceLine += "─────"
	if cli.targetIP != "" {
		//sword := "▬▬ι═══════ﺤ "
		//sword3 := "▬▬ι═══════ﺤ    -═══════ι▬▬"
		sword2 := "[black]▬▬[yellow]ι[white]═══════ﺤ[-]"
		interfaceLine += fmt.Sprintf("%s[::bl][purple] Target: %s [-][::-]]", sword2, cli.targetIP)
	}
	return interfaceLine
}

func (cli *CLI) setNewApp() {
	cli.app = tview.NewApplication()
}

func (cli *CLI) setIfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		cli.ifaces = map[string][]string{}
	}

	data := make(map[string][]string)
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			cli.ifaces = map[string][]string{}
		}

		// Skip loopback and docker interfaces
		if i.Name == "lo" || strings.Contains(i.Name, "docker") {
			continue
		}

		// Only add IPv4 addresses
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && ipnet.IP.To4() != nil {
				data[i.Name] = append(data[i.Name], addr.String())
			}
		}
	}
	cli.ifaces = data
}

func (cli *CLI) startApplication() {
	cli.app.SetRoot(cli.getApplicationPanes(), true).EnableMouse(true)
	if err := cli.app.Run(); err != nil {
		panic(err)
	}
}

func (cli *CLI) parseAPIKeyInput(apiKeyInput *tview.InputField) {
	// Retrieve the text from the input field
	apiKey := apiKeyInput.GetText()
	apiKeyTrimmed := strings.Trim(apiKey, " ")
	if apiKeyTrimmed == "" {
		// Change the text of the input field to "API Key Required"
		apiKeyInput.SetText("API Key Required")

		return
	} else if err := TestAPIKey(apiKeyTrimmed); err != nil {
		apiKeyInput.SetText("Invalid Key")
	} else {
		cli.apiSet = true
		cli.restart <- true
		fromCLI <- NewApiKey{ApiKey: apiKeyTrimmed}
	}
}

func (cli *CLI) handleUpdates() {
	for {
		select {
		case update := <-toCLI:
			switch update.GetType() {
			case CURRENT_USER:
				cli.UpdateCurrentUser(update)
			case ACTIVE_MACHINE:
				cli.UpdateActiveMachine(update)
			case MACHINES_LIST:
				cli.UpdateMachinesList(update)
			case ACTIVITY_LIST:
				cli.UpdateRecentActivityList(update)
			case VPN_LIST:

				cli.UpdateVPNList(update)
			case VPN_STATUS:

				cli.UpdateVPNStatus(update)
			case BADGES_LIST:
				cli.updateBadges(update)
			case PROLABS_PROGRESS:
				cli.updateProlabsProgress(update)
			default:
				return
			}
		}
	}
}

//// Listener for fromWarp to the application
//func (cli *CLI) Update(update Update) {
//	lock := cli.getLock()
//
//	switch update.GetType() {
//	case CURRENT_USER:
//		cli.UpdateCurrentUser(update)
//	case ACTIVE_MACHINE:
//		cli.UpdateActiveMachine(update)
//	case MACHINES_LIST:
//		cli.UpdateMachinesList(update)
//	case ACTIVITY_LIST:
//		cli.UpdateRecentActivityList(update)
//	case VPN_LIST:
//
//		cli.UpdateVPNList(update)
//	case VPN_STATUS:
//
//		cli.UpdateVPNStatus(update)
//	case BADGES_LIST:
//		cli.updateBadges(update)
//	case PROLABS_PROGRESS:
//		cli.updateProlabsProgress(update)
//	default:
//		return
//	}
//
//	lock.Unlock()
//}

func (cli *CLI) updateProlabsProgress(update Update) {
	var prolabProgress []ProLabProgress

	data := update.GetUpdate()

	// Convert the string to a list of ProLabProgress
	prolabProgressString := data["ProLabProgress"]
	err := json.Unmarshal([]byte(prolabProgressString), &prolabProgress)
	if err != nil {
		panic(err) // todo log here
	}

	// Update the ProLabProgress
	cli.prolabs = prolabProgress

	// Restart the Application
	cli.restart <- true
}

func (cli *CLI) updateBadges(update Update) {
	var badges []Badge

	data := update.GetUpdate()

	// Convert the string to a list of Badges
	badgesString := data["Badges"]
	err := json.Unmarshal([]byte(badgesString), &badges)
	if err != nil {
		panic(err) // todo log here
	}

	// Update the Badges
	cli.badges = badges

	cli.restart <- true
}

func (cli *CLI) UpdateCurrentUser(update Update) {
	var user UserBrief

	data := update.GetUpdate()

	cli.Username = data["Username"]
	userString := data["User"]
	err := json.Unmarshal([]byte(userString), &user)
	if err != nil {
		panic(err)
	}
	cli.user = user

	// Restart the Application
	cli.restart <- true
}

func (cli *CLI) UpdateActiveMachine(update Update) {
	var machine Machine

	data := update.GetUpdate()
	machinesString := data["ActiveMachine"]
	err := json.Unmarshal([]byte(machinesString), &machine)
	if err != nil {
		panic(err)
	}

	// Update the Active Machine
	cli.activeMachine = machine

	// Restart the Application
	cli.restart <- true
}

func (cli *CLI) UpdateMachinesList(update Update) {
	var machines []Machine

	data := update.GetUpdate()
	machinesListString := data["Machines"]
	err := json.Unmarshal([]byte(machinesListString), &machines)
	if err != nil {
		panic(err)
	}

	// Update the Machines List
	cli.activeMachines = machines

	// Restart the Application
	cli.restart <- true
}

func (cli *CLI) UpdateRecentActivityList(update Update) {
	data := update.GetUpdate()

	// Convert the string to a list of UserActivity
	activitiesString := data["Activities"]
	err := json.Unmarshal([]byte(activitiesString), &cli.activities)
	if err != nil {
		panic(err) // todo log here
	}

	cli.restart <- true
}

func (cli *CLI) UpdateVPNList(update Update) {
	data := update.GetUpdate()
	fmt.Printf("VPN List: %v\n", data)
	cli.restart <- true
}

func (cli *CLI) UpdateVPNStatus(update Update) {
	data := update.GetUpdate()
	fmt.Printf("VPN Status: %v\n", data)
	cli.restart <- true
}

// //////////// FILL APPLICATION PANES //////////////
func (cli *CLI) fillVPNPane() {
	vpnServer := htb.VpnServers
	rows := len(vpnServer) + 1
	row := 1

	for servername, id := range vpnServer {
		if !(row < rows) {
			break
		}
		servername := fmt.Sprintf(" %s ", servername)
		cli.vpnBox.SetCell(row, 0,
			tview.NewTableCell(servername).
				SetTextColor(white).
				SetSelectable(false).
				SetAlign(tview.AlignLeft).
				SetExpansion(2))
		cli.vpnBox.SetCell(row, 1,
			tview.NewTableCell("Yes").
				SetTextColor(tcell.ColorGreen).
				SetSelectable(false).
				SetAlign(tview.AlignCenter).
				SetExpansion(2))
		cli.vpnBox.SetCell(row, 2,
			tview.NewTableCell(" Disconnected ").
				SetTextColor(tcell.ColorRed).
				SetSelectable(false).
				SetAlign(tview.AlignCenter).
				SetExpansion(2))
		cli.vpnBox.SetCell(row, 3,
			tview.NewTableCell(" Connect ").
				SetTextColor(htbGreen).
				SetSelectable(true).
				SetAlign(tview.AlignRight).
				SetExpansion(2).
				SetClickedFunc(func() bool {
					setVPN(servername, id)
					return false
				}))
		row++
	}
}

func (cli *CLI) fillMachinePane() {
	rowsActive := len(cli.activeMachines) + 1
	rowsRetired := len(cli.retiredMachines) + 1
	row := 1

	if rowsActive > 1 {
		for _, machine := range cli.activeMachines {
			if !(row < rowsActive) {
				break
			}

			cli.machineTable.SetCell(row, 0,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Avatar)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 1,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Name)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 2,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Os)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 3,
				tview.NewTableCell(fmt.Sprintf(" %d ", machine.Points)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 4,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Star)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 5,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Free)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 6,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.DifficultyText)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 7,
				tview.NewTableCell(fmt.Sprintf(" %v ", false)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 8,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Release)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			row++
		}
	}
	if rowsRetired > 1 {
		row = 1
		for _, machine := range cli.retiredMachines {
			if !(row < rowsRetired) {
				break
			}

			cli.machineTable.SetCell(row, 0,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Avatar)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 1,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Name)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 2,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Os)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 3,
				tview.NewTableCell(fmt.Sprintf(" %d ", machine.Points)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 4,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Star)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 5,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Free)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 6,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.DifficultyText)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 7,
				tview.NewTableCell(fmt.Sprintf(" %v ", true)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 8,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Release)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(2))
			row++
		}
	}
	// TODO
}

func (cli *CLI) fillActivityPane() {

}

func setVPN(servername, id string) bool {
	vpn.ConnectVPN(servername, id)
	return false
}
