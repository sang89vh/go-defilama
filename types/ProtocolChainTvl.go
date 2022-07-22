package types

type ProtocolChainTvl struct {
	Binance           float64 `json:"Binance"`
	AvalancheTreasury float64 `json:"Avalanche-treasury"`
	BinanceTreasury   float64 `json:"Binance-treasury"`
	PolygonTreasury   float64 `json:"Polygon-treasury"`
	Treasury          float64 `json:"treasury"`
}
