package tictactoe

import (
	"crypto/sha256"
	"fmt"
	"math/bits"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/lisbon/jcs/x/tictactoe/keeper"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

func handleMsgAcceptInvitation(ctx sdk.Context, k keeper.Keeper, a auth.AccountKeeper, msg types.MsgAcceptInvitation) (*sdk.Result, error) {

	log := ctx.Logger()

	invitation, err := k.GetInvitation(ctx, msg.InvitationId)

	if err != nil {

		return nil, fmt.Errorf("Failed to initiate new game: %s", err)
	}

	var game types.TTTGame

	gameID := msg.InvitationId

	initiator := invitation.Creator
	acceptor := msg.Creator

	initiatorPK, err := a.GetPubKey(ctx, initiator)

	if err != nil {

		return nil, fmt.Errorf("Failed to initiate new game: %s", err)
	}

	acceptorPK, err := a.GetPubKey(ctx, acceptor)

	if err != nil {

		return nil, fmt.Errorf("Failed to initiate new game: %s", err)
	}

	concat := append(initiatorPK.Bytes(), acceptorPK.Bytes()...)
	hash := sha256.Sum256(concat)

	log.Info(fmt.Sprintf("Leading bits for %x: %d (first byte is %08b)\n", hash, bits.LeadingZeros8(hash[0]), hash[0]))

	if bits.LeadingZeros8(hash[0]) > 0 {

		game, err = types.NewTTTGame(gameID, initiator, acceptor)

	} else {

		game, err = types.NewTTTGame(gameID, acceptor, initiator)
	}

	if err != nil {

		return nil, fmt.Errorf("Failed to initiate new game: %s\n", err)
	}

	k.DeleteInvitation(ctx, invitation.ID)

	k.CreateTTTGame(ctx, game)

	log.Info(fmt.Sprintf("Game started with ID %s", game.ID))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
