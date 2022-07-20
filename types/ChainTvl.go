package types

type ChainTvl struct {
	Pool2             Pool2          `json:"pool2"`
	Avalanche         ChainTvlDetail `json:"Avalanche"`
	PolygonBorrowed   ChainTvlDetail `json:"Polygon-borrowed"`
	EthereumPool2     ChainTvlDetail `json:"Ethereum-pool2"`
	Ethereum          ChainTvlDetail `json:"Ethereum"`
	EthereumStaking   ChainTvlDetail `json:"Ethereum-staking"`
	Polygon           ChainTvlDetail `json:"Polygon"`
	EthereumBorrowed  ChainTvlDetail `json:"Ethereum-borrowed"`
	Borrowed          ChainTvlDetail `json:"borrowed"`
	AvalancheBorrowed ChainTvlDetail `json:"Avalanche-borrowed"`
	Staking           ChainTvlDetail `json:"staking"`
	Optimism          ChainTvlDetail `json:"Optimism"`
	ArbitrumBorrowed  ChainTvlDetail `json:"Arbitrum-borrowed"`
	Harmony           ChainTvlDetail `json:"Harmony"`
	HarmonyBorrowed   ChainTvlDetail `json:"Harmony-borrowed"`
	FantomBorrowed    ChainTvlDetail `json:"Fantom-borrowed"`
	OptimismBorrowed  ChainTvlDetail `json:"Optimism-borrowed"`
	Arbitrum          ChainTvlDetail `json:"Arbitrum"`
	Fantom            ChainTvlDetail `json:"Fantom"`
}
