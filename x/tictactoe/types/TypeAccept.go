package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Accept struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    InvitationId string `json:"invitationId" yaml:"invitationId"`
}