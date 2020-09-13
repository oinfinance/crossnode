package types

// deposit msg
type DepositParam struct {
	Chain   uint	`json:"chain"`
	Token	uint	`json:"token"`
	Amount  []byte 	`json:"amount"`
	TargetChain uint `json:"target"`
}

type Burning struct {
	Chain 	uint	`json:"chain"`
	Token	uint	`json:"token"`
	Amount  []byte 	`json:"amount"`
	TargetChain uint `json:"target"`
	Account string   `json:"account"`
}


