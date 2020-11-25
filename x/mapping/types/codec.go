package types

import "github.com/cosmos/cosmos-sdk/codec"

// ModuleCdc defines the codec to be used by evm module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterCodec registers concrete types and interfaces on the given codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&MsgRegister{}, "mapping/MsgRegister", nil)
	cdc.RegisterConcrete(&MsgMapVerify{}, "mapping/MsgMapVerify", nil)

	cdc.RegisterConcrete(&MappingInfo{}, "mapping/MappingInfo", nil)
}
