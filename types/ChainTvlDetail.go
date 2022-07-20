package types

type ChainTvlDetail struct {
	Tvl         []Tvl       `json:"tvl"`
	Tokens      interface{} `json:"tokens"`
	TokensInUsd interface{} `json:"tokensInUsd"`
}
