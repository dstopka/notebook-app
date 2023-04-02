package config

import (
	"github.com/dstopka/notebook-app/backend/common/server"
)

var _ server.GRPCServerConfig = (*Config)(nil)

// Config defines the information needed to configure users service.
type Config struct {
	// Address is the address on which the grpc server will be run on.
	Address string
}

// Addr returns string representing configured address.
func (c Config) Addr() string {
	return c.Address
}
