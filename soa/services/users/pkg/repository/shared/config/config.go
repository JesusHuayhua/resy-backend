package config

type ConfigData struct {
	DatabaseURL  string `json:"database_url"`
	Host         string `json:"nombre_de_host"`
	Port         string `json:"puerto"`
	User         string `json:"usuario"`
	DatabaseName string `json:"esquemabd"`
	Password     string `json:"contrasenha"` // Asumimos que está cifrada
}

// using yml for this
func LoadConfig() ConfigData {
	return ConfigData{DatabaseURL: ""}
}
