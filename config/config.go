// Package config contains definition of configuration
// for the pinger
package config

// Config contains config for the pinger
type Config struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Address  string `yaml:"address"`
	Token    string `yaml:"token"`
}
