package config

// Config defines the information needed to configure notebooks service.
type Config struct {
	Address string
}

// Addr returns string representing the configured address.
func (c Config) Addr() string {
	return c.Address
}