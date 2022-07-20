package types

type Pool2 struct {
	Tvl         []Tvl       `json:"tvl"`
	Tokens      interface{} `json:"tokens"`
	TokensInUsd interface{} `json:"tokensInUsd"`
}
