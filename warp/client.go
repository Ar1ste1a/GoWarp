package warp

import (
	"GoWarp/htb"
	"GoWarp/util"
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
	"os"
)

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
}

func GetWarp() (*Warp, error) {
	var (
		ghtb *Warp
		err  error
	)

	home, err := os.UserHomeDir()
	confPath := fmt.Sprintf("%s/.htb.conf", home)

	if err != nil {
		return nil, err
	}

	if util.HTBExists(confPath) {
		ghtb, err = loadHTB(confPath)
	} else {
		ghtb = &Warp{
			ApiKey:   "",
			client:   &http.Client{},
			req:      nil,
			headers:  map[string]string{},
			data:     map[string]string{},
			Vpn:      map[string]string{},
			ConfPath: fmt.Sprintf(confPath),
		}
	}

	if ghtb.apiSet() {
		userInfoResponse, err := ghtb.GetUserInfo()
		if !ghtb.userSet() {
			if err == nil {
				ghtb.setUser(userInfoResponse.Info)
			}
		} else {
			ghtb.setData("user_id", fmt.Sprintf("%d", userInfoResponse.Info.ServerID))
		}
	}

	return ghtb, err
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

func (warp *Warp) setRequest(url url2.URL) {
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
