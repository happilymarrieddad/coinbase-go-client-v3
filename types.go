package coinbasegoclientv3

import (
	"time"

	"github.com/google/uuid"
)

type ValueCurrency struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Account struct {
	UUID             uuid.UUID     `json:"uuid"`
	Name             string        `json:"name"`
	Currency         string        `json:"currency"`
	AvailableBalance ValueCurrency `json:"available_balance"`
	Default          bool          `json:"default"`
	Active           bool          `json:"active"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        *time.Time    `json:"updated_at"`
	DeletedAt        *time.Time    `json:"deleted_at"`
	Type             string        `json:"type"`
	Ready            bool          `json:"ready"`
	Hold             ValueCurrency `json:"hold"`
}

// {"product_id":"BTC-USD","price":"22824.43","price_percentage_change_24h":"-0.78375186320378","volume_24h":"12080.83720814","volume_percentage_change_24h":"87.6833024653734","base_increment":"0.00000001","quote_increment":"0.01","quote_min_size":"1","quote_max_size":"50000000","base_min_size":"0.000016","base_max_size":"2600","base_name":"Bitcoin","quote_name":"US Dollar","watched":true,"is_disabled":false,"new":false,"status":"online","cancel_only":false,"limit_only":false,"post_only":false,"trading_disabled":false,"auction_mode":false,"product_type":"SPOT","quote_currency_id":"USD","base_currency_id":"BTC","fcm_trading_session_details":null,"mid_market_price":"","alias":"","alias_to":[],"base_display_symbol":"BTC","quote_display_symbol":"USD"}
type Product struct {
	ProductID                 string        `json:"product_id"`
	Price                     string        `json:"price"`
	PricePercentageChange24H  string        `json:"price_percentage_change_24h"`
	Volume24H                 string        `json:"volume_24h"`
	VolumePercentageChange24H string        `json:"volume_percentage_change_24h"`
	BaseIncrement             string        `json:"base_increment"`
	QuoteIncrement            string        `json:"quote_increment"`
	QuoteMinSize              string        `json:"quote_min_size"`
	QuoteMaxSize              string        `json:"quote_max_size"`
	BaseMinSize               string        `json:"base_min_size"`
	BaseMaxSize               string        `json:"base_max_size"`
	BaseName                  string        `json:"base_name"`
	QuoteName                 string        `json:"quote_name"`
	Watched                   bool          `json:"watched"`
	IsDisabled                bool          `json:"is_disabled"`
	New                       bool          `json:"new"`
	Status                    string        `json:"status"`
	CancelOnly                bool          `json:"cancel_only"`
	PostOnly                  bool          `json:"post_only"`
	TradingDisabled           bool          `json:"trading_disabled"`
	AuctionMode               bool          `json:"auction_mode"`
	ProductType               string        `json:"product_type"`
	QuoteCurrencyID           string        `json:"quote_currency_id"` // USD
	BaseCurrencyID            string        `json:"base_currency_id"`  // BTC
	FCMTradingSessionDetails  interface{}   `json:"fcm_trading_session_details"`
	MidMarketPrice            string        `json:"mid_market_price"`
	Alias                     string        `json:"alias"`
	AliasTo                   []interface{} `json:"alias_to"`
	BaseDisplaySymbol         string        `json:"base_display_symbol"`  // BTC
	QuoteDisplaySymbol        string        `json:"quote_display_symbol"` // USD
}

type ProductCandle struct {
	Start  string `json:"start"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
}

type Granularity string

const (
	OneMinuteGranularity     Granularity = "ONE_MINUTE"
	FiveMinuteGranularity    Granularity = "FIVE_MINUTE"
	FifteenMinuteGranularity Granularity = "FIFTEEN_MINUTE"
	ThirtyMinuteGranularity  Granularity = "THIRTY_MINUTE"
	OneHourGranularity       Granularity = "ONE_HOUR"
	TwoHourGranularity       Granularity = "TWO_HOUR"
	SixHourGranularity       Granularity = "SIX_HOUR"
	OneDayGranularity        Granularity = "ONE_DAY"
)

type Side string

const (
	UnknownOrderSide Side = "UNKNOWN_ORDER_SIDE"
	BuySide          Side = "BUY"
	SellSide         Side = "SELL"
)

type MarketTrade struct {
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Time      string `json:"time"` // "time": "2021-05-31T09:59:59Z",
	Side      Side   `json:"side"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
}

type FeeTier struct {
	PricingTier  string `json:"pricing_tier"`
	USDFrom      string `json:"usd_from"`
	USDTo        string `json:"usd_to"`
	MakerFeeRate string `json:"maker_fee_rate"`
}

type MarginRate struct {
	Value string `json:"value"`
}

type GoodsAndServicesTaxType string

const (
	InclusiveGoodsAndServicesTaxType string = "INCLUSIVE"
	ExclusiveGoodsAndServicesTaxType string = "EXCLUSIVE"
)

type GoodsAndServicesTax struct {
	Rate string                  `json:"rate"`
	Type GoodsAndServicesTaxType `json:"type"`
}

type TransactionSummary struct {
	TotalVolume             float64             `json:"total_volume"`
	TotalFees               float64             `json:"total_fees"`
	FeeTier                 FeeTier             `json:"free_tier"`
	MarginRate              MarginRate          `json:"margin_rate"`
	GoodsAndServicesTax     GoodsAndServicesTax `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume float64             `json:"advanced_trade_only_volume"`
	AdvancedTradeOnlyFees   float64             `json:"advanced_trade_only_fees"`
	CoinbaseProVolume       float64             `json:"coinbase_pro_volume"`
	CoinbaseProFees         float64             `json:"coinbase_pro_fees"`
}

type ProductType string

const (
	SpotProductType ProductType = "SPOT"
)
