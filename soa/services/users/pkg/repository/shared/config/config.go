package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

type ConfigData struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

func LoadConfig() string {
	relPath := filepath.Join("soa", "services", "users", "pkg", "repository", "shared", "config", "config.json")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("config: no pude obtener working dir: %v", err)
	}
	var path string
	for dir := wd; ; dir = filepath.Dir(dir) {
		try := filepath.Join(dir, relPath)
		if _, err := os.Stat(try); err == nil {
			path = try
			break
		}
		// Llegó a la raíz y no encontró nada.
		if dir == filepath.Dir(dir) {
			log.Fatalf("config: no se encontró %s desde %s", relPath, wd)
		}
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("config: no pude abrir %s: %v", path, err)
	}
	defer f.Close()
	var rc ConfigData
	if err := json.NewDecoder(f).Decode(&rc); err != nil {

	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		rc.User,
		url.QueryEscape(rc.Password),
		rc.Host,
		rc.Port,
		rc.DBName,
		rc.SSLMode,
	)
	return dsn
}
