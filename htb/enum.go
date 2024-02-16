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
)

func (e ENDPOINT) Url(data map[string]string) (*url2.URL, error) {
	urlString := "https://labs.hackthebox.com/api/v4"

	switch e {
	case GET_CONNECTION_STATUS:
		urlString = fmt.Sprintf("%s/connection/status", urlString)
	case GET_ACTIVE_MACHINE:
		urlString = fmt.Sprintf("%s/machines/active", urlString)
	case GET_USER_INFO:
		urlString = fmt.Sprintf("%s/user/info", urlString)
	case GET_USER_SETTINGS:
		urlString = fmt.Sprintf("%s/user/settings", urlString)
	case GET_SUBSCRIPTION_STATUS:
		urlString = fmt.Sprintf("%s/user/subscriptions", urlString)
	case GET_SUBSCRIPTION_BALANCE:
		urlString = fmt.Sprintf("%s/user/subscriptions/balance", urlString)
	case GET_ENROLLED_TRACKS:
		urlString = fmt.Sprintf("%s/user/tracks", urlString)
	case GET_CHALLENGE_SUBMISSIONS:
		urlString = fmt.Sprintf("%s/submissions/challenges", urlString)
	case GET_USER_PROGRESS_MACHINES_BY_OS:
		urlString = fmt.Sprintf("%s/profile/progress/machines/os/%s", urlString, data["user_id"])
	case GET_USER_PROGRESS_CHALLENGES_BY_CATEGORY:
		urlString = fmt.Sprintf("%s/profile/progress/challenges/%s", urlString, data["user_id"])
	case GET_USER_PROGRESS_ENDGAME:
		urlString = fmt.Sprintf("%s/profile/progress/endgame/%s", urlString, data["user_id"])
	case GET_USER_PROGRESS_FORTRESS:
		urlString = fmt.Sprintf("%s/profile/progress/fortress/%s", urlString, data["user_id"])
	case GET_USER_PROGRESS_PROLABS:
		urlString = fmt.Sprintf("%s/profile/progress/prolabs/%s", urlString, data["user_id"])
	case GET_USER_ACTIVITY:
		urlString = fmt.Sprintf("%s/profile/activity/%s", urlString, data["user_id"])
	case GET_USER_BLOODS:
		urlString = fmt.Sprintf("%s/profile/bloods/%s", urlString, data["user_id"])
	case GET_USER_ACHIEVEMENTS_GRAPH:
		urlString = fmt.Sprintf("%s/profile/graph/%s/%s", urlString, data["progress_time"], data["user_id"]) // progress time can be one of 1Y, 6M, 3M, 1M, 1W
	case GET_MACHINE_OWNAGE_CHART_BY_ATTACK_PATH:
		urlString = fmt.Sprintf("%s/profile/chart/machines/attack/%s", urlString, data["user_id"])
	case GET_PROFILE_OVERVIEW:
		urlString = fmt.Sprintf("%s/profile/%s", urlString, data["user_id"])
	case GET_USER_BADGES:
		urlString = fmt.Sprintf("%s/profile/badges/%s", urlString, data["user_id"])
	case GET_VALIDATE_MACHINE_OWNAGE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/user/achievement/machine/%s/%s", urlString, data["user_id"], data["machine_id"])
	case GET_LIST_ENDGAMES:
		urlString = fmt.Sprintf("%s/endgames", urlString)
	case GET_ENDGAME_PROFILE:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/endgame/%s", urlString, data["endgame_id"])
	case GET_ENDGAMGE_FLAG_LIST:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/endgame/%s/flags", urlString, data["endgame_id"])
	case GET_ENDGAME_MACHINE_LIST:
		if _, ok := data["endgame_id"]; !ok {
			return nil, HTB_ERROR_ENDGAME_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/endgame/%s/machines", urlString, data["endgame_id"])
	case GET_PROLABS_LIST:
		urlString = fmt.Sprintf("%s/prolabs", urlString)
	case GET_PROLAB_INFO:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s", urlString, data["prolab_id"])
	case GET_PROLAB_OVERVIEW:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/overview", urlString, data["prolab_id"])
	case GET_PROLAB_MACHINES:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/machines", urlString, data["prolab_id"])
	case GET_PROLAB_FLAG_LIST:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/flags", urlString, data["prolab_id"])
	case GET_PROLAB_FAQ:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/faq", urlString, data["prolab_id"])
	case GET_PROLAB_REVIEWS:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/reviews_overview", urlString, data["prolab_id"])
	case GET_PROLAB_REVIEWS_PAGINATED:
		if _, ok := data["prolab_id"]; !ok {
			return nil, HTB_ERROR_PROLAB_ID_MISSING
		}
		if _, ok := data["page_num"]; !ok {
			return nil, HTB_ERROR_PAGE_NUMBER_MISSING
		}
		urlString = fmt.Sprintf("%s/prolabs/%s/reviews?page=%s", urlString, data["prolab_id"], data["page_num"])
	case GET_FORTRESS_LIST:
		urlString = fmt.Sprintf("%s/fortresses", urlString)
	case GET_FORTRESS_INFO:
		if _, ok := data["fortress_id"]; !ok {
			return nil, HTB_ERROR_FORTRESS_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/fortress/%s", urlString, data["fortress_id"])
	case GET_FORTRESS_FLAG_LIST:
		if _, ok := data["fortress_id"]; !ok {
			return nil, HTB_ERROR_FORTRESS_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/fortress/%s/flags", urlString, data["fortress_id"])
	case GET_STARTING_POINT_MACHINE_LIST:
		urlString = fmt.Sprintf("%s/sp/machines", urlString)
	case GET_BADGE_LIST_BY_CATEGORY:
		urlString = fmt.Sprintf("%s/badges", urlString)
	case GET_BADGE_LIST_BY_CATEGORY_ALTERNATE:
		urlString = fmt.Sprintf("%s/category/badges", urlString)
	case GET_DASHBOARD_MACHINE_CHALLENGE_COUNTS:
		urlString = fmt.Sprintf("%s/content/stats", urlString)
	case GET_HTB_SERVER_LIST:
		urlString = fmt.Sprintf("%s/lab/list", urlString)
	case GET_DASHBOARD:
		urlString = fmt.Sprintf("%s/user/dashboard", urlString)
	case GET_RANKINGS_BEST_BY_COUNTRY_FOR_YEAR:
		urlString = fmt.Sprintf("%s/rankings/country/best?period=1Y", urlString)
	case GET_RANKINGS_BEST_BY_COUNTRY:
		urlString = fmt.Sprintf("%s/rankings/country/ranking_bracket", urlString)
	case GET_LIST_ACTIVE_MACHINES:
		urlString = fmt.Sprintf("%s/machine/list", urlString)
	case GET_LIST_RETIRED_MACHINES:
		urlString = fmt.Sprintf("%s/machine/list/retired", urlString)
	case GET_LIST_TODO_MACHINES:
		urlString = fmt.Sprintf("%s/machine/list/todo", urlString)
	case GET_MACHINE_PROFILE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/profile/%s", urlString, data["machine_id"])
	case GET_MACHINE_ACTIVITY:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/activity/%s", urlString, data["machine_id"])
	case GET_MACHINE_TOP_25_OWNERS:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/owns/top/%s", urlString, data["machine_id"])
	case GET_MACHINE_TAGS_SUB_LANGUAGE:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/tags/%s", urlString, data["machine_id"])
	case GET_MACHINE_LIST_PAGINATED:
		urlString = fmt.Sprintf("%s/machine/paginated", urlString)
	case GET_MACHINE_LIST_TAGS:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/tags/%s", urlString, data["machine_id"])
	case GET_MACHINE_LIST_WALKTHROUGH:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/walkthroughs/%s", urlString, data["machine_id"])
	case GET_MACHINE_OFFICIAL_WRITEUP:
		if _, ok := data["machine_id"]; !ok {
			return nil, HTB_ERROR_MACHINE_ID_MISSING
		}
		urlString = fmt.Sprintf("%s/machine/writeup/%s", urlString, data["machine_id"])
	default:
		// do nothing
		return nil, HTB_ERROR_UNKNOWN
	}

	url, err := url2.Parse(urlString)
	if err != nil {
		return nil, err
	}
	return url, nil
}
