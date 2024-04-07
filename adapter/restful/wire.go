//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/assessment-bito/app/domain/match/biz"
	"github.com/blackhorseya/assessment-bito/app/domain/match/repo/player"
	"github.com/blackhorseya/assessment-bito/pkg/adapterx"
	"github.com/blackhorseya/assessment-bito/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	httpx.NewServer,
	biz.ProvideMatchBiz,
)

func NewMemory(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet, newService, player.NewPlayerRepoWithMemory))
}

func NewRBTree(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet, newService, player.NewPlayerRepoWithRBTree))
}
