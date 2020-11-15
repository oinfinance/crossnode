package coinswap

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oinfinance/crossnode/x/coinswap/types"

	//"github.com/oinfinance/crossnode/bridge"
	"github.com/oinfinance/crossnode/x/coinswap/keeper"
)

var (
	RefreshPoint = int64(150)
)

func GenerateReceipt() []byte {
	return []byte("target receipt")
}

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	blockNumber := ctx.BlockHeight()

	records := k.GetAllRecord(ctx)
	dealed := make([]*types.CoinSwapRecordStorage, 0)
	for _, record := range records {
		if record.Status != types.RecordStatusWaited {
			continue
		}
		if record.AddedBlock.Int64()-blockNumber < RefreshPoint {
			continue
		}
		// sign target chain tx with param.
		r := GenerateReceipt()
		record.Receipt = hex.EncodeToString(r)
		record.Status = types.RecordStatusSucceed
		dealed = append(dealed, record)
	}

	for _, r := range dealed {
		k.UpdateRecord(ctx, r.Hash(), r)
	}
}
