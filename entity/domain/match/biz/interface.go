//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
)

// ListPairsOption is the option for list pairs.
type ListPairsOption struct {
	Page int
	Size int
}

// ListPlayersOption is the option for list players.
type ListPlayersOption struct {
	Page int
	Size int
}

// IMatchBiz is an interface that represents the business logic of match.
type IMatchBiz interface {
	// EnrollPlayer is to enroll a player.
	EnrollPlayer(
		ctx contextx.Contextx,
		name string,
		height uint,
		gender model.Gender,
		age uint,
		numsOfWantedDates uint,
	) (item *agg.Player, err error)

	// UnregisterPlayer is to unregister a player.
	UnregisterPlayer(ctx contextx.Contextx, playerID string) (err error)

	// GetPlayerByIDWithPairs is to get a player by id with pairs.
	GetPlayerByIDWithPairs(
		ctx contextx.Contextx,
		playerID string,
		option ListPairsOption,
	) (item *agg.Player, err error)

	// ListPlayers is to list players.
	ListPlayers(ctx contextx.Contextx, option ListPlayersOption) (items []*agg.Player, total int, err error)

	// SubmitPair is to submit a pair.
	SubmitPair(ctx contextx.Contextx, leftID string, rightID string) (item *model.Pair, err error)
}
