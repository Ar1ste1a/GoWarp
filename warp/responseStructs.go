package warp

type Machine struct {
	ID             int     `json:"id"`
	Avatar         string  `json:"avatar"`
	Name           string  `json:"name"`
	StaticPoints   int     `json:"static_points"`
	SpFlag         int     `json:"sp_flag"`
	Os             string  `json:"os"`
	Points         int     `json:"points"`
	Star           float64 `json:"star"`
	Release        string  `json:"release"`
	EasyMonth      int     `json:"easy_month"`
	Poweroff       int     `json:"poweroff"`
	Free           bool    `json:"free"`
	Difficulty     int     `json:"difficulty"`
	DifficultyText string  `json:"difficultyText"`
	// Add other fields as needed
}

type Link struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Meta struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	LastPage    int    `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
}

type Track struct {
	ID       int `json:"id"`
	Complete int `json:"complete"`
}

type EnrolledTracksResponse []Track

type User struct {
	ID                         int    `json:"id"`
	Name                       string `json:"name"`
	Email                      string `json:"email"`
	Timezone                   string `json:"timezone"`
	IsVip                      bool   `json:"isVip"`
	IsModerator                bool   `json:"isModerator"`
	IsBGModerator              bool   `json:"isBGModerator"`
	IsChatBanned               bool   `json:"isChatBanned"`
	IsDedicatedVip             bool   `json:"isDedicatedVip"`
	CanAccessVIP               bool   `json:"canAccessVIP"`
	CanAccessDedilab           bool   `json:"canAccessDedilab"`
	IsServerVIP                bool   `json:"isServerVIP"`
	ServerID                   int    `json:"server_id"`
	Avatar                     string `json:"avatar"`
	BetaTester                 int    `json:"beta_tester"`
	RankID                     int    `json:"rank_id"`
	OnboardingCompleted        bool   `json:"onboarding_completed"`
	OnboardingTutorialComplete int    `json:"onboarding_tutorial_complete"`
	Verified                   bool   `json:"verified"`
	CanDeleteAvatar            bool   `json:"can_delete_avatar"`
	Team                       struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		AvatarThumbURL string `json:"avatar_thumb_url"`
	} `json:"team"`
	University        interface{} `json:"university"`
	Identifier        string      `json:"identifier"`
	HasTeamInvitation bool        `json:"hasTeamInvitation"`
	TwoFaEnabled      bool        `json:"TwoFaEnabled"`
	HasAppTokens      bool        `json:"hasAppTokens"`
	OptIn             int         `json:"opt_in"`
	IsSsoConnected    bool        `json:"is_sso_connected"`
	SubscriptionPlan  string      `json:"subscription_plan"`
	DunningExists     bool        `json:"dunning_exists"`
}
type UserInfoResponse struct {
	Info User `json:"info"`
}

type ListMachinesResponse struct {
	Data  []Machine `json:"data"`
	Links Link      `json:"links"`
	Meta  Meta      `json:"meta"`
}

type UserBrief struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	IsViP        bool   `json:"isVip"`
	CanAccessVIP bool   `json:"canAccessVIP"`
	IsServerVIP  bool   `json:"isServerVIP"`
	ServerID     int    `json:"server_id"`
	RankID       int    `json:"rank_id"`
	Team         struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		AvatarThumbURL string `json:"avatar_thumb_url"`
	}
	SubscriptionPlan string `json:"subscription_plan"`
	Identifier       string `json:"identifier"`
}

type UserActivityResponse struct {
	Profile struct {
		Activity []UserActivity `json:"activity"`
	} `json:"profile"`
}

type UserActivity struct {
	Date          string `json:"date"`
	DateDiff      string `json:"date_diff"`
	ObjectType    string `json:"object_type"`
	Type          string `json:"type"`
	FirstBlood    bool   `json:"first_blood"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Points        int    `json:"points"`
	MachineAvatar string `json:"machine_avatar"`
}

type Subscription struct {
	Name        string      `json:"name"`
	EndsAt      string      `json:"ends_at"`
	RenewsAt    string      `json:"renews_at"`
	RecurlyPlan string      `json:"recurly_plan"`
	PausedAt    string      `json:"paused_at"`
	CreatedAt   string      `json:"created_at"`
	ResumeAt    string      `json:"resume_at"`
	StripePlan  string      `json:"stripe_plan"`
	Items       interface{} `json:"items"`
}

type SubscriptionStatusResponse struct {
	Subscriptions []Subscription `json: "subscriptions"`
}
