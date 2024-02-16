package warp

import (
	"GoWarp/htb"
	"GoWarp/util"
	"encoding/json"
	"fmt"
	"net/http"
	urlLib "net/url"
	"os"
)

// Kris eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiNGUzY2YxNDA3OTI2ZDQ0OWNiNDgzMjNmY2ViNjg5OWM3MWUwNzlkNTZjYWY0ZmExY2Q1YjFkODViMjgzOWExM2MxMTk0MWVjYjdlODI4ZjAiLCJpYXQiOjE3MDgwOTM2MTMuODE1OTAzLCJuYmYiOjE3MDgwOTM2MTMuODE1OTA1LCJleHAiOjE3MDgxODAwMTMuODEwODk0LCJzdWIiOiI5NTI5NjYiLCJzY29wZXMiOlsiMmZhIl19.VkG4JJ75T_38BKGHKb6kLm5wKZDZ_1rIji5vkECLL8DVR3lKfKLXXN9TrG7L1AR2N8tziZu_R92q8cQ_6t8Ue58D3s8GWlPP8LgvTNEfjy7jHZnd6I_XQKrWcRSFmFlSkxvmOKFAYIjfxYtiK0KY1FjPkC3OZkx-ZOaFeyv4F9qSkahiKHNjQXR1dnvhn1hiQKhQcvuDKP-QZ9pP3626hvGTew8w9LdklnbjE-gM7wxQiRRiPeGlykuhDRYdIdvcRmrC1FBDolBaJarpgxoStshWLh9vUdcrFvjFKl_nRknlM2Zmpk1e1roebZD2m9-rwOvezTv2Zit6jvR3NRk_lsNJX8mwUE46SoBXEkY7rneqRiuiebedMu0gfjRKeas5NMgqm9IjUBuShJT0mb5Xo5mNPPXsAphEbXV_TeKuRtZMGI3MYxkX0fXV3n37PKKbAoESAQDFP9D5yIy6LmzizYNhpmUcoxGLzyN101uVeH7R9uE0UnAiyGUQh5TQNd7ShSvE0Ja68FP1kdwaN35-N0G9NIODODDOYpYz3clMuCVweP7pHWTwLRnqpGGpGWiR0ZjzwhpIjmeAy7BdFctMXnpt3cX7pWbDbrUZATQVg7c5ljuomVzvZjyka52QicCP7HcHvnNgzvx359I2_rH7E0NZAy5bZqADZooIkDN7uBE

type Warp struct {
	ApiKey   string `json:"ApiKey"`
	client   *http.Client
	req      *http.Request
	headers  map[string]string
	data     map[string]string
	Vpn      map[string]string `json:"Vpn"`
	ConfPath string            `json:"confPath"`
	machines map[string]Machine
	prolabs  map[string]Machine
	User     UserBrief `json:"user"`
	update   chan Update
	cli      *CLI
}

func newWarp(confPath string) *Warp {
	ghtb := &Warp{
		ApiKey:   "",
		client:   &http.Client{},
		req:      nil,
		headers:  map[string]string{},
		data:     map[string]string{},
		Vpn:      map[string]string{},
		update:   make(chan Update, 5),
		ConfPath: fmt.Sprintf(confPath),
	}
	return ghtb
}

func GetWarpClient() (*Warp, error) {
	var (
		ghtb *Warp
		err  error
	)

	home, err := os.UserHomeDir()
	confFolder := fmt.Sprintf("%s/.htb", home)
	confPath := fmt.Sprintf("%s/htb.conf", confFolder)

	if err != nil {
		return nil, err
	}

	// If the config folder does not exist, create it and generate a new Warp Client
	if !util.DirExists(confFolder) {
		err = os.Mkdir(confFolder, 0755)
		ghtb = newWarp(confPath)
	} else {
		// If the config folder exists, check if the config file exists, if not generate a new Warp Client else load the existing one
		if util.HTBExists(confPath) {
			ghtb, err = loadHTB(confPath)
		} else {
			ghtb = newWarp(confPath)
		}
	}

	// If the API key is set
	if ghtb.apiSet() {
		userInfoResponse, err := ghtb.GetUserInfo()

		// If the user is not set, set the user
		if !ghtb.userSet() {
			if err == nil {
				ghtb.setUser(userInfoResponse.Info)
			}
		}

		// Set the user id and server id
		ghtb.setData("server_id", fmt.Sprintf("%d", userInfoResponse.Info.ServerID))
		ghtb.setData("user_id", fmt.Sprintf("%d", userInfoResponse.Info.ID))
	}

	return ghtb, err
}

func (warp *Warp) Start() {
	machine, _ := warp.GetActiveMachine()
	warp.cli = GetCLI(warp.User.Name, warp.apiSet(), machine, warp.update)
	go warp.listen()
	if warp.apiSet() {
		if !warp.apiValid() {
			warp.cli.apiSet = false
		}
		warp.cli.Start()
	} else {
		warp.cli.Start()
	}
}

func (warp *Warp) listen() {
	for {
		select {
		case update := <-warp.update:
			switch update.GetType() {
			case NEW_API_KEY:
				warp.SetNewAPIKey(update.GetUpdate()["ApiKey"])
			}
			fmt.Println(update)
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

func loadHTB(path string) (*Warp, error) {
	if fileBytes, err := util.ReadFile(path); err != nil {
		return nil, htb.LOCAL_ERROR_FILE_NOT_FOUND
	} else {
		var ghtb Warp
		err = json.Unmarshal(fileBytes, &ghtb)
		ghtb.setClient()
		if ghtb.apiSet() {
			ghtb.setHeaders()
		}
		return &ghtb, err
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
