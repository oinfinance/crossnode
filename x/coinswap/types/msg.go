package types

// All coinswap message define.

type MsgRegister struct {
	SourceAccount []byte `json:"account"`   // source chain account
	ChainId       uint   `json:"chainId"`   // chain identification
	TokenType     uint   `json:"token"`     // token type identification
	MyAddress     []byte `json:"myAddress"` // binding address with local address
}

func NewMsgRegister(sourceAccount []byte, chainid uint, tokenType uint, myAddress []byte) *MsgRegister {
	return &MsgRegister{
		SourceAccount: sourceAccount,
		ChainId:       chainid,
		TokenType:     tokenType,
		MyAddress:     myAddress,
	}
}
