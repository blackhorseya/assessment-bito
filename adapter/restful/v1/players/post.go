package players

import (
	"net/http"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/model"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/response"
	"github.com/gin-gonic/gin"
)

// AddPlayerAndMatchPayload is the payload to add player and match.
type AddPlayerAndMatchPayload struct {
	model.Profile

	NumsOfWantedDates uint `json:"nums_of_wanted_dates,omitempty"`
}

// AddPlayerAndMatch is to add player and match.
// @Summary Add player and match.
// @Description add player and match.
// @Tags players
// @Accept json
// @Produce json
// @Param player body AddPlayerAndMatchPayload true "player"
// @Success 201 {object} response.Response{data=agg.Player}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/players [post]
func (i *impl) AddPlayerAndMatch(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload AddPlayerAndMatchPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
		return
	}

	player, err := i.match.EnrollPlayer(
		ctx,
		payload.Name,
		payload.Height,
		payload.Gender,
		payload.Age,
		payload.NumsOfWantedDates,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	player, err = i.match.GetPlayerByIDWithPairs(ctx, player.ID, biz.ListPairsOption{
		Page: 0,
		Size: 0,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(player))
}
