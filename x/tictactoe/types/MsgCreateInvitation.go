package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateInvitation{}

type MsgCreateInvitation struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Message string         `json:"message" yaml:"message"`
}

func NewMsgCreateInvitation(creator sdk.AccAddress, msg string) MsgCreateInvitation {
	return MsgCreateInvitation{
		ID:      uuid.New().String(),
		Creator: creator,
		Message: msg,
	}
}

func (msg MsgCreateInvitation) Route() string {
	return RouterKey
}

func (msg MsgCreateInvitation) Type() string {
	return "CreateInvitation"
}

func (msg MsgCreateInvitation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateInvitation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateInvitation) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
