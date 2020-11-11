package types

const (
	QueryListChains = "list-chains"
	QueryShowChain  = "show-chain"
)

type QueryListChainsParams struct {
	Page, Limit int
}
