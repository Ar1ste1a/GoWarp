package warp

import (
	"GoWarp/htb"
	"GoWarp/util"
	"encoding/json"
	"fmt"
	"net/http"
	urlLib "net/url"
	"os"
	"path"
	"strings"
	"time"
)

// Kris eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiNGUzY2YxNDA3OTI2ZDQ0OWNiNDgzMjNmY2ViNjg5OWM3MWUwNzlkNTZjYWY0ZmExY2Q1YjFkODViMjgzOWExM2MxMTk0MWVjYjdlODI4ZjAiLCJpYXQiOjE3MDgwOTM2MTMuODE1OTAzLCJuYmYiOjE3MDgwOTM2MTMuODE1OTA1LCJleHAiOjE3MDgxODAwMTMuODEwODk0LCJzdWIiOiI5NTI5NjYiLCJzY29wZXMiOlsiMmZhIl19.VkG4JJ75T_38BKGHKb6kLm5wKZDZ_1rIji5vkECLL8DVR3lKfKLXXN9TrG7L1AR2N8tziZu_R92q8cQ_6t8Ue58D3s8GWlPP8LgvTNEfjy7jHZnd6I_XQKrWcRSFmFlSkxvmOKFAYIjfxYtiK0KY1FjPkC3OZkx-ZOaFeyv4F9qSkahiKHNjQXR1dnvhn1hiQKhQcvuDKP-QZ9pP3626hvGTew8w9LdklnbjE-gM7wxQiRRiPeGlykuhDRYdIdvcRmrC1FBDolBaJarpgxoStshWLh9vUdcrFvjFKl_nRknlM2Zmpk1e1roebZD2m9-rwOvezTv2Zit6jvR3NRk_lsNJX8mwUE46SoBXEkY7rneqRiuiebedMu0gfjRKeas5NMgqm9IjUBuShJT0mb5Xo5mNPPXsAphEbXV_TeKuRtZMGI3MYxkX0fXV3n37PKKbAoESAQDFP9D5yIy6LmzizYNhpmUcoxGLzyN101uVeH7R9uE0UnAiyGUQh5TQNd7ShSvE0Ja68FP1kdwaN35-N0G9NIODODDOYpYz3clMuCVweP7pHWTwLRnqpGGpGWiR0ZjzwhpIjmeAy7BdFctMXnpt3cX7pWbDbrUZATQVg7c5ljuomVzvZjyka52QicCP7HcHvnNgzvx359I2_rH7E0NZAy5bZqADZooIkDN7uBE
// Evan eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiNDJkZmQwYjVkNDMwZmJlNWI1MzFiYmFlZDZhZTZiMmY0ZmNlM2ZiZGRjMjUzMWM5ZmZmM2RlMzg1N2I2MDc4NTRlZWUyNDJmMjAwNWVhOTAiLCJpYXQiOjE3MDgwNDE0NTUuMTYzMzEsIm5iZiI6MTcwODA0MTQ1NS4xNjMzMTEsImV4cCI6MTczOTU3NzQ1NS4xNTQ0ODMsInN1YiI6IjEzMTM0MjUiLCJzY29wZXMiOlsiMmZhIl19.o9J5W1hX8amzQase1G2UKnAEZMHNdrfboCGK0n3epaCClVCRNR78BIVCykKNVC1LIhzv7fm9qOJxSnCHlFGwkUiclEXSFASJ1ywiWorcLXWSGIhdqUQ11v3_Jv4XIRoR74Lp2PVVizhyas55-wzmw0j-qLWynsjUOVsPR5TJHo_IOj8bErckXh-VcUEAXZTNNsXZS3vML3QEJyQpWrecyGvWlsAsKRO_PuI7FYygB8_SLjf65sB1QsAlFUxhriB04ogC4grocXf5efOfbF-0QTqRmv26qWwoSxUYsAcVCYvySPIk20HtvLkOF0ik04VEboxSQ0GCZ3q1sWNunwX-PdU8HVKpnGRT3I5hM-oAJBPTB53iq-1zt5tz6yjBWRYH50AGb6lgdtqZ1tHc18R5jyYSE6-3ICZKrT0kysH6eSFzqi04cAoRTXIE4dA739xb9dZZUU4qqVx4JFx3TP74BYAXWcx5ocmvoKYBU8NY2r6_4h1w61Cffpgeo4gmnxuNIfAGnYJmww5NGG_LyXqgYzOKCwEoahhKDYrEGu9ulZlHMyaGXifmyG9mhGpvHzMATDgn-gMvBNhKTV5QHFynloDmCBNlyNQQuHm6KiLF2J-QMJtGdXWwRdDqe4tc88qjL1Mks5BuK8UhQPRH6lVrB8mJf5qJHDVYAwGsMoTfhK4

