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
	player.NewPlayerRepoWithMemory,
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(providerSet, newService))
}

func NewRestful() (adapterx.Restful, error) {
	panic(wire.Build(providerSet, newRestful))
}
