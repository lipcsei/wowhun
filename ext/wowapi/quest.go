package wowapi

type GetQuestResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	Area  struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"area"`
	Description  string `json:"description"`
	Requirements struct {
		MinCharacterLevel int `json:"min_character_level"`
		MaxCharacterLevel int `json:"max_character_level"`
		Faction           struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"requirements"`
	Rewards struct {
		Experience  int `json:"experience"`
		Reputations []struct {
			Reward struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"reward"`
			Value int `json:"value"`
		} `json:"reputations"`
		Money struct {
			Value int `json:"value"`
			Units struct {
				Gold   int `json:"gold"`
				Silver int `json:"silver"`
				Copper int `json:"copper"`
			} `json:"units"`
		} `json:"money"`
	} `json:"rewards"`
}
