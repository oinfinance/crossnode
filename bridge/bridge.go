package bridge

import "math/big"

type Address []byte
type Receipt []byte
type Receipts []Receipt

/*
 * Bridge define the interface that could query account's assert and do a transfer.
 */
type Bridge interface {
	GetBalance(account Address) *big.Int
	Transfer(from, to Address, value *big.Int) Receipt
	BatchTransfer(from Address, to []Address, value []*big.Int) Receipts
}
