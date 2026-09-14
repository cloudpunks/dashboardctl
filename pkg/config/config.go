package config

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
	Color  bool
}

// Config is a combination of all available configurations.
type Config struct {
	Logs Logs
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
