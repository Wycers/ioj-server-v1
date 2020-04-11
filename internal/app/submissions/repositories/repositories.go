package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlSubmissionsRepository)

//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(SubmissionRepository),new(MockUsersRepository)))
