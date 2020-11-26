package types

const (
	// module name
	ModuleName = "mapping"

	// StoreKey is the default store key for guardian
	MapinfoStoreKey = ModuleName
	VerifyStoreKey  = ModuleName + "verify"

	// RouterKey is the message route for guardian
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the guardian store.
	QuerierRoute = MapinfoStoreKey
)
