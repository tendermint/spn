package types

const (
	QueryListChains = "list-chains"
)

type QueryListChainsParams struct {
	Page, Limit int
}