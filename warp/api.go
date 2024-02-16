package warp

import (
	"GoWarp/htb"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
)

func (warp *Warp) ListMachines() (ListMachinesResponse, error) {
	var (
		err  error
		lmr  ListMachinesResponse
		url  *url2.URL
		resp *http.Response
	)

	if warp.apiSet() {
		url, err = htb.GET_MACHINE_LIST_PAGINATED.Url(warp.data)
		if err != nil {
			return lmr, err
		} else {
			// Make the request
			warp.setRequest(*url)

			// Send the request
			resp, err = warp.Do()
			if err != nil {
				return lmr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return lmr, err
				}

				err = json.Unmarshal(bodyBytes, &lmr)

				return lmr, err
			}
		}
	}

	fmt.Println("API Key Not Set")
	return lmr, htb.LOCAL_ERROR_API_KEY_UNSET
}

// Empty Body TODO
func (warp *Warp) GetConnectionStatus() {
	if warp.apiSet() {
		url, err := htb.GET_CONNECTION_STATUS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Connection Status Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")
	}
}

// Does not exist
func (warp *Warp) GetActiveMachine() (Machine, error) {
	var m Machine

	if warp.apiSet() {
		url, err := htb.GET_ACTIVE_MACHINE.Url(warp.data)
		if err != nil {
			return m, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				return m, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return m, err
				}
				err = json.Unmarshal(bodyBytes, &m)
				return m, err
			}
		}
	} else {
		return m, htb.LOCAL_ERROR_API_KEY_UNSET
	}
}

func (warp *Warp) GetUserInfo() (UserInfoResponse, error) {
	var uio UserInfoResponse

	if warp.apiSet() {
		url, err := htb.GET_USER_INFO.Url(warp.data)
		if err != nil {
			return uio, err
		} else {
			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				return uio, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return uio, err
				}

				err = json.Unmarshal(bodyBytes, &uio)
				return uio, err
			}
		}
	}

	return uio, htb.LOCAL_ERROR_API_KEY_UNSET
}

func TestAPIKey(key string) error {

	url := "https://labs.hackthebox.com/api/v4/machine/list/retired/paginated"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return htb.HTB_ERROR_INVALID_API_KEY
	} else {
		return nil
	}
}

// Works but cannot find use case
func (warp *Warp) GetUserSettings() {
	if warp.apiSet() {
		url, err := htb.GET_USER_SETTINGS.Url(warp.data)
		if err != nil {
			//return etr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				//return etr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				//err = json.Unmarshal(bodyBytes, &etr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				//return etr, err
			}

		}
	}

	//fmt.Println("API Key Not Set")
	//return etr, htb.LOCAL_ERROR_API_KEY_UNSET

}

func (warp *Warp) GetSubscriptionStatus() (SubscriptionStatusResponse, error) {
	var (
		ssr SubscriptionStatusResponse
		err error
	)

	if warp.apiSet() {
		url, err := htb.GET_SUBSCRIPTION_STATUS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Enrolled Tracks Error: %s", err)
			return ssr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return ssr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &ssr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return ssr, err
			}

		}
	}
	fmt.Printf("API Key Not Set")
	return ssr, err
}

func (warp *Warp) GetSubscriptionBalance() {

}

func (warp *Warp) GetEnrolledTracks() (EnrolledTracksResponse, error) {
	var (
		//err error
		etr EnrolledTracksResponse
		//url *url2.URL
	)

	if warp.apiSet() {
		url, err := htb.GET_ENROLLED_TRACKS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Enrolled Tracks Error: %s", err)
			return etr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return etr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &etr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return etr, err
			}

		}
	}

	fmt.Println("API Key Not Set")
	return etr, htb.LOCAL_ERROR_API_KEY_UNSET
}

// This page does not exist
func (warp *Warp) GetChallengeSubmissions() {
	if warp.apiSet() {
		url, err := htb.GET_CHALLENGE_SUBMISSIONS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Challenge Submissions Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")

	}
}

// This page does not exist
func (warp *Warp) GetUserProgressMachinesByOS() {
	if warp.apiSet() {
		url, err := htb.GET_USER_PROGRESS_MACHINES_BY_OS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Progress Machines By OS Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")

	}
}

// 404
func (warp *Warp) GetUserProgressChallengesByCategory() {
	if warp.apiSet() {
		url, err := htb.GET_USER_PROGRESS_CHALLENGES_BY_CATEGORY.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Progress Challenges By Category Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")

	}
}

// 404
func (warp *Warp) GetUserProgressEndgame() {
	if warp.apiSet() {
		url, err := htb.GET_USER_PROGRESS_ENDGAME.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Progress Endgame Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")

	}
}

func (warp *Warp) GetUserProgressFortress() {

}

// 404
func (warp *Warp) GetUserProgressProlabs() {
	if warp.apiSet() {
		url, err := htb.GET_USER_PROGRESS_PROLABS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Progress Prolabs Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")
	}
}

