package id

type UserIdPersonalstatsResponse struct {
	Personalstats struct {
		Attacking struct {
			Attacks struct {
				Won       int `json:"won"`
				Lost      int `json:"lost"`
				Stalemate int `json:"stalemate"`
				Assist    int `json:"assist"`
				Stealth   int `json:"stealth"`
			} `json:"attacks"`
			Defends struct {
				Won       int `json:"won"`
				Lost      int `json:"lost"`
				Stalemate int `json:"stalemate"`
				Total     int `json:"total"`
			} `json:"defends"`
			Elo                int `json:"elo"`
			UnarmoredWins      int `json:"unarmored_wins"`
			HighestLevelBeaten int `json:"highest_level_beaten"`
			Escapes            struct {
				Player int `json:"player"`
				Foes   int `json:"foes"`
			} `json:"escapes"`
			Killstreak struct {
				Best    int `json:"best"`
				Current int `json:"current"`
			} `json:"killstreak"`
			Hits struct {
				Success     int `json:"success"`
				Miss        int `json:"miss"`
				Critical    int `json:"critical"`
				OneHitKills int `json:"one_hit_kills"`
			} `json:"hits"`
			Damage struct {
				Total int64 `json:"total"`
				Best  int   `json:"best"`
			} `json:"damage"`
			Networth struct {
				MoneyMugged int64 `json:"money_mugged"`
				LargestMug  int64 `json:"largest_mug"`
				ItemsLooted int   `json:"items_looted"`
			} `json:"networth"`
			Ammunition struct {
				Total       int `json:"total"`
				Special     int `json:"special"`
				HollowPoint int `json:"hollow_point"`
				Tracer      int `json:"tracer"`
				Piercing    int `json:"piercing"`
				Incendiary  int `json:"incendiary"`
			} `json:"ammunition"`
			Faction struct {
				Respect       int `json:"respect"`
				Retaliations  int `json:"retaliations"`
				RankedWarHits int `json:"ranked_war_hits"`
				RaidHits      int `json:"raid_hits"`
				Territory     struct {
					WallJoins  int `json:"wall_joins"`
					WallClears int `json:"wall_clears"`
					WallTime   int `json:"wall_time"`
				} `json:"territory"`
			} `json:"faction"`
		} `json:"attacking"`
		BattleStats struct {
			Strength  int64 `json:"strength"`
			Defense   int64 `json:"defense"`
			Speed     int64 `json:"speed"`
			Dexterity int64 `json:"dexterity"`
			Total     int64 `json:"total"`
		} `json:"battle_stats"`
		Jobs struct {
			JobPointsUsed  int `json:"job_points_used"`
			TrainsReceived int `json:"trains_received"`
			Stats          struct {
				Manual       int `json:"manual"`
				Intelligence int `json:"intelligence"`
				Endurance    int `json:"endurance"`
				Total        int `json:"total"`
			} `json:"stats"`
		} `json:"jobs"`
		Trading struct {
			Items struct {
				Bought struct {
					Market int `json:"market"`
					Shops  int `json:"shops"`
				} `json:"bought"`
				Auctions struct {
					Won  int `json:"won"`
					Sold int `json:"sold"`
				} `json:"auctions"`
				Sent int `json:"sent"`
			} `json:"items"`
			Trades int `json:"trades"`
			Points struct {
				Bought int `json:"bought"`
				Sold   int `json:"sold"`
			} `json:"points"`
			Bazaar struct {
				Customers int   `json:"customers"`
				Sales     int   `json:"sales"`
				Profit    int64 `json:"profit"`
			} `json:"bazaar"`
			ItemMarket struct {
				Customers int   `json:"customers"`
				Sales     int   `json:"sales"`
				Revenue   int64 `json:"revenue"`
				Fees      int64 `json:"fees"`
			} `json:"item_market"`
		} `json:"trading"`
		Jail struct {
			TimesJailed int `json:"times_jailed"`
			Busts       struct {
				Success int `json:"success"`
				Fails   int `json:"fails"`
			} `json:"busts"`
			Bails struct {
				Amount int   `json:"amount"`
				Fees   int64 `json:"fees"`
			} `json:"bails"`
		} `json:"jail"`
		Hospital struct {
			TimesHospitalized int `json:"times_hospitalized"`
			MedicalItemsUsed  int `json:"medical_items_used"`
			BloodWithdrawn    int `json:"blood_withdrawn"`
			Reviving          struct {
				Skill           int `json:"skill"`
				Revives         int `json:"revives"`
				RevivesReceived int `json:"revives_received"`
			} `json:"reviving"`
		} `json:"hospital"`
		FinishingHits struct {
			HeavyArtillery int `json:"heavy_artillery"`
			MachineGuns    int `json:"machine_guns"`
			Rifles         int `json:"rifles"`
			SubMachineGuns int `json:"sub_machine_guns"`
			Shotguns       int `json:"shotguns"`
			Pistols        int `json:"pistols"`
			Temporary      int `json:"temporary"`
			Piercing       int `json:"piercing"`
			Slashing       int `json:"slashing"`
			Clubbing       int `json:"clubbing"`
			Mechanical     int `json:"mechanical"`
			HandToHand     int `json:"hand_to_hand"`
		} `json:"finishing_hits"`
		Communication struct {
			MailsSent struct {
				Total      int `json:"total"`
				Friends    int `json:"friends"`
				Faction    int `json:"faction"`
				Colleagues int `json:"colleagues"`
				Spouse     int `json:"spouse"`
			} `json:"mails_sent"`
			ClassifiedAds int `json:"classified_ads"`
			Personals     int `json:"personals"`
		} `json:"communication"`
		Crimes struct {
			Total            int    `json:"total"`
			SellIllegalGoods int    `json:"sell_illegal_goods"`
			Theft            int    `json:"theft"`
			AutoTheft        int    `json:"auto_theft"`
			DrugDeals        int    `json:"drug_deals"`
			Computer         int    `json:"computer"`
			Fraud            int    `json:"fraud"`
			Murder           int    `json:"murder"`
			Other            int    `json:"other"`
			OrganizedCrimes  int    `json:"organized_crimes"`
			Version          string `json:"version"`
		} `json:"crimes"`
		Bounties struct {
			Placed struct {
				Amount int   `json:"amount"`
				Value  int64 `json:"value"`
			} `json:"placed"`
			Collected struct {
				Amount int   `json:"amount"`
				Value  int64 `json:"value"`
			} `json:"collected"`
			Received struct {
				Amount int   `json:"amount"`
				Value  int64 `json:"value"`
			} `json:"received"`
		} `json:"bounties"`
		Investments struct {
			Bank struct {
				Total         int64 `json:"total"`
				Profit        int64 `json:"profit"`
				Current       int64 `json:"current"`
				TimeRemaining int   `json:"time_remaining"`
			} `json:"bank"`
			Stocks struct {
				Profits    int64 `json:"profits"`
				Losses     int64 `json:"losses"`
				Fees       int64 `json:"fees"`
				NetProfits int64 `json:"net_profits"`
				Payouts    int64 `json:"payouts"`
			} `json:"stocks"`
		} `json:"investments"`
		Items struct {
			Found struct {
				City       int `json:"city"`
				Dump       int `json:"dump"`
				EasterEggs int `json:"easter_eggs"`
			} `json:"found"`
			Trashed int `json:"trashed"`
			Used    struct {
				Books         int `json:"books"`
				Boosters      int `json:"boosters"`
				Consumables   int `json:"consumables"`
				Candy         int `json:"candy"`
				Alcohol       int `json:"alcohol"`
				EnergyDrinks  int `json:"energy_drinks"`
				StatEnhancers int `json:"stat_enhancers"`
				EasterEggs    int `json:"easter_eggs"`
			} `json:"used"`
			VirusesCoded int `json:"viruses_coded"`
		} `json:"items"`
		Travel struct {
			Total       int `json:"total"`
			TimeSpent   int `json:"time_spent"`
			ItemsBought int `json:"items_bought"`
			Hunting     struct {
				Skill int `json:"skill"`
			} `json:"hunting"`
			AttacksWon         int `json:"attacks_won"`
			DefendsLost        int `json:"defends_lost"`
			Argentina          int `json:"argentina"`
			Canada             int `json:"canada"`
			CaymanIslands      int `json:"cayman_islands"`
			China              int `json:"china"`
			Hawaii             int `json:"hawaii"`
			Japan              int `json:"japan"`
			Mexico             int `json:"mexico"`
			UnitedArabEmirates int `json:"united_arab_emirates"`
			UnitedKingdom      int `json:"united_kingdom"`
			SouthAfrica        int `json:"south_africa"`
			Switzerland        int `json:"switzerland"`
		} `json:"travel"`
		Drugs struct {
			Cannabis        int `json:"cannabis"`
			Ecstasy         int `json:"ecstasy"`
			Ketamine        int `json:"ketamine"`
			Lsd             int `json:"lsd"`
			Opium           int `json:"opium"`
			Pcp             int `json:"pcp"`
			Shrooms         int `json:"shrooms"`
			Speed           int `json:"speed"`
			Vicodin         int `json:"vicodin"`
			Xanax           int `json:"xanax"`
			Total           int `json:"total"`
			Overdoses       int `json:"overdoses"`
			Rehabilitations struct {
				Amount int   `json:"amount"`
				Fees   int64 `json:"fees"`
			} `json:"rehabilitations"`
		} `json:"drugs"`
		Missions struct {
			Missions  int `json:"missions"`
			Contracts struct {
				Total int `json:"total"`
				Duke  int `json:"duke"`
			} `json:"contracts"`
			Credits int `json:"credits"`
		} `json:"missions"`
		Racing struct {
			Skill  int `json:"skill"`
			Points int `json:"points"`
			Races  struct {
				Entered int `json:"entered"`
				Won     int `json:"won"`
			} `json:"races"`
		} `json:"racing"`
		Networth struct {
			Total        int64 `json:"total"`
			Wallet       int64 `json:"wallet"`
			Vaults       int64 `json:"vaults"`
			Bank         int64 `json:"bank"`
			OverseasBank int64 `json:"overseas_bank"`
			Points       int64 `json:"points"`
			Inventory    int64 `json:"inventory"`
			DisplayCase  int64 `json:"display_case"`
			Bazaar       int64 `json:"bazaar"`
			ItemMarket   int64 `json:"item_market"`
			Property     int64 `json:"property"`
			StockMarket  int64 `json:"stock_market"`
			AuctionHouse int64 `json:"auction_house"`
			Bookie       int64 `json:"bookie"`
			Company      int64 `json:"company"`
			EnlistedCars int64 `json:"enlisted_cars"`
			PiggyBank    int64 `json:"piggy_bank"`
			Pending      int64 `json:"pending"`
			Loans        int64 `json:"loans"`
			UnpaidFees   int64 `json:"unpaid_fees"`
		} `json:"networth"`
		Other struct {
			Activity struct {
				Time   int `json:"time"`
				Streak struct {
					Best    int `json:"best"`
					Current int `json:"current"`
				} `json:"streak"`
			} `json:"activity"`
			Awards       int `json:"awards"`
			MeritsBought int `json:"merits_bought"`
			Refills      struct {
				Energy int `json:"energy"`
				Nerve  int `json:"nerve"`
				Token  int `json:"token"`
			} `json:"refills"`
			DonatorDays   int `json:"donator_days"`
			RankedWarWins int `json:"ranked_war_wins"`
		} `json:"other"`
	} `json:"personalstats"`
}
