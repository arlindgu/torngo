package id

import "torngo/internal/api"

type MarketIdItemmarketResponse struct {
	Itemmarket struct {
		Item struct {
			ID           int64  `json:"id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			AveragePrice int64  `json:"average_price"`
		} `json:"item"`
		Listings []struct {
			Price       int64 `json:"price"`
			Amount      int   `json:"amount"`
			ItemDetails struct {
				UID   int64 `json:"uid"`
				Stats struct {
					Damage   float64 `json:"damage"`
					Accuracy float64 `json:"accuracy"`
					Armor    float64 `json:"armor"`
					Quality  float64 `json:"quality"`
				} `json:"stats"`
				Bonuses []struct {
					ID          int    `json:"id"`
					Title       string `json:"title"`
					Description string `json:"description"`
					Value       int    `json:"value"`
				} `json:"bonuses"`
				Rarity string `json:"rarity"`
			} `json:"item_details,omitempty"`
		} `json:"listings"`
		CacheTimestamp int `json:"cache_timestamp"`
	} `json:"itemmarket"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type IdItemmarketBonus string

const (
	IdItemmarketBonusAny         = "any"
	IdItemmarketBonusDouble      = "double"
	IdItemmarketBonusYellow      = "yellow"
	IdItemmarketBonusOrange      = "orange"
	IdItemmarketBonusRed         = "red"
	IdItemmarketBonusAchilles    = "achilles"
	IdItemmarketBonusAssassinate = "assassinate"
	IdItemmarketBonusBackstab    = "backstab"
	IdItemmarketBonusBerserk     = "berserk"
	IdItemmarketBonusBleed       = "bleed"
	IdItemmarketBonusBlindfire   = "blindfire"
	IdItemmarketBonusBlindside   = "blindside"
	IdItemmarketBonusBloodlust   = "bloodlust"
	IdItemmarketBonusBurn        = "burn"
	IdItemmarketBonusComeback    = "comeback"
	IdItemmarketBonusConserve    = "conserve"
	IdItemmarketBonusCripple     = "cripple"
	IdItemmarketBonusCrusher     = "crusher"
	IdItemmarketBonusCupid       = "cupid"
	IdItemmarketBonusDeadeye     = "deadeye"
	IdItemmarketBonusDeadly      = "deadly"
	IdItemmarketBonusDemoralize  = "demoralize"
	IdItemmarketBonusDisarm      = "disarm"
	IdItemmarketBonusDoubleEdged = "double edged"
	IdItemmarketBonusDoubleTap   = "double tap"
	IdItemmarketBonusEmasculate  = "emasculate"
	IdItemmarketBonusEmpower     = "empower"
	IdItemmarketBonusEviscerate  = "eviscerate"
	IdItemmarketBonusExecute     = "execute"
	IdItemmarketBonusExpose      = "expose"
	IdItemmarketBonusFinale      = "finale"
	IdItemmarketBonusFocus       = "focus"
	IdItemmarketBonusFreeze      = "freeze"
	IdItemmarketBonusFrenzy      = "frenzy"
	IdItemmarketBonusFury        = "fury"
	IdItemmarketBonusGrace       = "grace"
	IdItemmarketBonusHazardous   = "hazardous"
	IdItemmarketBonusHomeRun     = "home run"
	IdItemmarketBonusIrradiate   = "irradiate"
	IdItemmarketBonusLacerate    = "lacerate"
	IdItemmarketBonusMotivation  = "motivation"
	IdItemmarketBonusParalyze    = "paralyze"
	IdItemmarketBonusParry       = "parry"
	IdItemmarketBonusPenetrate   = "penetrate"
	IdItemmarketBonusPlunder     = "plunder"
	IdItemmarketBonusPoison      = "poison"
	IdItemmarketBonusPowerful    = "powerful"
	IdItemmarketBonusProficience = "proficience"
	IdItemmarketBonusPuncture    = "puncture"
	IdItemmarketBonusQuicken     = "quicken"
	IdItemmarketBonusRage        = "rage"
	IdItemmarketBonusRevitalize  = "revitalize"
	IdItemmarketBonusRoshambo    = "roshambo"
	IdItemmarketBonusShock       = "shock"
	IdItemmarketBonusSleep       = "sleep"
	IdItemmarketBonusSlow        = "slow"
	IdItemmarketBonusSmash       = "smash"
	IdItemmarketBonusSmurf       = "smurf"
	IdItemmarketBonusSpecialist  = "specialist"
	IdItemmarketBonusSpray       = "spray"
	IdItemmarketBonusStorage     = "storage"
	IdItemmarketBonusStricken    = "stricken"
	IdItemmarketBonusStun        = "stun"
	IdItemmarketBonusSuppress    = "suppress"
	IdItemmarketBonusSureShot    = "sure shot"
	IdItemmarketBonusThrottle    = "throttle"
	IdItemmarketBonusToxin       = "toxin"
	IdItemmarketBonusWarlord     = "warlord"
	IdItemmarketBonusWeaken      = "weaken"
	IdItemmarketBonusWindUp      = "wind up"
	IdItemmarketBonusWither      = "wither"
)

type MarketIdItemmarketParams struct {
	ItemId    int64
	Bonus     IdItemmarketBonus
	Limit     api.ApiLimit
	Sort      api.ApiSort
	Offset    api.ApiOffset
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
