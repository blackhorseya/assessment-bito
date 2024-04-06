package players

import (
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	match biz.IMatchBiz
}

// Handle is to handle the matches api.
func Handle(g *gin.RouterGroup, match biz.IMatchBiz) {
	instance := &impl{
		match: match,
	}

	g.GET("/:id", instance.QuerySinglePeople)
	g.GET("", instance.ListPlayers)
	g.POST("", instance.AddPlayerAndMatch)
	g.DELETE("/:id", instance.RemoveSinglePerson)
}
