package biz

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/repo"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var (
	errMock = errors.New("mock error")
)

type suiteTester struct {
	suite.Suite

	ctrl    *gomock.Controller
	players *repo.MockIPlayerRepo
	biz     biz.IMatchBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.players = repo.NewMockIPlayerRepo(s.ctrl)
	s.biz = NewMatchBiz(s.players)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_EnrollPlayer() {
	player1, err := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	s.Require().NoError(err)

	type args struct {
		ctx               contextx.Contextx
		name              string
		height            uint
		gender            model.Gender
		age               uint
		numsOfWantedDates uint
		mock              func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Player
		wantErr  bool
	}{
		{
			name:     "invalid player infor then error",
			args:     args{name: "   ", height: 0},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "join player then error",
			args: args{
				name:              player1.Profile.Name,
				height:            player1.Profile.Height,
				gender:            player1.Profile.Gender,
				age:               player1.Profile.Age,
				numsOfWantedDates: player1.NumsOfWantedDates,
				mock: func() {
					s.players.EXPECT().JoinPlayer(gomock.Any(), player1).Return(errMock).Times(1)
				},
			},
			wantItem: nil,
			wantErr:  true,
		},
		{
			name: "join player then success",
			args: args{
				name:              player1.Profile.Name,
				height:            player1.Profile.Height,
				gender:            player1.Profile.Gender,
				age:               player1.Profile.Age,
				numsOfWantedDates: player1.NumsOfWantedDates,
				mock: func() {
					s.players.EXPECT().JoinPlayer(gomock.Any(), player1).Return(nil).Times(1)
				},
			},
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
				tt.args.name,
				tt.args.height,
				tt.args.gender,
				tt.args.age,
				tt.args.numsOfWantedDates,
			)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("EnrollPlayer() error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("EnrollPlayer() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteTester) Test_impl_UnregisterPlayer() {
	player1, err := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	s.Require().NoError(err)

	type args struct {
		ctx      contextx.Contextx
		playerID string
		mock     func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get player by id then error",
			args: args{playerID: "not found", mock: func() {
				s.players.EXPECT().GetByID(gomock.Any(), "not found").Return(nil, errMock).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "not found player then error",
			args: args{playerID: "not found", mock: func() {
				s.players.EXPECT().GetByID(gomock.Any(), "not found").Return(nil, nil).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "got player to unregister then error",
			args: args{playerID: player1.ID, mock: func() {
				s.players.EXPECT().GetByID(gomock.Any(), player1.ID).Return(player1, nil).Times(1)

				s.players.EXPECT().LeavePlayer(gomock.Any(), player1).Return(errMock).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "got player to unregister then ok",
			args: args{playerID: player1.ID, mock: func() {
				s.players.EXPECT().GetByID(gomock.Any(), player1.ID).Return(player1, nil).Times(1)

				s.players.EXPECT().LeavePlayer(gomock.Any(), player1).Return(nil).Times(1)
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

			if err = s.biz.UnregisterPlayer(tt.args.ctx, tt.args.playerID); (err != nil) != tt.wantErr {
				t.Errorf("UnregisterPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *suiteTester) Test_impl_GetPlayerByIDWithPairs() {
	player1, _ := agg.NewPlayer("player1", 180, model.GenderMale, 20, 3)
	player2, _ := agg.NewPlayer("player2", 170, model.GenderFemale, 20, 3)

	pair := model.NewPair(player1.User, player2.User)
	player1.Pairs = []*model.Pair{pair}

	type args struct {
		ctx      contextx.Contextx
		playerID string
		option   biz.ListPairsOption
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *agg.Player
		wantErr  bool
	}{
		{
			name: "ok",
			args: args{playerID: player1.ID, mock: func() {
				s.players.EXPECT().GetByID(gomock.Any(), player1.ID).Return(player1, nil).Times(1)

				wantGender, _ := player1.WantGender()
				s.players.EXPECT().ListPlayers(gomock.Any(), repo.ListPlayersCondition{
					Gender:            wantGender,
					Height:            player1.WantHeight(),
					NumsOfWantedDates: 1,
					Limit:             0,
					Offset:            0,
				}).Return([]*agg.Player{player2}, 1, nil).Times(1)
			}},
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

			gotItem, err := s.biz.GetPlayerByIDWithPairs(tt.args.ctx, tt.args.playerID, tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerByIDWithPairs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetPlayerByIDWithPairs() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func (s *suiteTester) Test_impl_SubmitPair() {
	left1, _ := agg.NewPlayer("left1", 180, model.GenderMale, 20, 3)
	right1, _ := agg.NewPlayer("right1", 170, model.GenderFemale, 20, 3)
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
			name: "matched pair then error",
			args: args{
				leftID:  left1.ID,
				rightID: right1.ID,
				mock: func() {
					s.players.EXPECT().GetByID(gomock.Any(), left1.ID).Return(left1, nil).Times(1)
					s.players.EXPECT().GetByID(gomock.Any(), right1.ID).Return(right1, nil).Times(1)

					s.players.EXPECT().MatchedPair(gomock.Any(), left1, right1, pair1).Return(errMock).Times(1)
				},
			},
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

			gotItem, err := s.biz.SubmitPair(tt.args.ctx, tt.args.leftID, tt.args.rightID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubmitPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("SubmitPair() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
