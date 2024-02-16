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
			fmt.Printf("List Machines Error: %s", err)
		} else {
			// Make the request
			warp.setRequest(*url)

			// Send the request
			resp, err = warp.Do()
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return lmr, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
				}

				err = json.Unmarshal(bodyBytes, &lmr)

				return lmr, err
			}
		}
	}

	fmt.Println("API Key Not Set")
	return lmr, htb.LOCAL_API_KEY_UNSET
}

// Empty Body
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
func (warp *Warp) GetActiveMachine() {
	if warp.apiSet() {
		url, err := htb.GET_ACTIVE_MACHINE.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Active Machine Error: %s", err)
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

func (warp *Warp) GetUserInfo() (UserInfoResponse, error) {
	var uio UserInfoResponse

	if warp.apiSet() {
		url, err := htb.GET_USER_INFO.Url(warp.data)
		if err != nil {
			fmt.Printf("Get User Info Error: %s", err)
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
				return UserInfoResponse{}, err
			} else {
				defer resp.Body.Close()

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error Parsing Body: %s", err)
					return UserInfoResponse{}, err
				}

				err = json.Unmarshal(bodyBytes, &uio)
				fmt.Printf("\n\nStatus: %s", resp.Status)
				fmt.Printf("\n\nBody: %s", string(bodyBytes))
				return uio, err
			}
		}
	}

	fmt.Println("API Key Not Set")
	return uio, htb.LOCAL_API_KEY_UNSET
}

func (warp *Warp) GetUserSettings() {
	if warp.apiSet() {
		url, err := htb.GET_USER_SETTINGS.Url(warp.data)
		if err != nil {
			fmt.Printf("Get Enrolled Tracks Error: %s", err)
			//return etr, err
		} else {
			fmt.Printf("URL: %s", url)

			// Make the request
			warp.setRequest(*url)

			// Log the response
			resp, err := warp.client.Do(warp.req)
			if err != nil {
				fmt.Printf("Error Performing Request: %s", err)
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
	//return etr, htb.LOCAL_API_KEY_UNSET

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
	return etr, htb.LOCAL_API_KEY_UNSET
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

func (warp *Warp) GetUserBloods() {

}

func (warp *Warp) GetUserAchievementsGraph() {

}

func (warp *Warp) GetMachineOwnageChartByAttackPath() {

}

func (warp *Warp) GetProfileOverview() {

}

func (warp *Warp) GetUserBadges() {

}

func (warp *Warp) GetValidateMachineOwnage() {

}

func (warp *Warp) GetListEndgames() {

}

func (warp *Warp) GetEndgameProfile() {

}

func (warp *Warp) GetEndgamgeFlagList() {

}

func (warp *Warp) GetEndgameMachineList() {

}
