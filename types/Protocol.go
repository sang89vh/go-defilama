package types

type Protocol struct {
	Id               string          `json:"id"`
	Name             string          `json:"name"`
	Url              string          `json:"url"`
	Description      string          `json:"description"`
	Logo             string          `json:"logo"`
	Chains           []interface{}   `json:"chains"`
	GeckoId          string          `json:"gecko_id"`
	CmcId            string          `json:"cmcId"`
	Twitter          string          `json:"twitter"`
	CurrentChainTvls CurrentChainTvl `json:"currentChainTvls"`
	ChainTvls        ChainTvl        `json:"chainTvls"`
	Tokens           interface{}     `json:"tokens"`
	TokensInUsd      interface{}     `json:"tokensInUsd"`
	Tvl              Tvl             `json:"tvl"`
	OtherProtocols   []string        `json:"otherProtocols"`
}
