package types

const (
	// module name
	ModuleName = "mining"

	// StoreKey is the default store key for guardian
	StoreKey = ModuleName

	// RouterKey is the message route for guardian
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the guardian store.
	QuerierRoute = StoreKey
)
