package player

import (
	"sync"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/gods/trees/rbtree"
)

type rbtreeImpl struct {
	sync.RWMutex
	rbtree *rbtree.RBTree
}

// NewPlayerRepoWithRBTree is to create a new player repo with rbtreeImpl.
func NewPlayerRepoWithRBTree() repo.IPlayerRepo {
	return &rbtreeImpl{
		rbtree: rbtree.New(),
	}
}

func (i *rbtreeImpl) GetByID(ctx contextx.Contextx, id string) (item *agg.Player, err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtreeImpl) JoinPlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtreeImpl) LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtreeImpl) ListPlayers(
	ctx contextx.Contextx,
	condition repo.ListPlayersCondition,
) (items []*agg.Player, total int, err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}

func (i *rbtreeImpl) MatchedPair(ctx contextx.Contextx, left, right *agg.Player, pair *model.Pair) (err error) {
	// todo: 2024/4/5|sean|implement me
	panic("implement me")
}
