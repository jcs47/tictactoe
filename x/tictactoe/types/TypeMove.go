package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Move struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	GameId  string         `json:"gameId" yaml:"gameId"`
	X       int            `json:"x" yaml:"x"`
	Y       int            `json:"y" yaml:"y"`
}
