package player

import (
	"sync"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
)

type rbtree struct {
	sync.RWMutex
}

// NewPlayerRepoWithRBTree is to create a new player repo with rbtree.
func NewPlayerRepoWithRBTree() repo.IPlayerRepo {
	return &rbtree{}
}

func (i *rbtree) GetByID(ctx contextx.Contextx, id string) (item *agg.Player, err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtree) JoinPlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtree) LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtree) ListPlayers(
	ctx contextx.Contextx,
	condition repo.ListPlayersCondition,
) (items []*agg.Player, total int, err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtree) MatchedPair(ctx contextx.Contextx, left, right *agg.Player, pair *model.Pair) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}
