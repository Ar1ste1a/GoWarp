package warp

import (
	"GoWarp/htb"
	"GoWarp/util"
	"GoWarp/vpn"
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"image"
	"net"
	"path"
	"strings"
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
	activityBox      *tview.Table
	badgeBox         *tview.Table
	proLabBox        *tview.Table
	app              *tview.Application
	apiSet           bool
	activeMachine    Machine
	interfaces       map[string][]string
	restart          chan bool
	targetIP         string
	imageFolder      string
	retiredMachines  []RetiredMachine
	activeMachines   []Machine
	activities       []UserActivity
	badges           []Badge
	proLabs          []ProLabProgress
	vpnProtocol      string
}

func GetCLI(username, folder string, apiSet bool, activeMachine Machine) *CLI {
	cli := &CLI{
		Username:      username,
		apiSet:        apiSet,
		activeMachine: activeMachine,
		restart:       make(chan bool),
		imageFolder:   folder,
	}
	cli.setDefaultAvatar()
	cli.setInterfaces()
	cli.setNewApp()
	go cli.handleUpdates()
	return cli
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

func (cli *CLI) setDefaultAvatar() {
	cli.activeMachine.Avatar = path.Join(cli.imageFolder, "gowarp-min.png")
}

func (cli *CLI) setVPNProtocol(protocol string) {
	cli.vpnProtocol = protocol
}

func (cli *CLI) getApplicationPanes() *tview.Flex {
	leftPane := cli.getLeftPane()
	rightPane := cli.getRightPane()

	// Get the top pane
	top := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 1, false).
		AddItem(rightPane, 0, 3, false)

	// Get the bottom pane
	bottom := cli.getBottomPane()
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(top, 0, 10, false).
		AddItem(bottom, 0, 1, false)

	return flex
}

