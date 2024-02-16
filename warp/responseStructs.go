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

type UserBlood struct {
	Machines   []interface{} `json:"machines"`
	Challenges []interface{} `json:"challenges"`
}

type UserBloodsResponse struct {
	Profile struct {
		Bloods []UserBlood `json:"bloods"`
	} `json:"profile"`
}

type MachineOwnageChart struct {
	MachineOwns struct {
		Solved int `json:"solved"`
		Total  int `json:"total"`
	} `json:"machine_owns"`
	MachineAttackPaths struct {
		Reconnaissance struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Reconnaissance"`
		SoftwareOSExploitation struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Software & OS exploitation"`
		Authentication struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Authentication"`
		SecurityTools struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Security Tools"`
		Injections struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Injections"`
		PasswordReuse struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Password Reuse"`
		PasswordCracking struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Password Cracking"`
		WebSiteStructureDiscovery struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Web Site Structure Discovery"`
		CommonApplications struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"Common Applications"`
		SUDOExploitation struct {
			Solved        int     `json:"solved"`
			Total         int     `json:"total"`
			AvgUserSolved float64 `json:"avg_user_solved"`
		} `json:"SUDO Exploitation"`
	} `json:"machine_attack_paths"`
}

type MachineOwnageChartResponse struct {
	Profile MachineOwnageChart `json:"profile"`
}

type UserQueryResponse struct {
	Profile struct {
		ID           int    `json:"id"`
		SsoID        string `json:"sso_id"`
		Name         string `json:"name"`
		SystemOwns   int    `json:"system_owns"`
		UserOwn      int    `json:"user_owns"`
		UserBloods   int    `json:"user_bloods"`
		SystemBloods int    `json:"system_bloods"`
		Team         struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Ranking int    `json:"ranking"`
			Avatar  string `json:"avatar"`
		} `json:"team"`
		Respects            int    `json:"respects"`
		Rank                string `json:"rank"`
		RankID              int    `json:"rank_id"`
		CurrentRankProgress int    `json:"current_rank_progress"`
		NextRank            string `json:"next_rank"`
		NextRankPoints      string `json:"next_rank_points"`
		RankOwnerShip       int    `json:"rank_ownership"`
		RankRequirement     string `json:"rank_requirement"`
		Ranking             string `json:"ranking"`
		Avatar              string `json:"avatar"`
		TimeZone            string `json:"time_zone"`
		Points              int    `json:"points"`
		CountryName         string `json:"country_name"`
		CountryCode         string `json:"country_code"`
		UniversityName      string `json:"university_name"`
		Description         string `json:"description"`
		Github              string `json:"github"`
		LinkedIn            string `json:"linkedIn"`
		Twitter             string `json:"twitter"`
		Website             string `json:"website"`
	} `json:"profile"`
}

type Badge struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	DescriptionEn   string `json:"description_en"`
	Color           string `json:"color"`
	Icon            string `json:"icon"`
	BadgeCategoryID int    `json:"badge_category_id"`
	BadgableType    string `json:"badgable_type"`
	BadgableID      int    `json:"badgable_id"`
	Percentage      int    `json:"percentage"`
	Pivot           struct {
		UserID    int    `json:"user_id"`
		BadgeID   int    `json:"badge_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"pivot"`
}

type UserBadgesResponse struct {
	Badges []Badge `json:"badges"`
}

type Creator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AvatarThumb string `json:"avatar_thumb"`
}

type UserAvailability struct {
	Available bool   `json:"available"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type Data struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	AvatarURL            string           `json:"avatar_url"`
	CoverImageURL        string           `json:"cover_image_url"`
	Retired              bool             `json:"retired"`
	VIP                  bool             `json:"vip"`
	Creators             []Creator        `json:"creators"`
	EndgameMachinesCount int              `json:"endgame_machines_count"`
	EndgameFlagsCount    int              `json:"endgame_flags_count"`
	UserAvailability     UserAvailability `json:"user_availability"`
	New                  bool             `json:"new"`
}

type ListEndgameResponse struct {
	Status bool   `json:"status"`
	Data   []Data `json:"data"`
}

type FeedbackForChart struct {
	CounterCake      int `json:"counterCake"`
	CounterVeryEasy  int `json:"counterVeryEasy"`
	CounterEasy      int `json:"counterEasy"`
	CounterTooEasy   int `json:"counterTooEasy"`
	CounterMedium    int `json:"counterMedium"`
	CounterBitHard   int `json:"counterBitHard"`
	CounterHard      int `json:"counterHard"`
	CounterTooHard   int `json:"counterTooHard"`
	CounterExHard    int `json:"counterExHard"`
	CounterBrainFuck int `json:"counterBrainFuck"`
}

type PlayInfo struct {
	IsActive  interface{} `json:"isActive"`
	ExpiresAt interface{} `json:"expires_at"`
}

type RetiredData struct {
	ID                  int              `json:"id"`
	Avatar              string           `json:"avatar"`
	Name                string           `json:"name"`
	StaticPoints        int              `json:"static_points"`
	SpFlag              int              `json:"sp_flag"`
	Os                  string           `json:"os"`
	Points              int              `json:"points"`
	Star                float64          `json:"star"`
	Release             string           `json:"release"`
	EasyMonth           int              `json:"easy_month"`
	Poweroff            int              `json:"poweroff"`
	Free                bool             `json:"free"`
	Difficulty          int              `json:"difficulty"`
	DifficultyText      string           `json:"difficultyText"`
	UserOwnsCount       int              `json:"user_owns_count"`
	AuthUserInUserOwns  bool             `json:"authUserInUserOwns"`
	RootOwnsCount       int              `json:"root_owns_count"`
	AuthUserHasReviewed bool             `json:"authUserHasReviewed"`
	AuthUserInRootOwns  bool             `json:"authUserInRootOwns"`
	IsTodo              bool             `json:"isTodo"`
	IsCompetitive       bool             `json:"is_competitive"`
	Active              interface{}      `json:"active"`
	FeedbackForChart    FeedbackForChart `json:"feedbackForChart"`
	Ip                  interface{}      `json:"ip"`
	PlayInfo            PlayInfo         `json:"playInfo"`
	Labels              []interface{}    `json:"labels"`
	Recommended         int              `json:"recommended"`
}

type RetiredLink struct {
	First string      `json:"first"`
	Last  string      `json:"last"`
	Prev  interface{} `json:"prev"`
	Next  string      `json:"next"`
}

type Links struct {
	URL    interface{} `json:"url"`
	Label  string      `json:"label"`
	Active bool        `json:"active"`
}

type RetiredMeta struct {
	CurrentPage int     `json:"current_page"`
	From        int     `json:"from"`
	LastPage    int     `json:"last_page"`
	Links       []Links `json:"links"`
	Path        string  `json:"path"`
	PerPage     int     `json:"per_page"`
	To          int     `json:"to"`
	Total       int     `json:"total"`
}

type RetiredMachinesResponse struct {
	Data  []RetiredData `json:"data"`
	Links RetiredLink   `json:"links"`
	Meta  RetiredMeta   `json:"meta"`
}

type ActiveMachineResponse struct {
	Info struct {
		Id   int         `json:"id"`
		Name string      `json:"name"`
		Type interface{} `json:"type"`
		Ip   string      `json:"ip"`
	} `json:"info"`
}
