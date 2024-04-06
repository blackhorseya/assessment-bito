//go:build integration

package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/assessment-bito/app/domain/match/repo/player"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/stretchr/testify/suite"
)

type suiteMemory struct {
	suite.Suite

	players repo.IPlayerRepo
	biz     biz.IMatchBiz
}

func (s *suiteMemory) SetupTest() {
	s.players = player.NewPlayerRepoWithMemory()
	s.biz = NewMatchBiz(s.players)
}

func TestMemory(t *testing.T) {
	suite.Run(t, new(suiteMemory))
}

func (s *suiteMemory) Test_impl_EnrollPlayer() {
	player1, _ := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)

	type args struct {
		ctx    contextx.Contextx
		player *agg.Player
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Player
		wantErr  bool
	}{
		{
			name:     "ok",
			args:     args{player: player1},
			wantItem: player1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err2 := s.biz.EnrollPlayer(
				tt.args.ctx,
				tt.args.player.Profile.Name,
				tt.args.player.Profile.Height,
				tt.args.player.Profile.Gender,
				tt.args.player.Profile.Age,
				tt.args.player.NumsOfWantedDates,
			)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("EnrollPlayer() error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
			gotItem.ID = tt.wantItem.ID
			gotItem.CreatedAt = tt.wantItem.CreatedAt
			gotItem.UpdatedAt = tt.wantItem.UpdatedAt
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("EnrollPlayer() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteMemory) Test_impl_SubmitPair() {
	left1, _ := agg.NewPlayer("left1", 180, model.GenderMale, 20, 3)
	right1, _ := agg.NewPlayer("right1", 170, model.GenderFemale, 20, 3)

	players := []*agg.Player{left1, right1}
	for _, p := range players {
		_ = s.players.JoinPlayer(contextx.Background(), p)
	}

	pair1 := model.NewPair(left1.User, right1.User)

	type args struct {
		ctx     contextx.Contextx
		leftID  string
		rightID string
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.Pair
		wantErr  bool
	}{
		{
			name:     "matched pair then ok",
			args:     args{leftID: left1.ID, rightID: right1.ID},
			wantItem: pair1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.biz.SubmitPair(tt.args.ctx, tt.args.leftID, tt.args.rightID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubmitPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotItem.ID = tt.wantItem.ID
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("SubmitPair() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
