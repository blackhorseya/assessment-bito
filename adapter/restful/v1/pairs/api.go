package pairs

import (
	"github.com/blackhorseya/assessment-bito/entity/domain/match/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	match biz.IMatchBiz
}

// Handle is to handle pair api.
func Handle(g *gin.RouterGroup, match biz.IMatchBiz) {
	instance := &impl{
		match: match,
	}

	g.POST("", instance.CreatePair)
}
