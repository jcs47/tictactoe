package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgMakeMove{}, "tictactoe/MakeMove", nil)
	cdc.RegisterConcrete(MsgAcceptInvitation{}, "tictactoe/AcceptInvitation", nil)
	cdc.RegisterConcrete(MsgCreateInvitation{}, "tictactoe/CreateInvitation", nil)
	cdc.RegisterConcrete(MsgDeleteInvitation{}, "tictactoe/DeleteInvitation", nil)
	// TODO: Register the modules msgs
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
