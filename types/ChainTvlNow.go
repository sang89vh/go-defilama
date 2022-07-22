package types

type ChainTvlNow struct {
	GeckoId     interface{} `json:"gecko_id"`
	Tvl         float64     `json:"tvl"`
	TokenSymbol interface{} `json:"tokenSymbol"`
	CmcId       interface{} `json:"cmcId"`
	Name        string      `json:"name"`
	ChainId     interface{} `json:"chainId,string,omitempty"`
}
