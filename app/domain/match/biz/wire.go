package biz

import (
	"github.com/google/wire"
)

// ProvideMatchBiz is to provide match biz.
var ProvideMatchBiz = wire.NewSet(NewMatchBiz)
