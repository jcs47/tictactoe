package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgMakeMove{}

type MsgMakeMove struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	GameId  string         `json:"gameId" yaml:"gameId"`
	X       int            `json:"x" yaml:"x"`
	Y       int            `json:"y" yaml:"y"`
}

func NewMsgMakeMove(creator sdk.AccAddress, gameId string, x, y int) MsgMakeMove {
	return MsgMakeMove{
		ID:      uuid.New().String(),
		Creator: creator,
		GameId:  gameId,
		X:       x,
		Y:       y,
	}
}

func (msg MsgMakeMove) Route() string {
	return RouterKey
}

func (msg MsgMakeMove) Type() string {
	return "CreateMove"
}

func (msg MsgMakeMove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgMakeMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgMakeMove) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
