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
	men     *rbtree.RBTree
	women   *rbtree.RBTree
}

// NewPlayerRepoWithRBTree is to create a new player repo with rbtreeImpl.
func NewPlayerRepoWithRBTree() repo.IPlayerRepo {
	return &rbtreeImpl{
		RWMutex: sync.RWMutex{},
		players: make(map[string]*playerDTO),
		pairs:   make(map[string]*pairDTO),
		men:     rbtree.New(),
		women:   rbtree.New(),
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

	if player.Profile.Gender == model.GenderMale {
		i.men.Insert(created)
	} else if player.Profile.Gender == model.GenderFemale {
		i.women.Insert(created)
	}

	return nil
}

func (i *rbtreeImpl) LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	i.Lock()
	defer i.Unlock()

	got, ok := i.players[player.ID]
	if !ok {
		return errPlayerNotFound
	}

	if player.Profile.Gender == model.GenderMale {
		i.men.Delete(got)
	} else if player.Profile.Gender == model.GenderFemale {
		i.women.Delete(got)
	}

	delete(i.players, got.ID)

	return nil
}

func (i *rbtreeImpl) ListPlayers(
	ctx contextx.Contextx,
	condition repo.ListPlayersCondition,
) (items []*agg.Player, total int, err error) {
	i.RLock()
	defer i.RUnlock()

	var players []*agg.Player
	add := func(item rbtree.Item) bool {
		player, ok := item.(*playerDTO)
		if !ok {
			return false
		}

		players = append(players, player.ToAgg())
		return true
	}

	if condition.Height < 0 {
		i.women.Descend(&playerDTO{Height: uint(-condition.Height)}, add)
	}

	if condition.Height > 0 {
		i.men.Ascend(&playerDTO{Height: uint(condition.Height)}, add)
	}

	if condition.Height == 0 {
		for _, player := range i.players {
			players = append(players, player.ToAgg())
		}
	}

	var ret []*agg.Player
	for _, player := range players {
		if condition.Gender != model.GenderUnspecified && condition.Gender != player.Profile.Gender {
			continue
		}

		if condition.NumsOfWantedDates > 0 && !(player.NumsOfWantedDates >= uint(condition.NumsOfWantedDates)) {
			continue
		}

		if condition.NumsOfWantedDates < 0 && !(player.NumsOfWantedDates < uint(-condition.NumsOfWantedDates)) {
			continue
		}

		ret = append(ret, player)
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
