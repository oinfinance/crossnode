package coinswap

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/types"
	"golang.org/x/crypto/sha3"

	//"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/coinswap/keeper"
)

var (
	RefreshPoint = int64(150)
)

func GenerateReceipt() []byte {
	data := []byte("target receipt")
	hash := sha3.Sum256(data)
	return hash[:]
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	blockNumber := ctx.BlockHeight()

	if blockNumber <= 1 {
		return
	}

	records := k.GetAllRecord(ctx)
	dealed := make([]*types.CoinSwapRecordStorage, 0)
	for _, record := range records {
		if record.Receipt.Status != types.RecordStatusWaited {
			continue
		}
		if (blockNumber - int64(record.Record.AddedBlock)) < RefreshPoint {
			continue
		}
		// sign target chain tx with param.
		// todo: make a signature for user to mint coin.
		r := GenerateReceipt()
		//
		record.Receipt.Receipt = hex.EncodeToString(r)
		record.Receipt.Status = types.RecordStatusSucceed
		dealed = append(dealed, record)
	}

	for _, r := range dealed {
		k.UpdateRecord(ctx, r.Record.Hash(), r)
	}
}
