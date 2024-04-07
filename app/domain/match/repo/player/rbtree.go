package player

import (
	"sync"
	"time"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/gods/trees/rbtree"
	"github.com/google/uuid"
)

type rbtreeImpl struct {
	sync.RWMutex
	players map[string]*playerDTO
	pairs   map[string]*pairDTO
	rbtree  *rbtree.RBTree
}

// NewPlayerRepoWithRBTree is to create a new player repo with rbtreeImpl.
func NewPlayerRepoWithRBTree() repo.IPlayerRepo {
	return &rbtreeImpl{
		players: make(map[string]*playerDTO),
		pairs:   make(map[string]*pairDTO),
		rbtree:  rbtree.New(),
	}
}

func (i *rbtreeImpl) GetByID(ctx contextx.Contextx, id string) (item *agg.Player, err error) {
	i.RLock()
	defer i.RUnlock()

	got, ok := i.players[id]
	if !ok {
		return nil, errPlayerNotFound
	}

	return got.ToAgg(), nil
}

func (i *rbtreeImpl) JoinPlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	i.Lock()
	defer i.Unlock()

	now := time.Now()
	player.ID = uuid.New().String()
	player.CreatedAt = now
	player.UpdatedAt = now

	created := newPlayerDTO(player)

	_, ok := i.players[created.ID]
	if ok {
		return errPlayerDuplicate
	}

	i.players[created.ID] = created

	i.rbtree.Insert(created)

	return nil
}

func (i *rbtreeImpl) LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	i.Lock()
	defer i.Unlock()

	got, ok := i.players[player.ID]
	if !ok {
		return errPlayerNotFound
	}
	i.rbtree.Delete(got)

	delete(i.players, got.ID)

	return nil
}

func (i *rbtreeImpl) ListPlayers(
	ctx contextx.Contextx,
	condition repo.ListPlayersCondition,
) (items []*agg.Player, total int, err error) {
	i.RLock()
	defer i.RUnlock()

	var ret []*agg.Player
	filter := func(item rbtree.Item) bool {
		player, ok := item.(*playerDTO)
		if !ok {
			return false
		}

		if condition.Gender != model.GenderUnspecified && condition.Gender != player.Gender {
			return false
		}

		if condition.NumsOfWantedDates > 0 && !(player.NumsOfWantedDates >= uint(condition.NumsOfWantedDates)) {
			return false
		}

		if condition.NumsOfWantedDates < 0 && !(player.NumsOfWantedDates < uint(-condition.NumsOfWantedDates)) {
			return false
		}

		ret = append(ret, player.ToAgg())
		return true
	}

	if condition.Height < 0 {
		i.rbtree.Descend(&playerDTO{Height: uint(-condition.Height)}, filter)
	}

	if condition.Height > 0 {
		i.rbtree.Ascend(&playerDTO{Height: uint(condition.Height)}, filter)
	}

	if condition.Height == 0 {
		i.rbtree.Ascend(&playerDTO{Height: 0}, filter)
	}

	return ret, len(ret), nil
}

func (i *rbtreeImpl) MatchedPair(ctx contextx.Contextx, left, right *agg.Player, pair *model.Pair) (err error) {
	i.Lock()
	defer i.Unlock()

	now := time.Now()
	left.UpdatedAt = now
	right.UpdatedAt = now

	leftDTO, ok := i.players[left.ID]
	if ok {
		leftDTO.NumsOfWantedDates = left.NumsOfWantedDates
		leftDTO.UpdatedAt = left.UpdatedAt
	}

	rightDTO, ok := i.players[right.ID]
	if ok {
		rightDTO.NumsOfWantedDates = right.NumsOfWantedDates
		rightDTO.UpdatedAt = right.UpdatedAt
	}

	pair.ID = uuid.New().String()
	pair.CreatedAt = now
	newPair := newPairDTO(pair)
	i.pairs[pair.ID] = newPair

	return nil
}
