package cryptogate

type Network struct {
	Slug     NetworkSlug
	Name     string
	Status   NetworkStatus
	Settings NetworkSettings
}

type NetworkSettings struct {
	RPCUrlHttp string      `json:"rpc_url_http"`
	RPCUrlWs   string      `json:"rpc_url_ws"`
	Type       NetworkType `json:"type"`
}

type NetworkType string

const (
	NetworkTypeTestnet NetworkType = "testnet"
	NetworkTypeMainnet NetworkType = "mainnet"
)

type TypeConnection string

const (
	TypeConnectionHTTP TypeConnection = "http"
	TypeConnectionWS   TypeConnection = "ws"
)

type NetworkStatus string

const (
	NetworkStatusActive  NetworkStatus = "active"
	NetworkStatusDisable NetworkStatus = "disable"
)
