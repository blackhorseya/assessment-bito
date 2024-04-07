package player

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteRBTree struct {
	suite.Suite

	repo repo.IPlayerRepo
}

func (s *suiteRBTree) SetupTest() {
	s.repo = NewPlayerRepoWithRBTree()
}

func TestRBTree(t *testing.T) {
	suite.Run(t, new(suiteRBTree))
}

func (s *suiteRBTree) Test_memory_GetByID() {
	player1, err := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	s.Require().NoError(err)

	type args struct {
		ctx  contextx.Contextx
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Player
		wantErr  bool
	}{
		{
			name:     "get player then error",
			args:     args{id: player1.ID},
			wantItem: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err2 := s.repo.GetByID(tt.args.ctx, tt.args.id)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetByID() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteRBTree) Test_memory_JoinPlayer() {
	player1, err := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	s.Require().NoError(err)

	type args struct {
		ctx    contextx.Contextx
		player *agg.Player
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "add player then success",
			args:    args{player: player1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err = s.repo.JoinPlayer(tt.args.ctx, tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("JoinPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *suiteRBTree) Test_memory_LeavePlayer() {
	player1, err := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	s.Require().NoError(err)

	type args struct {
		ctx    contextx.Contextx
		player *agg.Player
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "remove player then success",
			args: args{player: player1, mock: func() {
				_ = s.repo.JoinPlayer(contextx.Background(), player1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err = s.repo.LeavePlayer(tt.args.ctx, tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("LeavePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *suiteRBTree) Test_memory_ListPlayers() {
	player1, _ := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	player2, _ := agg.NewPlayer("player2", 170, model.GenderFemale, 19, 2)
	player3, _ := agg.NewPlayer("player3", 160, model.GenderMale, 18, 1)
	player4, _ := agg.NewPlayer("player4", 150, model.GenderFemale, 17, 0)

	players := []*agg.Player{player1, player2, player3, player4}
	for _, player := range players {
		s.Require().NoError(s.repo.JoinPlayer(contextx.Background(), player))
	}

	type args struct {
		ctx       contextx.Contextx
		condition repo.ListPlayersCondition
		mock      func()
	}
	tests := []struct {
		name      string
		args      args
		wantItems []*agg.Player
		wantTotal int
		wantErr   bool
	}{
		{
			name: "ok",
			args: args{condition: repo.ListPlayersCondition{
				Gender:            model.GenderFemale,
				Height:            -175,
				NumsOfWantedDates: 1,
				Limit:             0,
				Offset:            0,
			}},
			wantItems: []*agg.Player{player2},
			wantTotal: 1,
			wantErr:   false,
		},
		{
			name: "ok",
			args: args{condition: repo.ListPlayersCondition{
				Gender:            model.GenderMale,
				Height:            175,
				NumsOfWantedDates: 1,
				Limit:             0,
				Offset:            0,
			}},
			wantItems: []*agg.Player{player1},
			wantTotal: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItems, gotTotal, err := s.repo.ListPlayers(tt.args.ctx, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItems, tt.wantItems) {
				tt.args.ctx.Error(
					"ListPlayers()",
					zap.Any("gotItems", gotItems),
					zap.Any("wantItems", tt.wantItems),
				)
				t.Errorf("ListPlayers() gotItems = %v, want %v", gotItems, tt.wantItems)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListPlayers() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
