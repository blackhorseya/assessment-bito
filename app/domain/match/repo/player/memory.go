package player

import (
	"errors"
	"sync"
	"time"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/google/uuid"
)

var (
	errPlayerNotFound  = errors.New("player not found in memory")
	errPlayerDuplicate = errors.New("player duplicate")
)

type memory struct {
	sync.RWMutex
	players map[string]*playerDTO
	pairs   map[string]*pairDTO
}

// NewPlayerRepoWithMemory is to create a new player repo with memory.
func NewPlayerRepoWithMemory() repo.IPlayerRepo {
	return &memory{
		players: make(map[string]*playerDTO),
		pairs:   make(map[string]*pairDTO),
	}
}

func (i *memory) GetByID(ctx contextx.Contextx, id string) (item *agg.Player, err error) {
	i.RLock()
	defer i.RUnlock()

	got, ok := i.players[id]
	if !ok {
		return nil, errPlayerNotFound
	}

	return got.ToAgg(), nil
}

func (i *memory) JoinPlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	i.Lock()
	defer i.Unlock()

	now := time.Now()
	player.ID = uuid.New().String()
	player.CreatedAt = now
	player.UpdatedAt = now

	created := newPlayerDTO(player)

	_, ok := i.players[player.ID]
	if ok {
		return errPlayerDuplicate
	}

	i.players[player.ID] = created

	return nil
}

func (i *memory) LeavePlayer(ctx contextx.Contextx, player *agg.Player) (err error) {
	i.Lock()
	defer i.Unlock()

	delete(i.players, player.ID)

	return nil
}

func (i *memory) ListPlayers(
	ctx contextx.Contextx,
	condition repo.ListPlayersCondition,
) (items []*agg.Player, total int, err error) {
	i.RLock()
	defer i.RUnlock()

	var ret []*agg.Player
	for _, player := range i.players {
		if condition.Gender != model.GenderUnspecified && condition.Gender != player.Gender {
			continue
		}

		if condition.Height > 0 && !(player.Height >= uint(condition.Height)) {
			continue
		}

		if condition.Height < 0 && !(player.Height < uint(-condition.Height)) {
			continue
		}

		if condition.NumsOfWantedDates > 0 && !(player.NumsOfWantedDates >= uint(condition.NumsOfWantedDates)) {
			continue
		}

		if condition.NumsOfWantedDates < 0 && !(player.NumsOfWantedDates < uint(-condition.NumsOfWantedDates)) {
			continue
		}

		ret = append(ret, player.ToAgg())
	}

	return ret, len(ret), nil
}

func (i *memory) MatchedPair(
	ctx contextx.Contextx,
	left, right *agg.Player,
	pair *model.Pair,
) (err error) {
	i.Lock()
	defer i.Unlock()

	now := time.Now()

	left.UpdatedAt = now
	leftDTO := newPlayerDTO(left)
	_, ok := i.players[leftDTO.ID]
	if ok {
		i.players[leftDTO.ID] = leftDTO
	}

	right.UpdatedAt = now
	rightDTO := newPlayerDTO(right)
	_, ok = i.players[rightDTO.ID]
	if ok {
		i.players[rightDTO.ID] = rightDTO
	}

	pair.ID = uuid.New().String()
	pair.CreatedAt = now
	newPair := newPairDTO(pair)
	i.pairs[pair.ID] = newPair

	return nil
}