var (
	toCLI   chan Update
	fromCLI chan Update
)

type Warp struct {
	ApiKey          string            `json:"ApiKey"`
	Vpn             map[string]string `json:"Vpn"`
	ConfPath        string            `json:"confPath"`
	FolderPath      string            `json:"filePath"`
	User            UserBrief         `json:"user"`
	ProlabsProgress []ProLabProgress  `json:"prolabsProgress"`
	Badges          []Badge           `json:"badges"`
	Activity        []UserActivity    `json:"activity"`
	currentVPN      string
	client          *http.Client
	req             *http.Request
	headers         map[string]string
	data            map[string]string
	machines        map[string]Machine
	prolabs         map[string]Machine
	cli             *CLI
}

/*
NewWarp

Return: *Warp

	This function is used to create a new Warp client.
*/
func newWarp(confPath, filePath string) *Warp {
	ghtb := &Warp{
		ApiKey:     "",
		client:     &http.Client{},
		req:        nil,
		headers:    map[string]string{},
		data:       map[string]string{},
		Vpn:        map[string]string{},
		ConfPath:   confPath,
		FolderPath: filePath,
	}
	return ghtb
}

/*
GetWarpClient

Return: *Warp, error

	This function is used to get the Warp client.
	If the config folder does not exist, it will be created and a new Warp client will be generated.
	If the config folder exists, the config file will be checked,
	if it does not exist a new Warp client will be generated,
	else the existing one will be loaded.
*/
func GetWarpClient() (*Warp, error) {
	var (
		warpClient *Warp
		err        error
	)
	toCLI = make(chan Update, 5)
	fromCLI = make(chan Update, 5)

	home, err := os.UserHomeDir()
	confFolder := fmt.Sprintf("%s/.htb", home)
	confPath := path.Join(confFolder, ".htb.conf")
	fileFolder := path.Join(confFolder, "images")

	if err != nil {
		return nil, err
	}

	// If the config folder does not exist, create it and generate a new Warp Client
	if !util.DirExists(confFolder) {
		err = os.Mkdir(confFolder, 0755)
		warpClient = newWarp(confPath, fileFolder)
	} else {
		// If the config folder exists, check if the config file exists, if not generate a new Warp Client else load the existing one
		if util.HTBExists(confPath) {
			warpClient, err = loadHTB(confPath, fileFolder)
		} else {
			warpClient = newWarp(confPath, fileFolder)
		}
	}

	// If the file folder does not exist, create it
	if !util.DirExists(fileFolder) {
		err = os.Mkdir(fileFolder, 0755)
	}

	// If the API key is set
	if warpClient.apiSet() {
		userInfoResponse, err := warpClient.GetUserInfo()
		go warpClient.grabImage(userInfoResponse.Info.Avatar, fileFolder)

		// If the user is not set, set the user
		if !warpClient.userSet() {
			if err == nil {
				warpClient.setUser(userInfoResponse.Info)
			}
		}

		// Set the user id and server id
		warpClient.setData("server_id", fmt.Sprintf("%d", userInfoResponse.Info.ServerID))
		warpClient.setData("user_id", fmt.Sprintf("%d", userInfoResponse.Info.ID))
	}

	return warpClient, err
}

/*
Start

Return: void

	This function is used to start the Warp client.
	This function will start the CLI and listen for fromWarp.
*/
func (warp *Warp) Start() {
	machine, _ := warp.GetActiveMachine()
	warp.cli = GetCLI(warp.User.Name, warp.FolderPath, warp.apiSet(), machine)
	go warp.listen()
	if warp.apiSet() {
		if !warp.apiValid() {
			warp.cli.apiSet = false
		}
		go warp.cli.Start()
	} else {
		go warp.cli.Start()
	}

	time.Sleep(500 * time.Millisecond)
	if warp.apiSet() {
		warp.prepCLI()
	}
	warp.listen()
}

