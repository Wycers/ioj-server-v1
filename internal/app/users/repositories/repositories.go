package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMysqlUsersRepository)

//var MockProviderSet = wire.NewSet(wire.InterfaceValue(new(UsersRepository),new(MockUsersRepository)))
