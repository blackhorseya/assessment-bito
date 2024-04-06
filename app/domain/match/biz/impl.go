package biz

import (
	"errors"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"go.uber.org/zap"
)

var (
	errPlayerNotFound = errors.New("player not found")
)

type impl struct {
	players repo.IPlayerRepo
}

// NewMatchBiz is to create a new match biz.
func NewMatchBiz(players repo.IPlayerRepo) biz.IMatchBiz {
	return &impl{
		players: players,
	}
}

func (i *impl) EnrollPlayer(
	ctx contextx.Contextx,
	name string,
	height uint,
	gender model.Gender,
	age uint,
	numsOfWantedDates uint,
) (item *agg.Player, err error) {
	player, err := agg.NewPlayer(name, height, gender, age, numsOfWantedDates)
	if err != nil {
		ctx.Error("new player error", zap.Error(err))
		return nil, err
	}

	err = i.players.JoinPlayer(ctx, player)
	if err != nil {
		ctx.Error("join player error", zap.Error(err))
		return nil, err
	}

	return player, nil
}

func (i *impl) UnregisterPlayer(ctx contextx.Contextx, playerID string) (err error) {
	player, err := i.players.GetByID(ctx, playerID)
	if err != nil {
		ctx.Error("get player by id error", zap.Error(err))
		return err
	}
	if player == nil {
		ctx.Error("player not found", zap.String("playerID", playerID), zap.Error(errPlayerNotFound))
		return errPlayerNotFound
	}

	err = i.players.LeavePlayer(ctx, player)
	if err != nil {
		ctx.Error("leave player error", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) GetPlayerByIDWithPairs(
	ctx contextx.Contextx,
	playerID string,
	option biz.ListPairsOption,
) (item *agg.Player, err error) {
	player, err := i.players.GetByID(ctx, playerID)
	if err != nil {
		ctx.Error("get player by id error", zap.Error(err))
		return nil, err
	}
	if player == nil {
		ctx.Error("player not found", zap.String("playerID", playerID), zap.Error(errPlayerNotFound))
		return nil, errPlayerNotFound
	}

	wantGender, err := player.WantGender()
	if err != nil {
		ctx.Error("get want gender error", zap.Error(err))
		return nil, err
	}

	condition := repo.ListPlayersCondition{
		Gender:            wantGender,
		Height:            player.WantHeight(),
		NumsOfWantedDates: 1,
		Limit:             option.Size,
		Offset:            (option.Page - 1) * option.Size,
	}
	targets, _, err := i.players.ListPlayers(ctx, condition)
	if err != nil {
		ctx.Error("list players error", zap.Error(err), zap.Any("condition", condition))
		return nil, err
	}

	var pairs []*model.Pair
	for _, target := range targets {
		pairs = append(pairs, model.NewPair(player.User, target.User))
	}

	player.Pairs = pairs
	return player, nil
}

func (i *impl) ListPlayers(
	ctx contextx.Contextx,
	option biz.ListPlayersOption,
) (items []*agg.Player, total int, err error) {
	condition := repo.ListPlayersCondition{
		Gender:            0,
		Height:            0,
		NumsOfWantedDates: 0,
		Limit:             option.Size,
		Offset:            (option.Page - 1) * option.Size,
	}
	players, t, err := i.players.ListPlayers(ctx, condition)
	if err != nil {
		ctx.Error("list players error", zap.Error(err), zap.Any("condition", condition))
		return nil, 0, err
	}

	return players, t, nil
}

func (i *impl) SubmitPair(ctx contextx.Contextx, leftID string, rightID string) (item *model.Pair, err error) {
	left, err := i.players.GetByID(ctx, leftID)
	if err != nil {
		ctx.Error("get left player by id error", zap.Error(err))
		return nil, err
	}

	if left == nil {
		ctx.Error("left player not found", zap.String("playerID", leftID), zap.Error(errPlayerNotFound))
		return nil, errPlayerNotFound
	}

	right, err := i.players.GetByID(ctx, rightID)
	if err != nil {
		ctx.Error("get right player by id error", zap.Error(err))
		return nil, err
	}

	if right == nil {
		ctx.Error("right player not found", zap.String("playerID", rightID), zap.Error(errPlayerNotFound))
		return nil, errPlayerNotFound
	}

	pair := model.NewPair(left.User, right.User)

	left.NumsOfWantedDates--
	left.Pairs = append(left.Pairs, pair)

	right.NumsOfWantedDates--
	right.Pairs = append(right.Pairs, pair)

	err = i.players.MatchedPair(ctx, left, right, pair)
	if err != nil {
		ctx.Error("matched pair error", zap.Error(err))
		return nil, err
	}

	return pair, nil
}