func (warp *Warp) grabImage(avatar, fileFolder string) {
	bytes, err := warp.GetFile(avatar)
	if err != nil {
		return
	} else {
		fileName := strings.Split(avatar, "/")[len(strings.Split(avatar, "/"))-1]
		filePath := path.Join(fileFolder, fileName)
		if !util.FileExists(filePath) {
			util.StoreFile(bytes, filePath)
		}
	}
}

func (warp *Warp) prepCLI() {
	warp.updateUserInfo()
	warp.updateActivity()
	warp.updateBadges()
	warp.updateProLabsProgress()
	warp.updateActiveMachine()
	warp.updateRetiredMachines()
	warp.updateMachines()
}

func (warp *Warp) fetchTargetIP() {

}

func (warp *Warp) updateUserInfo() {
	userInfo, err := warp.GetUserInfo()
	if err != nil {
		// todo log here
		return
	}
	userBrief := UserBrief{
		ID:               userInfo.Info.ID,
		Name:             userInfo.Info.Name,
		Avatar:           userInfo.Info.Avatar,
		IsViP:            userInfo.Info.IsVip,
		CanAccessVIP:     userInfo.Info.CanAccessVIP,
		IsServerVIP:      userInfo.Info.IsServerVIP,
		ServerID:         userInfo.Info.ServerID,
		RankID:           userInfo.Info.RankID,
		Team:             userInfo.Info.Team,
		SubscriptionPlan: userInfo.Info.SubscriptionPlan,
		Identifier:       userInfo.Info.Identifier,
	}
	cu := CurrentUser{User: userBrief, Username: userInfo.Info.Name}
	warp.User = userBrief
	warp.sendUpdate(cu)
	warp.grabImage(userInfo.Info.Avatar, warp.FolderPath)
}

func (warp *Warp) updateProLabsProgress() {
	prolabs, err := warp.GetUserProgressProlabs()
	if err != nil {
		// todo log here
		return
	}
	warp.ProlabsProgress = prolabs.Profile.Prolabs
	prolabsProgress := ProlabsProgress{Progress: prolabs.Profile.Prolabs}
	warp.sendUpdate(prolabsProgress)
}

func (warp *Warp) updateActivity() {
	activity, err := warp.GetUserActivity()
	if err != nil {
		// todo log here
		return
	}
	al := ActivityList{Activity: activity.Profile.Activity[:10]}
	warp.Activity = activity.Profile.Activity[:10]
	warp.sendUpdate(al)
}

func (warp *Warp) updateBadges() {
	badges, err := warp.GetUserBadges()
	if err != nil {
		// todo log here
		return
	}
	bu := BadgesList{Badges: badges.Badges}
	warp.Badges = badges.Badges
	warp.sendUpdate(bu)
	for _, badge := range badges.Badges {
		warp.grabImage(badge.Icon, warp.FolderPath)
	}
}

func (warp *Warp) updateRetiredMachines() {
	retired, err := warp.ListRetiredMachines()
	if err != nil {
		// todo log here
		return
	}
	rl := RetiredMachinesList{Machines: retired.Data}
	warp.sendUpdate(rl)
	for _, machine := range retired.Data {
		warp.grabImage(machine.Avatar, warp.FolderPath)
	}
}

func (warp *Warp) updateMachines() {
	machines, err := warp.ListMachines()
	if err != nil {
		// todo log here
		return
	}
	ml := MachinesList{Machines: machines.Data}
	warp.sendUpdate(ml)
	for _, machine := range machines.Data {
		warp.grabImage(machine.Avatar, warp.FolderPath)
	}
}

func (warp *Warp) updateActiveMachine() {
	active, err := warp.GetActiveMachine()
	if err != nil {
		// todo log here
		return
	}
	am := ActiveMachine{active}
	warp.sendUpdate(am)
}

func (warp *Warp) sendUpdate(update Update) {
	toCLI <- update
}