func (cli *CLI) getBottomPane() *tview.Flex {
	badgesPane := tview.NewTable()
	badgesPane.SetBorders(false).
		SetSelectable(false, false).
		SetBorder(true).
		SetBorderColor(htbGreen).
		SetTitle("───────[ [::b][purple]Badges[-] [::-]]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(htbGreen).
		SetTitleColor(htbGreen).
		SetBackgroundColor(black)
	cli.badgeBox = badgesPane
	cli.fillBadgesPane()

	proLabsPane := tview.NewTable()
	proLabsPane.SetBorders(false).
		SetBorder(true).
		SetBorderColor(htbGreen).
		SetTitle("───────[ [::b][purple]ProLabs[-] [::-]]───────").
		SetTitleAlign(tview.AlignRight).
		SetTitleColor(htbGreen).
		SetBorderColor(htbGreen).
		SetBackgroundColor(black)
	cli.proLabBox = proLabsPane
	cli.fillProLabsPane()

	botFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(badgesPane, 0, 1, false).
		AddItem(proLabsPane, 0, 3, false)

	return botFlex
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
	userPaneTitle := "───────[ [::b][purple]GoWarp[-][::-] [::-]]" + apiStatus
	vpnPaneTitle := "───────[ [::b][purple]VPN[-] [::-]]"

	// Create User Box
	userBox := tview.NewBox()
	userBox.SetBorder(true).SetTitle(userPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.userBox = userBox

	// Create VPN Box
	vpnBox := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	vpnBox.SetCell(0, 0, tview.NewTableCell("VPN").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 1, tview.NewTableCell("Status").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 2, tview.NewTableCell("Connect").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetBorder(true).SetTitle(vpnPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	vpnBox.SetSelectedStyle(tcell.StyleDefault.Background(black).Foreground(tcell.ColorCadetBlue))
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
	apiStatus := "────────────[ [::b][red]API Key \u2717 [-][::-]]"
	userPaneTitle := "───────[ [::b][purple]GoWarp[-][::-] ]" + apiStatus
	vpnPaneTitle := "───────[ [::b][purple]VPN[-] [::-]]"

	// Create User Box
	userBox := tview.NewBox()
	userBox.SetBorder(true).SetTitle(userPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.userBox = userBox

	// Create VPN Box
	vpnBox := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	vpnBox.SetCell(0, 0, tview.NewTableCell("VPN").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 1, tview.NewTableCell("Status").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetCell(0, 2, tview.NewTableCell("Connect").SetTextColor(htbGreen).SetSelectable(false).SetAlign(tview.AlignCenter))
	vpnBox.SetBorder(true).SetTitle(vpnPaneTitle).SetTitleAlign(tview.AlignLeft).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	vpnBox.SetSelectedStyle(tcell.StyleDefault.Background(black).Foreground(tcell.ColorCadetBlue))
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
	setKeyButton := tview.NewButton("[ [::b][red]Set Key[-] [::-]]")
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

// //////////////////////////////////////////////////////// TCP UDP MODAL ///////////////////////////////////////////////////////
//func (cli *CLI) getTCPUDPModal() *tview.Modal {
//	modal := tview.NewModal().
//		SetText("Would you like to connect over TCP or UDP?").
//		AddButtons([]string{"TCP", "UDP", "Cancel"}).
//		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "Cancel" {
//				app.Stop()
//			} else if buttonLabel == "TCP" {
//				cli.setVPNProtocol(htb.TCP)
//			} else {
//				cli.setVPNProtocol(htb.UDP)
//			}
//		})
//}

// //////////////////////////////////////////////////////// RIGHT PANE //////////////////////////////////////////////////////////
func (cli *CLI) getAuthenticatedRightPane() *tview.Flex {
	activeMachineTitle := "───────[ [::b][purple]Active Machine[-] [::-]]───────"
	activityTitle := "───────[ [::b][purple]Recent Activity[-] [::-]]───────"
	interfaceLine := cli.getInterfacesLine()

	// ActiveMachine Box
	activeMachineBox := tview.NewBox()
	activeMachineBox.SetBorder(true).SetTitle(activeMachineTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activeMachineBox = activeMachineBox

	picture := tview.NewImage()
	photo := cli.getActiveMachinePhoto()
	picture.SetImage(photo)
	picture.SetBorder(true).SetBorderColor(htbGreen).SetBackgroundColor(black)

	midFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(picture, 0, 1, false).
		AddItem(activeMachineBox, 0, 9, false)

	// Activity Box
	activityBox := tview.NewTable()
	activityBox.SetBorders(false)
	activityBox.SetBorder(true).SetTitle(activityTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activityBox = activityBox
	cli.fillActivityPane()

	// Machines Table
	machineTable := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	machineTable.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 1, tview.NewTableCell("OS").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 2, tview.NewTableCell("Points").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 3, tview.NewTableCell("Stars").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 4, tview.NewTableCell("Free").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 5, tview.NewTableCell("Difficulty").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 6, tview.NewTableCell("Retired").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 7, tview.NewTableCell("Release").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 8, tview.NewTableCell("Start").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetBorder(true).SetTitle(interfaceLine).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.machineTable = machineTable
	if len(cli.activeMachines) > 0 || len(cli.retiredMachines) > 0 {
		cli.fillMachinePane()
	}

	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(machineTable, 0, 8, false).
		AddItem(midFlex, 5, 1, false).
		AddItem(activityBox, 0, 1, false)

	return rightPane
}

func (cli *CLI) getUnauthenticatedRightPane() *tview.Flex {
	activeMachineTitle := "───────[ [::b][purple]Active Machine[-] [::-]]───────"
	activityTitle := "───────[ [::b][purple]Recent Activity[-] [::-]]───────"
	interfaceLine := cli.getInterfacesLine()

	// ActiveMachine Box
	activeMachineBox := tview.NewBox()
	activeMachineBox.SetBorder(true).SetTitle(activeMachineTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activeMachineBox = activeMachineBox

	picture := tview.NewImage()
	photo := cli.getActiveMachinePhoto()
	picture.SetImage(photo)
	picture.SetBorder(true).SetBorderColor(htbGreen).SetBackgroundColor(black)

	midFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(picture, 0, 1, false).
		AddItem(activeMachineBox, 0, 9, false)

	// Ranking Box
	activityBox := tview.NewTable()
	activityBox.SetBorders(false)
	activityBox.SetBorder(true).SetTitle(activityTitle).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.activityBox = activityBox

	// Machines Table
	machineTable := tview.NewTable().SetBorders(true).SetBordersColor(htbGreen).SetSelectable(true, false)
	machineTable.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 1, tview.NewTableCell("OS").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 2, tview.NewTableCell("Points").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 3, tview.NewTableCell("Stars").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 4, tview.NewTableCell("Free").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 5, tview.NewTableCell("Difficulty").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 6, tview.NewTableCell("Retired").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 7, tview.NewTableCell("Release").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetCell(0, 8, tview.NewTableCell("Start").SetTextColor(htbGreen).SetSelectable(false).SetExpansion(2).SetAlign(tview.AlignCenter))
	machineTable.SetBorder(true).SetTitle(interfaceLine).SetTitleAlign(tview.AlignRight).SetBorderColor(htbGreen).SetTitleColor(htbGreen).SetBackgroundColor(black)
	cli.machineTable = machineTable

	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(machineTable, 0, 3, false).
		AddItem(midFlex, 5, 2, false).
		AddItem(activityBox, 0, 1, false)

	return rightPane
}

func (cli *CLI) getInterfacesLine() string {
	interfaceLine := ""
	for name, addresses := range cli.interfaces {
		for _, addr := range addresses {
			interfaceLine += fmt.Sprintf("─────[[::b][purple] %s: %s [-][::-]]", name, addr)
		}
	}
	interfaceLine += "─────"
	if cli.targetIP != "" {
		//sword := "▬▬ι═══════ﺤ "
		//sword3 := "▬▬ι═══════ﺤ    -═══════ι▬▬"
		sword2 := "[black]▬▬[yellow]ι[white]═══════ﺤ[-]"
		interfaceLine += fmt.Sprintf("%s[::b][purple] Target: %s [-][::-]]", sword2, cli.targetIP)
	}
	return interfaceLine
}

func (cli *CLI) setNewApp() {
	cli.app = tview.NewApplication()
}

func (cli *CLI) setInterfaces() {
	interfaces, err := net.Interfaces()
	if err != nil {
		cli.interfaces = map[string][]string{}
	}

	data := make(map[string][]string)
	for _, i := range interfaces {
		addresses, err := i.Addrs()
		if err != nil {
			cli.interfaces = map[string][]string{}
		}

		// Skip loopback and docker interfaces
		if i.Name == "lo" || strings.Contains(i.Name, "docker") {
			continue
		}

		// Only add IPv4 addresses
		for _, addr := range addresses {
			ip, ok := addr.(*net.IPNet)
			if ok && ip.IP.To4() != nil {
				data[i.Name] = append(data[i.Name], addr.String())
			}
		}
	}
	cli.interfaces = data
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
			case RETIRED_MACHINES_LIST:
				cli.updateRetiredMachines(update)
			case ACTIVITY_LIST:
				cli.UpdateRecentActivityList(update)
			case VPN_LIST:

				cli.UpdateVPNList(update)
			case VPN_STATUS:

				cli.UpdateVPNStatus(update)
			case BADGES_LIST:
				cli.updateBadges(update)
			case PROLABS_PROGRESS:
				cli.updateProLabsProgress(update)
			default:
				return
			}
		}
	}
}

func (cli *CLI) updateProLabsProgress(update Update) {
	var proLabProgress []ProLabProgress

	data := update.GetUpdate()
	proLabProgressString := data["Progress"]
	plpSplit := strings.Split(proLabProgressString, "${delim}")

	for _, jsonString := range plpSplit {
		var plp ProLabProgress
		err := json.Unmarshal([]byte(jsonString), &plp)
		if err != nil {
			panic(err) // todo log here
		}
		proLabProgress = append(proLabProgress, plp)

	}

	// Update the ProLabProgress
	cli.proLabs = proLabProgress

	// Restart the Application
	cli.fillProLabsPane()
	//cli.restart <- true
}

func (cli *CLI) updateBadges(update Update) {
	var badges []Badge

	data := update.GetUpdate()

	// Convert the string to a list of Badges
	badgesString := data["Badges"]
	badgesSplit := strings.Split(badgesString, "${delim}")

	for _, badgeString := range badgesSplit {
		var badge Badge
		err := json.Unmarshal([]byte(badgeString), &badge)
		if err != nil {
			panic(err) // todo log here
		}
		badges = append(badges, badge)
	}

	// Update the Badges
	cli.badges = badges

	cli.fillBadgesPane()

	//cli.restart <- true
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
	//cli.restart <- true
}

func (cli *CLI) UpdateActiveMachine(update Update) {
	var machine Machine

	data := update.GetUpdate()
	machinesString := data["Machine"]
	err := json.Unmarshal([]byte(machinesString), &machine)
	if err != nil {
		panic(err)
	}

	// Update the Active Machine
	cli.activeMachine = machine

	// Restart the Application
	//cli.restart <- true
}

func (cli *CLI) UpdateMachinesList(update Update) {
	var machines []Machine

	data := update.GetUpdate()
	machinesListString := data["Machines"]
	machinesSplit := strings.Split(machinesListString, "${delim}")
	for _, machineString := range machinesSplit {
		var machine Machine
		err := json.Unmarshal([]byte(machineString), &machine)
		if err != nil {
			panic(err) // todo log here
		}
		machines = append(machines, machine)
	}

	// Update the Machines List
	cli.activeMachines = machines

	// Restart the Application
	cli.fillMachinePane()
	//cli.restart <- true
}

func (cli *CLI) updateRetiredMachines(update Update) {
	var machines []RetiredMachine

	data := update.GetUpdate()
	machinesListString := data["Machines"]
	machinesSplit := strings.Split(machinesListString, "${delim}")
	for _, machineString := range machinesSplit {
		var machine RetiredMachine
		err := json.Unmarshal([]byte(machineString), &machine)
		if err != nil {
			panic(err) // todo log here
		}
		machines = append(machines, machine)
	}

	// Update the Machines List
	cli.retiredMachines = machines

	// Restart the Application
	//cli.fillMachinePane()
	cli.restart <- true
}

func (cli *CLI) UpdateRecentActivityList(update Update) {
	data := update.GetUpdate()
	var activities []UserActivity

	// Convert the string to a list of UserActivity
	activitiesString := data["Activity"]
	activitiesSplit := strings.Split(activitiesString, "${delim}")
	for _, activityString := range activitiesSplit {
		var activity UserActivity
		err := json.Unmarshal([]byte(activityString), &activity)
		if err != nil {
			panic(err) // todo log here
		}
		activities = append(activities, activity)
	}

	cli.activities = activities
	cli.fillActivityPane()

	//cli.restart <- true
}

func (cli *CLI) UpdateVPNList(update Update) {
	data := update.GetUpdate()
	fmt.Printf("VPN List: %v\n", data)
	//cli.restart <- true
}

func (cli *CLI) UpdateVPNStatus(update Update) {
	data := update.GetUpdate()
	fmt.Printf("VPN Status: %v\n", data)
	//cli.restart <- true
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
			tview.NewTableCell("N/A").
				SetTextColor(tcell.ColorRed).
				SetSelectable(false).
				SetAlign(tview.AlignCenter).
				SetExpansion(2))
		cli.vpnBox.SetCell(row, 2,
			tview.NewTableCell("Connect").
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
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Name)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 1,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Os)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 2,
				tview.NewTableCell(fmt.Sprintf(" %d ", machine.Points)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 3,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Star)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 4,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Free)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 5,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.DifficultyText)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 6,
				tview.NewTableCell(fmt.Sprintf(" %v ", false)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 7,
				tview.NewTableCell(fmt.Sprintf(" %s ", util.HumanReadableDate(machine.Release))).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 8,
				tview.NewTableCell(fmt.Sprintf(" %s ", "✓")).
					SetTextColor(tcell.ColorRed).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(0).
					SetSelectable(true))
			// todo add function for click
			row++
		}
	}
	if rowsRetired > 1 {
		for _, machine := range cli.retiredMachines {
			if !(row < rowsRetired+rowsActive) {
				break
			}

			cli.machineTable.SetCell(row, 0,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Name)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 1,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.Os)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 2,
				tview.NewTableCell(fmt.Sprintf(" %d ", machine.Points)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 3,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Star)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 4,
				tview.NewTableCell(fmt.Sprintf(" %v ", machine.Free)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 5,
				tview.NewTableCell(fmt.Sprintf(" %s ", machine.DifficultyText)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 6,
				tview.NewTableCell(fmt.Sprintf(" %v ", true)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 7,
				tview.NewTableCell(fmt.Sprintf(" %s ", util.HumanReadableDate(machine.Release))).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(2))
			cli.machineTable.SetCell(row, 8,
				tview.NewTableCell(fmt.Sprintf(" %s ", "✓")).
					SetTextColor(tcell.ColorRed).
					SetSelectable(false).
					SetAlign(tview.AlignCenter).
					SetExpansion(0).
					SetSelectable(true))
			// todo set function for click
			row++
		}
	}
}

func (cli *CLI) fillActivityPane() {
	var cols int
	var printColor tcell.Color

	if len(cli.activities) > 0 {
		if len(cli.activities) >= 10 {
			cols = 11
		} else {
			cols = len(cli.activities) + 1
		}
		col := 0

		for _, activity := range cli.activities {
			if !(col < cols) {
				break
			}

			if col%2 == 1 {
				printColor = tcell.ColorPurple
			} else {
				printColor = htbGreen
			}

			cli.activityBox.SetCell(0, col,
				tview.NewTableCell(fmt.Sprintf(" %s ", activity.Name)).
					SetTextColor(printColor).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.activityBox.SetCell(1, col,
				tview.NewTableCell(fmt.Sprintf(" Type: %s ", activity.Type)).
					SetTextColor(printColor).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.activityBox.SetCell(2, col,
				tview.NewTableCell(fmt.Sprintf(" Pts: %d ", activity.Points)).
					SetTextColor(printColor).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.activityBox.SetCell(2, col,
				tview.NewTableCell(fmt.Sprintf(" %s ", util.HumanReadableDate(activity.Date))).
					SetTextColor(printColor).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			col++
		}

	}
}

func (cli *CLI) fillActiveMachinePane() {

}

func (cli *CLI) fillBadgesPane() {
	if len(cli.badges) > 0 {
		cols := len(cli.badges) + 1
		col := 1
		left := true
		i := 0
		row := 0

		for _, badge := range cli.badges {
			if !(col < cols) {
				break
			}

			if i == 2 {
				row += 2
				i = 0
			}

			if left {
				col = 0
			} else {
				col = 1
			}

			cli.badgeBox.SetCell(row, col,
				tview.NewTableCell(fmt.Sprintf(" %s ", badge.Name)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.badgeBox.SetCell(row+1, col,
				tview.NewTableCell(fmt.Sprintf(" %s ", badge.DescriptionEn)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			left = !left
			i++
			col++
		}
	}
}

func (cli *CLI) fillProLabsPane() {
	var nameColor tcell.Color
	var color string

	if len(cli.proLabs) > 0 {
		cols := len(cli.proLabs) + 1
		col := 0

		for _, lab := range cli.proLabs {
			if !(col < cols) {
				break
			}

			if lab.CompletionPercentage == 100 {
				nameColor = tcell.ColorGreen
				color = "green"
			} else if lab.CompletionPercentage >= 50 {
				nameColor = tcell.ColorYellow
				color = "yellow"
			} else if lab.CompletionPercentage >= 25 {
				nameColor = tcell.ColorCadetBlue
				color = "blue"
			} else {
				nameColor = tcell.ColorRed
				color = "red"
			}

			cli.proLabBox.SetCell(0, col,
				tview.NewTableCell(fmt.Sprintf(" %s ", lab.Name)).
					SetTextColor(nameColor).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.proLabBox.SetCell(1, col,
				tview.NewTableCell(fmt.Sprintf(" %d/%d [%s](%v)", lab.OwnedFlags, lab.TotalFlags, color, lab.CompletionPercentage)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			cli.proLabBox.SetCell(2, col,
				tview.NewTableCell(fmt.Sprintf(" Avg. Rating: %v ", lab.AverageRatings)).
					SetTextColor(white).
					SetSelectable(false).
					SetAlign(tview.AlignLeft).
					SetExpansion(3))
			col++
		}
	}
}

func setVPN(servername, id string) bool {
	vpn.ConnectVPN(servername, id)
	return false
}

func (cli *CLI) getActiveMachinePhoto() image.Image {
	if cli.activeMachine.Avatar == "" {
		cli.setDefaultAvatar()
	}
	avatar := strings.Split(cli.activeMachine.Avatar, "/")[len(strings.Split(cli.activeMachine.Avatar, "/"))-1]
	avatarPath := path.Join(cli.imageFolder, avatar)
	photo := util.GetImage(avatarPath)
	return photo
}
