package market

import "torngo/internal/api"

type MarketBazaarResponse struct {
	Bazaar struct {
		Busiest []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			IsOpen          bool   `json:"is_open"`
			WeeklyCustomers int    `json:"weekly_customers"`
		} `json:"busiest"`
		MostPopular []struct {
			ID             int    `json:"id"`
			Name           string `json:"name"`
			IsOpen         bool   `json:"is_open"`
			TotalFavorites int    `json:"total_favorites"`
		} `json:"most_popular"`
		Trending []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			IsOpen          bool   `json:"is_open"`
			RecentFavorites int    `json:"recent_favorites"`
		} `json:"trending"`
		TopGrossing []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			IsOpen       bool   `json:"is_open"`
			WeeklyIncome int64  `json:"weekly_income"`
		} `json:"top_grossing"`
		Bulk []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			IsOpen    bool   `json:"is_open"`
			BulkSales int    `json:"bulk_sales"`
		} `json:"bulk"`
		AdvancedItem []struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			IsOpen            bool   `json:"is_open"`
			AdvancedItemSales int    `json:"advanced_item_sales"`
		} `json:"advanced_item"`
		Bargain []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			IsOpen       bool   `json:"is_open"`
			BargainSales int    `json:"bargain_sales"`
		} `json:"bargain"`
		DollarSale []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			IsOpen      bool   `json:"is_open"`
			DollarSales int    `json:"dollar_sales"`
		} `json:"dollar_sale"`
	} `json:"bazaar"`
}

type MarketBazaarCategory string

const (
	MarketBazaarCategoryAlcohol     MarketBazaarCategory = "alcohol"
	MarketBazaarCategoryArtifact    MarketBazaarCategory = "artifact"
	MarketBazaarCategoryBooster     MarketBazaarCategory = "booster"
	MarketBazaarCategoryCandy       MarketBazaarCategory = "candy"
	MarketBazaarCategoryCar         MarketBazaarCategory = "car"
	MarketBazaarCategoryClothing    MarketBazaarCategory = "clothing"
	MarketBazaarCategoryCollectible MarketBazaarCategory = "collectible"
	MarketBazaarCategoryDefensive   MarketBazaarCategory = "defensive"
	MarketBazaarCategoryDrug        MarketBazaarCategory = "drug"
	MarketBazaarCategoryEnergyDrink MarketBazaarCategory = "energy drink"
	MarketBazaarCategoryEnhancer    MarketBazaarCategory = "enhancer"
	MarketBazaarCategoryFlower      MarketBazaarCategory = "flower"
	MarketBazaarCategoryJewelry     MarketBazaarCategory = "jewelry"
	MarketBazaarCategoryMaterial    MarketBazaarCategory = "material"
	MarketBazaarCategoryMedical     MarketBazaarCategory = "medical"
	MarketBazaarCategoryMelee       MarketBazaarCategory = "melee"
	MarketBazaarCategoryOther       MarketBazaarCategory = "other"
	MarketBazaarCategoryPlushie     MarketBazaarCategory = "plushie"
	MarketBazaarCategoryPrimary     MarketBazaarCategory = "primary"
	MarketBazaarCategorySecondary   MarketBazaarCategory = "secondary"
	MarketBazaarCategorySpecial     MarketBazaarCategory = "special"
	MarketBazaarCategorySupplyPack  MarketBazaarCategory = "supply pack"
	MarketBazaarCategoryTemporary   MarketBazaarCategory = "temporary"
	MarketBazaarCategoryTool        MarketBazaarCategory = "tool"
)

type MarketBazaarParams struct {
	Category  MarketBazaarCategory
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
