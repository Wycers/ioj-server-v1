package grpcclients

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUsersClient)
