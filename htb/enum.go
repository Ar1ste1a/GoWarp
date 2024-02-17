package htb

import (
	"fmt"
	url2 "net/url"
)

type VPN uint

const (
	ACADEMY_VPN VPN = iota
	LABS_VPN
	PROLABS_VPN
	ENDGAME_VPN
	FORTRESS_VPN
)

func (v VPN) String() string {
	switch v {
	case ACADEMY_VPN:
		return "academy"
	case LABS_VPN:
		return "labs"
	case PROLABS_VPN:
		return "prolabs"
	case ENDGAME_VPN:
		return "endgame"
	case FORTRESS_VPN:
		return "fortress"
	default:
		return "unknown"
	}
}

func GetVPNFromString(vpn string) VPN {
	switch vpn {
	case "academy":
		return ACADEMY_VPN
	case "labs":
		return LABS_VPN
	case "prolabs":
		return PROLABS_VPN
	case "endgame":
		return ENDGAME_VPN
	case "fortress":
		return FORTRESS_VPN
	default:
		return 0
	}
}

type ENDPOINT uint

const (
	GET_CONNECTION_STATUS ENDPOINT = iota
	GET_ACTIVE_MACHINE
	GET_USER_INFO
	GET_USER_SETTINGS
	GET_SUBSCRIPTION_STATUS
	GET_SUBSCRIPTION_BALANCE
	GET_ENROLLED_TRACKS
	GET_CHALLENGE_SUBMISSIONS
	GET_USER_PROGRESS_MACHINES_BY_OS
	GET_USER_PROGRESS_CHALLENGES_BY_CATEGORY
	GET_USER_PROGRESS_ENDGAME
	GET_USER_PROGRESS_FORTRESS
	GET_USER_PROGRESS_PROLABS
	GET_USER_ACTIVITY
	GET_USER_BLOODS
	GET_USER_ACHIEVEMENTS_GRAPH
	GET_MACHINE_OWNAGE_CHART_BY_ATTACK_PATH
	GET_PROFILE_OVERVIEW
	GET_USER_BADGES
	GET_VALIDATE_MACHINE_OWNAGE
	GET_LIST_ENDGAMES
	GET_ENDGAME_PROFILE
	GET_ENDGAMGE_FLAG_LIST
	GET_ENDGAME_MACHINE_LIST
	GET_PROLABS_LIST
	GET_PROLAB_INFO
	GET_PROLAB_OVERVIEW
	GET_PROLAB_MACHINES
	GET_PROLAB_FLAG_LIST
	GET_PROLAB_FAQ
	GET_PROLAB_REVIEWS
	GET_PROLAB_REVIEWS_PAGINATED
	GET_FORTRESS_LIST
	GET_FORTRESS_INFO
	GET_FORTRESS_FLAG_LIST
	GET_STARTING_POINT_MACHINE_LIST
	GET_BADGE_LIST_BY_CATEGORY
	GET_BADGE_LIST_BY_CATEGORY_ALTERNATE
	GET_DASHBOARD_MACHINE_CHALLENGE_COUNTS
	GET_HTB_SERVER_LIST
	GET_DASHBOARD
	GET_RANKINGS_BEST_BY_COUNTRY_FOR_YEAR
	GET_RANKINGS_BEST_BY_COUNTRY
	GET_LIST_ACTIVE_MACHINES
	GET_LIST_RETIRED_MACHINES
	GET_LIST_TODO_MACHINES
	GET_MACHINE_PROFILE
	GET_MACHINE_ACTIVITY
	GET_MACHINE_TOP_25_OWNERS
	GET_MACHINE_TAGS_SUB_LANGUAGE
	GET_MACHINE_LIST_TAGS
	GET_MACHINE_LIST_WALKTHROUGH
	GET_MACHINE_OFFICIAL_WRITEUP
	GET_MACHINE_LIST_PAGINATED
	VPN_REQUEST_CONFIG
	VPN_SWITCH_SERVER
	GET_MACHINE_AVATAR_FILE
	GET_USER_AVATAR_FILE
)

