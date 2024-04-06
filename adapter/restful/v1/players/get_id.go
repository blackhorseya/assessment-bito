package players

import (
	"net/http"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// QuerySinglePeopleQuery is the request for query single people.
type QuerySinglePeopleQuery struct {
	N int `form:"n" default:"10" minimum:"1"  maximum:"100"`
}

// QuerySinglePeople is to query single people.
// @Summary Query single people
// @Description Query single people
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "player id"
// @Param param query QuerySinglePeopleQuery true "parameter"
// @Success 200 {object} response.Response{data=agg.Player}
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/players/{id} [get]
func (i *impl) QuerySinglePeople(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
		return
	}

	var query QuerySinglePeopleQuery
	err = c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
		return
	}

	player, err := i.match.GetPlayerByIDWithPairs(ctx, id.String(), biz.ListPairsOption{
		Page: 1,
		Size: query.N,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(player))
}
