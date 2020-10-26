package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgAcceptInvitation{}

type MsgAcceptInvitation struct {
	ID           string
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	InvitationId string         `json:"invitationId" yaml:"invitationId"`
}

func NewMsgAcceptInvitation(creator sdk.AccAddress, invitationId string) MsgAcceptInvitation {
	return MsgAcceptInvitation{
		ID:           uuid.New().String(),
		Creator:      creator,
		InvitationId: invitationId,
	}
}

func (msg MsgAcceptInvitation) Route() string {
	return RouterKey
}

func (msg MsgAcceptInvitation) Type() string {
	return "CreateAccept"
}

func (msg MsgAcceptInvitation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgAcceptInvitation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgAcceptInvitation) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	return nil
}