func (e ENDPOINT) Url(data map[string]string) (*url2.URL, error) {
	var (
		labsString  = "https://labs.hackthebox.com/api/v4"
		htbString   = "https://www.hackthebox.com/api/v4"
		htbStandard = "https://www.hackthebox.com"
		lab         = false
		htb         = false
		standard    = false
	)

	switch e {
	case GET_MACHINE_AVATAR_FILE:
		htbStandard = fmt.Sprintf("%s%s", htbStandard, data["avatar"])
		standard = true
	case GET_USER_AVATAR_FILE:
		htbString = fmt.Sprintf("%s%s", htbStandard, data["avatar"])
		standard = true
	case VPN_REQUEST_CONFIG:
		labsString = fmt.Sprintf("%s/\"https://labs.hackthebox.com/api/v4/access/ovpnfile/%s%s", labsString, data["vpn_id"], data["tcp"]) // tcp /0 udp /1
		lab = true
	case VPN_SWITCH_SERVER:
		labsString = fmt.Sprintf("%s/connections/servers/switch/%s", labsString, data["vpn_id"])
		lab = true
	case GET_CONNECTION_STATUS:
		labsString = fmt.Sprintf("%s/connection/status", labsString)
		lab = true
	case GET_ACTIVE_MACHINE:
		labsString = fmt.Sprintf("%s/machine/active", labsString)
		lab = true
	case GET_USER_INFO:
		labsString = fmt.Sprintf("%s/user/info", labsString)
		lab = true
	case GET_USER_SETTINGS:
		labsString = fmt.Sprintf("%s/user/settings", labsString)
		lab = true
	case GET_SUBSCRIPTION_STATUS:
		labsString = fmt.Sprintf("%s/user/subscriptions", labsString)
		lab = true
	case GET_SUBSCRIPTION_BALANCE:
		labsString = fmt.Sprintf("%s/user/subscriptions/balance", labsString)
		lab = true
	case GET_ENROLLED_TRACKS:
		labsString = fmt.Sprintf("%s/user/tracks", labsString)
		lab = true
	case GET_CHALLENGE_SUBMISSIONS:
		labsString = fmt.Sprintf("%s/submissions/challenges", labsString)
		lab = true
	case GET_USER_PROGRESS_MACHINES_BY_OS:
		labsString = fmt.Sprintf("%s/profile/progress/machines/os/%s", labsString, data["user_id"])
		lab = true
	case GET_USER_PROGRESS_CHALLENGES_BY_CATEGORY:
		labsString = fmt.Sprintf("%s/profile/progress/challenges/%s", labsString, data["user_id"])
		lab = true
	case GET_USER_PROGRESS_ENDGAME:
		labsString = fmt.Sprintf("%s/profile/progress/endgame/%s", labsString, data["user_id"])
		lab = true
	case GET_USER_PROGRESS_FORTRESS:
		labsString = fmt.Sprintf("%s/profile/progress/fortress/%s", labsString, data["user_id"])
		lab = true
	case GET_USER_PROGRESS_PROLABS:
		labsString = fmt.Sprintf("%s/profile/progress/prolab/%s", htbString, data["user_id"])
		lab = true
	case GET_USER_ACTIVITY:
		labsString = fmt.Sprintf("%s/profile/activity/%s", labsString, data["server_id"])
		lab = true
	case GET_USER_BLOODS:
		labsString = fmt.Sprintf("%s/profile/bloods/%s", labsString, data["user_id"])
		lab = true
	case GET_USER_ACHIEVEMENTS_GRAPH:
		labsString = fmt.Sprintf("%s/profile/graph/%s/%s", labsString, data["progress_time"], data["user_id"]) // progress time can be one of 1Y, 6M, 3M, 1M, 1W
		lab = true
	case GET_MACHINE_OWNAGE_CHART_BY_ATTACK_PATH:
		labsString = fmt.Sprintf("%s/profile/chart/machines/attack/%s", labsString, data["user_id"])
		lab = true
	case GET_PROFILE_OVERVIEW:
		labsString = fmt.Sprintf("%s/user/profile/basic/%s", htbString, data["query_id"])
		lab = true
	case GET_USER_BADGES:
		labsString = fmt.Sprintf("%s/profile/badges/%s", htbString, data["user_id"])
		lab = true
	case GET_VALIDATE_MACHINE_OWNAGE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/user/achievement/machine/%s/%s", labsString, data["user_id"], data["machine_id"])
		lab = true
	case GET_LIST_ENDGAMES:
		labsString = fmt.Sprintf("%s/endgames", labsString)
		lab = true
	case GET_ENDGAME_PROFILE:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/endgame/%s", labsString, data["endgame_id"])
		lab = true
	case GET_ENDGAMGE_FLAG_LIST:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/endgame/%s/flags", labsString, data["endgame_id"])
		lab = true
	case GET_ENDGAME_MACHINE_LIST:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/endgame/%s/machines", labsString, data["endgame_id"])
		lab = true
	case GET_PROLABS_LIST:
		labsString = fmt.Sprintf("%s/prolabs", labsString)
		lab = true
	case GET_PROLAB_INFO:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_OVERVIEW:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/overview", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_MACHINES:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/machines", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_FLAG_LIST:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/flags", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_FAQ:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/faq", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_REVIEWS:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/reviews_overview", labsString, data["prolab_id"])
		lab = true
	case GET_PROLAB_REVIEWS_PAGINATED:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		if _, ok := data["page_num"]; !ok {
			return nil, HTB_ERROR_PAGE_NUMBER_MISSING
		}
		labsString = fmt.Sprintf("%s/prolabs/%s/reviews?page=%s", labsString, data["prolab_id"], data["page_num"])
		lab = true
	case GET_FORTRESS_LIST:
		labsString = fmt.Sprintf("%s/fortresses", labsString)
		lab = true
	case GET_FORTRESS_INFO:
		if _, ok := data["fortress_id"]; !ok {
			return nil, HTB_ERROR_FORTRESS_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/fortress/%s", labsString, data["fortress_id"])
		lab = true
	case GET_FORTRESS_FLAG_LIST:
		if _, ok := data["fortress_id"]; !ok {
			return nil, HTB_ERROR_FORTRESS_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/fortress/%s/flags", labsString, data["fortress_id"])
		lab = true
	case GET_STARTING_POINT_MACHINE_LIST:
		labsString = fmt.Sprintf("%s/sp/machines", labsString)
		lab = true
	case GET_BADGE_LIST_BY_CATEGORY:
		labsString = fmt.Sprintf("%s/badges", labsString)
		lab = true
	case GET_BADGE_LIST_BY_CATEGORY_ALTERNATE:
		labsString = fmt.Sprintf("%s/category/badges", labsString)
		lab = true
	case GET_DASHBOARD_MACHINE_CHALLENGE_COUNTS:
		labsString = fmt.Sprintf("%s/content/stats", labsString)
		lab = true
	case GET_HTB_SERVER_LIST:
		labsString = fmt.Sprintf("%s/lab/list", labsString)
		lab = true
	case GET_DASHBOARD:
		labsString = fmt.Sprintf("%s/user/dashboard", labsString)
		lab = true
	case GET_RANKINGS_BEST_BY_COUNTRY_FOR_YEAR:
		labsString = fmt.Sprintf("%s/rankings/country/best?period=1Y", labsString)
		lab = true
	case GET_RANKINGS_BEST_BY_COUNTRY:
		labsString = fmt.Sprintf("%s/rankings/country/ranking_bracket", labsString)
		lab = true
	case GET_LIST_ACTIVE_MACHINES:
		labsString = fmt.Sprintf("%s/machine/list", labsString)
		lab = true
	case GET_LIST_RETIRED_MACHINES:
		labsString = fmt.Sprintf("%s/machine/list/retired", labsString)
		lab = true
	case GET_LIST_TODO_MACHINES:
		labsString = fmt.Sprintf("%s/machine/list/todo", labsString)
		lab = true
	case GET_MACHINE_PROFILE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/profile/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_ACTIVITY:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/activity/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_TOP_25_OWNERS:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/owns/top/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_TAGS_SUB_LANGUAGE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/tags/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_LIST_PAGINATED:
		labsString = fmt.Sprintf("%s/machine/paginated", labsString)
		lab = true
	case GET_MACHINE_LIST_TAGS:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/tags/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_LIST_WALKTHROUGH:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/walkthroughs/%s", labsString, data["machine_id"])
		lab = true
	case GET_MACHINE_OFFICIAL_WRITEUP:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		labsString = fmt.Sprintf("%s/machine/writeup/%s", labsString, data["machine_id"])
		lab = true
	default:
		// do nothing
		return nil, HTB_ERROR_UNKNOWN
	}

	var url *url2.URL
	var err error

	if lab {
		url, err = url2.Parse(labsString)
	} else if htb {
		url, err = url2.Parse(htbString)
	} else if standard {
		url, err = url2.Parse(htbStandard)
	}

	if err != nil {
		return nil, err
	}
	return url, nil
}
