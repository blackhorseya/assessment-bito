//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
)

// ListPlayersCondition is the condition for list players.
type ListPlayersCondition struct {
	Gender model.Gender

	// Height cm unit,
	// 0 means no limit
	// >0 means the height is greater than the value
	// <0 means the height is less than the value.
	Height int

	// NumOfWantedDates is the number of wanted dates.
	// 0 means no limit
	// >0 means the number is greater than the value
	// <0 means the number is less than the value.
	NumsOfWantedDates int

	Limit  int
	Offset int
}

// IPlayerRepo is an interface that represents the repository of player.
type IPlayerRepo interface {
	// GetByID is to get a player by id.
	GetByID(ctx contextx.Contextx, id string) (item *agg.Player, err error)

	// JoinPlayer is to join a player.
	JoinPlayer(ctx contextx.Contextx, player *agg.Player) (err error)

	// LeavePlayer is to leave a player.
	LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error)

	// ListPlayers is to list players.
	ListPlayers(ctx contextx.Contextx, condition ListPlayersCondition) (items []*agg.Player, total int, err error)

	// MatchedPair is to get a matched pair.
	MatchedPair(ctx contextx.Contextx, left, right *agg.Player, pair *model.Pair) (err error)
}
