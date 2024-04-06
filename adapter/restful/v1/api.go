package v1

import (
	"github.com/blackhorseya/assessment-bito/adapter/restful/v1/pairs"
	"github.com/blackhorseya/assessment-bito/adapter/restful/v1/players"
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/gin-gonic/gin"
)

// Handle is to handle the api.
func Handle(g *gin.RouterGroup, match biz.IMatchBiz) {
	pairs.Handle(g.Group("/pairs"), match)
	players.Handle(g.Group("/players"), match)
}
