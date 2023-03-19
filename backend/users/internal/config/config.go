package config

// Config defines the information needed to configure users service.
type Config struct {
	Address string
}

// Addr returns string representing configured address.
func (c Config) Addr() string {
	return c.Address
}