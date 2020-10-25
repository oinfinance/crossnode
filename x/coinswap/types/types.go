package types

import (
	"math/big"
)

type Object struct {
	Coins map[int]big.Int		// chain/token ==> amount
}

type State struct {
	Objects map[string]Object	// all user object
}