package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Invitation struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Message string         `json:"message" yaml:"message"`
}
