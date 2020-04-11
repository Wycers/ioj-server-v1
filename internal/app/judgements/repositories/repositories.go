package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlJudgementsRepository)

//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(JudgementsRepository),new(MockJudgementsRepository)))