func (warp *Warp) listen() {
	for {
		select {
		case update := <-fromCLI:
			switch update.GetType() {
			case NEW_API_KEY:
				warp.SetNewAPIKey(update.GetUpdate()["ApiKey"])
			}
		}
	}
}

func (warp *Warp) UpdateActiveMachine() {

}

func (warp *Warp) UpdateCurrentUser() {

}

func (warp *Warp) apiValid() bool {
	_, err := warp.GetUserInfo()
	return err == nil
}

func (warp *Warp) setData(key, value string) {
	if warp.data == nil {
		warp.data = make(map[string]string)
	}
	warp.data[key] = value
}

func loadHTB(path, fileFolder string) (*Warp, error) {
	if fileBytes, err := util.ReadFile(path); err != nil {
		return nil, htb.LOCAL_ERROR_FILE_NOT_FOUND
	} else {
		var warpClient Warp
		err = json.Unmarshal(fileBytes, &warpClient)
		warpClient.setClient()
		if warpClient.apiSet() {
			warpClient.setHeaders()
		}
		warpClient.FolderPath = fileFolder
		return &warpClient, err
	}
}

func (warp *Warp) setUser(user User) {
	warp.User = UserBrief{
		ID:               user.ID,
		Name:             user.Name,
		Avatar:           user.Avatar,
		IsViP:            user.IsVip,
		CanAccessVIP:     user.CanAccessVIP,
		IsServerVIP:      user.IsServerVIP,
		ServerID:         user.ServerID,
		RankID:           user.RankID,
		Team:             user.Team,
		SubscriptionPlan: user.SubscriptionPlan,
		Identifier:       user.Identifier,
	}

	err := warp.persist()
	if err != nil {
		panic(err)
	}

	if warp.data == nil {
		warp.data = make(map[string]string)
	}
	warp.data["user_id"] = fmt.Sprintf("%d", user.ID)
}

func (warp *Warp) setClient() {
	warp.client = &http.Client{}
}

func (warp *Warp) setHeaders() {
	if warp.headers == nil {
		warp.headers = make(map[string]string)
	}
	// Add the authorization header
	warp.headers = map[string]string{
		"Authorization": "Bearer " + warp.ApiKey,
		"Content-Type":  "application/json",
	}
}

func (warp *Warp) Do() (*http.Response, error) {
	return warp.client.Do(warp.req)
}

func (warp *Warp) setRequest(url urlLib.URL) {
	// Set the request
	warp.req, _ = http.NewRequest(http.MethodGet, url.String(), nil)
	for key, value := range warp.headers {
		warp.req.Header.Add(key, value)
	}
}

func (warp *Warp) apiSet() bool {
	return warp.ApiKey != ""
}

func (warp *Warp) vpnSet(vpn string) bool {
	return warp.Vpn[vpn] != ""
}

func (warp *Warp) userSet() bool {
	return warp.User.ID != 0
}

func (warp *Warp) ListVPNs() string {
	var vpnList string
	for k, _ := range warp.Vpn {
		vpnList += k + "\n"
	}
	return vpnList
}

func (warp *Warp) SetNewAPIKey(apiKey string) {
	warp.ApiKey = apiKey

	userInfoResponse, err := warp.GetUserInfo()
	if err != nil {
		warp.setUser(userInfoResponse.Info)
	} else {
		panic(err)
	}

	err = warp.persist()
	if err != nil {
		panic(err)
	}
}

func (warp *Warp) persist() error {
	if fileBytes, err := json.Marshal(warp); err == nil {
		err = util.WriteFile(warp.ConfPath, fileBytes)
		return err
	} else {
		return err
	}
}

func (warp *Warp) AddVPN(vpn htb.VPN, path string) error {
	if util.FileExists(path) {
		// If the Vpn map is nil create it
		if warp.Vpn == nil {
			warp.Vpn = make(map[string]string)
		}
		// If the Vpn already exists, update it, otherwise add it
		warp.Vpn[vpn.String()] = path
		return nil
	} else {
		return htb.LOCAL_ERROR_FILE_NOT_FOUND
	}
}

func (warp *Warp) RemoveVPN(vpn htb.VPN) {
	delete(warp.Vpn, vpn.String())
}
