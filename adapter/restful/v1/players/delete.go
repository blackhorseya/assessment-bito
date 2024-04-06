package players

import (
	"net/http"

	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RemoveSinglePerson is to remove single person.
// @Summary Remove single person
// @Description Remove single person
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "player id"
// @Success 204
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/players/{id} [delete]
func (i *impl) RemoveSinglePerson(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
	}

	err = i.match.UnregisterPlayer(ctx, id.String())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
