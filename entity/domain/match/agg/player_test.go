package agg

import (
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		name              string
		height            uint
		gender            model.Gender
		age               uint
		numsOfWantedDates uint
	}
	tests := []struct {
		name    string
		args    args
		want    *Player
		wantErr bool
	}{
		{
			name:    "invalid name then error",
			args:    args{name: "   ", height: 100, gender: model.GenderMale, age: 20, numsOfWantedDates: 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid height then error",
			args:    args{name: "player1", height: 0, gender: model.GenderMale, age: 20, numsOfWantedDates: 5},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayer(tt.args.name, tt.args.height, tt.args.gender, tt.args.age, tt.args.numsOfWantedDates)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_WantGender(t *testing.T) {
	type fields struct {
		User              model.User
		NumsOfWantedDates uint
		Pairs             []*model.Pair
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Gender
		wantErr bool
	}{
		{
			name: "male then female",
			fields: fields{
				User: model.User{
					ID: "",
					Profile: model.Profile{
						Name:   "",
						Age:    0,
						Gender: model.GenderMale,
						Height: 0,
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				NumsOfWantedDates: 0,
				Pairs:             nil,
			},
			want:    model.GenderFemale,
			wantErr: false,
		},
		{
			name: "other then error",
			fields: fields{
				User: model.User{
					ID: "",
					Profile: model.Profile{
						Name:   "",
						Age:    0,
						Gender: model.GenderOther,
						Height: 0,
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				NumsOfWantedDates: 0,
				Pairs:             nil,
			},
			want:    model.GenderUnspecified,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Player{
				User:              tt.fields.User,
				NumsOfWantedDates: tt.fields.NumsOfWantedDates,
				Pairs:             tt.fields.Pairs,
			}
			got, err := x.WantGender()
			if (err != nil) != tt.wantErr {
				t.Errorf("WantGender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WantGender() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_WantHeight(t *testing.T) {
	type fields struct {
		User              model.User
		NumsOfWantedDates uint
		Pairs             []*model.Pair
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "other then 0",
			fields: fields{
				User: model.User{
					ID: "",
					Profile: model.Profile{
						Name:   "",
						Age:    0,
						Gender: model.GenderOther,
						Height: 180,
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				NumsOfWantedDates: 0,
				Pairs:             nil,
			},
			want: 0,
		},
		{
			name: "from male then -180",
			fields: fields{
				User: model.User{
					ID: "",
					Profile: model.Profile{
						Name:   "",
						Age:    0,
						Gender: model.GenderMale,
						Height: 180,
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				NumsOfWantedDates: 0,
				Pairs:             nil,
			},
			want: -180,
		},
		{
			name: "from female then 160",
			fields: fields{
				User: model.User{
					ID: "",
					Profile: model.Profile{
						Name:   "",
						Age:    0,
						Gender: model.GenderFemale,
						Height: 160,
					},
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				NumsOfWantedDates: 0,
				Pairs:             nil,
			},
			want: 160,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Player{
				User:              tt.fields.User,
				NumsOfWantedDates: tt.fields.NumsOfWantedDates,
				Pairs:             tt.fields.Pairs,
			}
			if got := x.WantHeight(); got != tt.want {
				t.Errorf("WantHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
