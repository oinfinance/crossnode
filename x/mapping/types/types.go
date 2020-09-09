package types

import "math/big"

type MappingInfo struct {
	RemoteAddr []byte
	ChainId    uint
	TokenType  uint
	MyAddress  []byte
	Balance    *big.Int
}

type QueryInfoResponse struct {
	Found      bool   `json:"found"`
	RemoteAddr []byte `json:"account"'`
	ChainId    uint   `json:"chainId"`
	TokenType  uint   `json:"tokenType"`
	MyAddress  []byte `json:"myAddress"`
	Balance    string `json:"balance"'`
}
