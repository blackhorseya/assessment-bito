package players

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/response"
	"github.com/gin-gonic/gin"
)

// ListPlayersQuery is the request for list players.
type ListPlayersQuery struct {
}

// ListPlayers is to list players.
// @Summary List players
// @Description List players
// @Tags players
// @Accept json
// @Produce json
// @Param param query ListPlayersQuery true "parameter"
// @Success 200 {object} response.Response{data=[]agg.Player}
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Header 200 {string} X-Total-Count "total player count"
// @Router /v1/players [get]
func (i *impl) ListPlayers(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var query ListPlayersQuery
	err = c.ShouldBindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
		return
	}

	players, total, err := i.match.ListPlayers(ctx, biz.ListPlayersOption{
		Page: 0,
		Size: 0,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	c.JSON(http.StatusOK, response.OK.WithData(players))
}
