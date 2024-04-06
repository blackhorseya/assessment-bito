package pairs

import (
	"net/http"

	"github.com/blackhorseya/assessment-bito/pkg/contextx"
	"github.com/blackhorseya/assessment-bito/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreatePairPayload is the request for create pair.
type CreatePairPayload struct {
	LeftID  string `json:"left_id"`
	RightID string `json:"right_id"`
}

// CreatePair is to create pair.
// @Summary Create pair
// @Description Create pair
// @Tags pairs
// @Accept json
// @Produce json
// @Param pair body CreatePairPayload true "pair"
// @Success 201 {object} response.Response{data=model.Pair}
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/pairs [post]
func (i *impl) CreatePair(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload CreatePairPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err.WithMessage(err.Error()))
		return
	}

	ret, err := i.match.SubmitPair(ctx, payload.LeftID, payload.RightID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err.WithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}
