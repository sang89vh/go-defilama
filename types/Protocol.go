package types

type Protocol struct {
	Id          string           `json:"id"`
	Name        string           `json:"name"`
	Address     string           `json:"address"`
	Symbol      string           `json:"symbol"`
	Url         string           `json:"url"`
	Description string           `json:"description"`
	Chain       string           `json:"chain"`
	Logo        string           `json:"logo"`
	Audits      string           `json:"audits"`
	AuditNote   interface{}      `json:"audit_note"`
	GeckoId     interface{}      `json:"gecko_id"`
	CmcId       string           `json:"cmcId"`
	Category    string           `json:"category"`
	Chains      []string         `json:"chains"`
	Oracles     []string         `json:"oracles"`
	ForkedFrom  []interface{}    `json:"forkedFrom"`
	Module      string           `json:"module"`
	Twitter     string           `json:"twitter"`
	AuditLinks  []string         `json:"audit_links"`
	ListedAt    float64          `json:"listedAt"`
	Slug        string           `json:"slug"`
	Tvl         float64          `json:"tvl"`
	ChainTvls   ProtocolChainTvl `json:"chainTvls"`
	Change1H    interface{}      `json:"change_1h"`
	Change1D    interface{}      `json:"change_1d"`
	Change7D    interface{}      `json:"change_7d"`
}