func (warp *Warp) GetUserActivity() (UserActivityResponse, error) {
	var (
		uar UserActivityResponse
		err error
	)

	if warp.apiSet() {
		url, err := htb.GET_USER_ACTIVITY.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Activity Error: %s", err)
			return uar, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return uar, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &uar)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))

				return uar, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return uar, err
	}
}

func (warp *Warp) GetUserBloods() (UserBloodsResponse, error) {
	var (
		ubr UserBloodsResponse
		err error
	)

	if warp.apiSet() {
		url, err := htb.GET_USER_BLOODS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Bloods Error: %s", err)
			return ubr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return ubr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &ubr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return ubr, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return ubr, err
	}
}

// 404
func (warp *Warp) GetUserAchievementsGraph() {
	if warp.apiSet() {
		url, err := htb.GET_USER_ACHIEVEMENTS_GRAPH.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Achievements Graph Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")
	}
}

func (warp *Warp) GetMachineOwnageChartByAttackPath() (MachineOwnageChartResponse, error) {
	var (
		mocr MachineOwnageChartResponse
		err  error
	)

	if warp.apiSet() {
		url, err := htb.GET_MACHINE_OWNAGE_CHART_BY_ATTACK_PATH.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Machine Ownage Chart By Attack Path Error: %s", err)
			return mocr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return mocr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &mocr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return mocr, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return mocr, err
	}
}

func (warp *Warp) GetProfileOverview(queryId int) (UserQueryResponse, error) {
	var (
		uqr UserQueryResponse
		err error
	)

	if warp.apiSet() {
		warp.setData("query_id", fmt.Sprintf("%d", queryId))
		url, err := htb.GET_PROFILE_OVERVIEW.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Profile Overview Error: %s", err)
			return uqr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return uqr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &uqr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return uqr, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return uqr, err
	}
}

func (warp *Warp) GetUserBadges() (UserBadgesResponse, error) {
	var (
		ubr UserBadgesResponse
		err error
	)

	if warp.apiSet() {
		url, err := htb.GET_USER_BADGES.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Badges Error: %s", err)
			return ubr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return ubr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &ubr)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return ubr, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return ubr, err
	}
}

func (warp *Warp) GetValidateMachineOwnage(machineID int) {
	if warp.apiSet() {
		warp.setData("machine_id", fmt.Sprintf("%d", machineID))
		url, err := htb.GET_VALIDATE_MACHINE_OWNAGE.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Validate Machine Ownage Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
			}
		}
	} else {
		fmt.Println("API Key Not Set")
	}
}

func (warp *Warp) GetListEndgames() (ListEndgameResponse, error) {
	var (
		ler ListEndgameResponse
		err error
	)

	if warp.apiSet() {
		url, err := htb.GET_LIST_ENDGAMES.Url(warp.data)
		if err != nil {
			fmt.Printf("Get List Endgames Error: %s", err)
			return ler, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return ler, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}
				err = json.Unmarshal(bodyBytes, &ler)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return ler, err
			}
		}
	} else {
		fmt.Println("API Key Not Set")
		return ler, err
	}
}

// Works but cant find ID without supplying endgameId 404 with it
func (warp *Warp) GetEndgameProfile(endgameId int) {
	if warp.apiSet() {
		warp.setData("endgame_id", fmt.Sprintf("%d", endgameId))
		url, err := htb.GET_ENDGAME_PROFILE.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Endgame Profile Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			resp, err := warp.client.Do(warp.req)

			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error Parsing Body: %s", err)
			}
			//err = json.Unmarshal(bodyBytes, &ler)
			fmt.Printf("\n\nStatus: %s", resp.Status)
			fmt.Printf("\n\nBody: %s", string(bodyBytes))
		}
	} else {
		fmt.Println("API Key Not Set")
	}
}

func (warp *Warp) GetEndgameFlagList() {

}

func (warp *Warp) GetEndgameMachineList() {

}

func (warp *Warp) GetProlabsList() {

}

func (warp *Warp) GetProlabInfo() {

}

func (warp *Warp) GetProlabOverview() {

}

func (warp *Warp) Test() {
	url, _ := url2.Parse("https://labs.hackthebox.com/api/v4/sp/tier/0")

	warp.setRequest(*url)
	resp, _ := warp.Do()
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("\n\nStatus: %s", resp.Status)
	fmt.Printf("\n\nBody: %s", string(bodyBytes))
}

func (warp *Warp) ListRetiredMachines() (RetiredMachinesResponse, error) {
	var rmr RetiredMachinesResponse

	url, _ := url2.Parse("https://labs.hackthebox.com/api/v4/machine/list/retired/paginated")

	warp.setRequest(*url)
	resp, _ := warp.Do()
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(bodyBytes, &rmr)
	return rmr, err
}
