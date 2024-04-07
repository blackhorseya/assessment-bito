package player

import (
	"time"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/agg"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/pkg/gods/trees/rbtree"
)

type playerDTO struct {
	ID                string       `json:"id,omitempty"`
	Name              string       `json:"name,omitempty"`
	Age               uint         `json:"age,omitempty"`
	Height            uint         `json:"height,omitempty"`
	Gender            model.Gender `json:"gender,omitempty"`
	NumsOfWantedDates uint         `json:"nums_of_wanted_dates,omitempty"`
	CreatedAt         time.Time    `json:"created_at,omitempty"`
	UpdatedAt         time.Time    `json:"updated_at,omitempty"`
}

func newPlayerDTO(v *agg.Player) *playerDTO {
	return &playerDTO{
		ID:                v.ID,
		Name:              v.Profile.Name,
		Age:               v.Profile.Age,
		Height:            v.Profile.Height,
		Gender:            v.Profile.Gender,
		NumsOfWantedDates: v.NumsOfWantedDates,
		CreatedAt:         v.CreatedAt,
		UpdatedAt:         v.UpdatedAt,
	}
}

func (x *playerDTO) Less(than rbtree.Item) bool {
	return x.Height <= than.(*playerDTO).Height
}

// ToAgg is to convert playerDTO to agg.Player.
func (x *playerDTO) ToAgg() *agg.Player {
	return &agg.Player{
		User: model.User{
			ID: x.ID,
			Profile: model.Profile{
				Name:   x.Name,
				Age:    x.Age,
				Gender: x.Gender,
				Height: x.Height,
			},
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
		},
		NumsOfWantedDates: x.NumsOfWantedDates,
		Pairs:             nil,
		PotentialPairs:    nil,
		MatchedPairs:      nil,
	}
}

type pairDTO struct {
	ID        string    `json:"id,omitempty"`
	LeftID    string    `json:"left_id,omitempty"`
	RightID   string    `json:"right_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func newPairDTO(v *model.Pair) *pairDTO {
	return &pairDTO{
		ID:        v.ID,
		LeftID:    v.Left.ID,
		RightID:   v.Right.ID,
		CreatedAt: v.CreatedAt,
	}
}
