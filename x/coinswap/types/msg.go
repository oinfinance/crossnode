package types

// All coinswap message define.
// deposit msg
type MsgDeposit struct {
	Chain   uint	`json:"chain"`
	Token	uint	`json:"token"`
	Amount  []byte 	`json:"amount"`
	TargetChain uint `json:"target"`
}

type MsgBurning struct {
	Chain   uint	`json:"chain"`
	Token	uint	`json:"token"`
	Amount  []byte 	`json:"amount"`
	TargetChain uint `json:"target"`
}

