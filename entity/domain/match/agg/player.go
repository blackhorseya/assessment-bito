package agg

import (
	"errors"
	"strings"
	"time"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
)

var (
	errGenderNotSupport = errors.New("gender not support")
)

// Player is an entity that represents a player.
type Player struct {
	model.User `json:",inline"`

	NumsOfWantedDates uint          `json:"nums_of_wanted_dates"`
	Pairs             []*model.Pair `json:"pairs,omitempty"`

	PotentialPairs []*model.Pair `json:"-"`
	MatchedPairs   []*model.Pair `json:"-"`
}

// NewPlayer is to create a new player.
func NewPlayer(name string, height uint, gender model.Gender, age uint, numsOfWantedDates uint) (*Player, error) {
	if strings.ReplaceAll(name, " ", "") == "" {
		return nil, errors.New("[name] is required")
	}

	if height == 0 {
		return nil, errors.New("[height] MUST be greater than 0")
	}

	return &Player{
		User: model.User{
			ID: "",
			Profile: model.Profile{
				Name:   name,
				Age:    age,
				Gender: gender,
				Height: height,
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		NumsOfWantedDates: numsOfWantedDates,
		Pairs:             nil,
	}, nil
}

// WantGender is to get the player to want target gender.
func (x *Player) WantGender() (model.Gender, error) {
	switch x.Profile.Gender {
	case model.GenderUnspecified:
		return model.GenderUnspecified, errGenderNotSupport
	case model.GenderMale:
		return model.GenderFemale, nil
	case model.GenderFemale:
		return model.GenderMale, nil
	case model.GenderOther:
		return model.GenderUnspecified, nil
	default:
		return model.GenderUnspecified, errGenderNotSupport
	}
}

// WantHeight is to get the player to want target height.
func (x *Player) WantHeight() int {
	wantGender, err := x.WantGender()
	if err != nil {
		return 0
	}

	if wantGender == model.GenderFemale {
		return -int(x.Profile.Height)
	}

	if wantGender == model.GenderMale {
		return int(x.Profile.Height)
	}

	return 0
}
