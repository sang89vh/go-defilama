package types

type CurrentChainTvl struct {
	Pool2             float64 `json:"pool2"`
	Avalanche         float64 `json:"Avalanche"`
	PolygonBorrowed   float64 `json:"Polygon-borrowed"`
	EthereumPool2     float64 `json:"Ethereum-pool2"`
	Ethereum          float64 `json:"Ethereum"`
	EthereumStaking   float64 `json:"Ethereum-staking"`
	Polygon           float64 `json:"Polygon"`
	EthereumBorrowed  float64 `json:"Ethereum-borrowed"`
	Borrowed          float64 `json:"borrowed"`
	AvalancheBorrowed float64 `json:"Avalanche-borrowed"`
	Staking           float64 `json:"staking"`
	Optimism          float64 `json:"Optimism"`
	ArbitrumBorrowed  float64 `json:"Arbitrum-borrowed"`
	Harmony           float64 `json:"Harmony"`
	HarmonyBorrowed   float64 `json:"Harmony-borrowed"`
	FantomBorrowed    float64 `json:"Fantom-borrowed"`
	OptimismBorrowed  float64 `json:"Optimism-borrowed"`
	Arbitrum          float64 `json:"Arbitrum"`
	Fantom            float64 `json:"Fantom"`
}
