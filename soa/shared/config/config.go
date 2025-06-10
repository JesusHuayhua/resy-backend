package config

type ConfigData struct {
	DatabaseURL string
}

// using yml for this
func LoadConfig() ConfigData {
	return ConfigData{}
}
