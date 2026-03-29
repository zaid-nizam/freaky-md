package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BotName     string   `yaml:"bot_name"`
	Prefixes    []string `yaml:"prefixes"`
	Owners      []string `yaml:"owners"`       // Format: "1234567890@s.whatsapp.net"
	SessionPath string   `yaml:"session_path"` // Where the .db file lives
	AutoRead    bool     `yaml:"auto_read"`    // Mark messages as read automatically
	Debug       bool     `yaml:"debug"`        // Enable whatsmeow log spam
}

func Default() *Config {
	return &Config{
		BotName:     "FreakyMD",
		Prefixes:    []string{"!", "."},
		Owners:      []string{},
		SessionPath: "./freaky.db",
		AutoRead:    false,
		Debug:       false,
	}
}

func (c *Config) IsOwner(jid string) bool {
	cleanJid := strings.Split(jid, ":")[0]
	if !strings.Contains(cleanJid, "@") {
		cleanJid += "@s.whatsapp.net"
	}

	for _, owner := range c.Owners {
		if owner == cleanJid {
			return true
		}
	}
	return false
}

func Load(path string) *Config {
	conf := Default()
	data, err := os.ReadFile(path)
	if err != nil {
		return conf
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		panic("Invalid config format: " + err.Error())
	}
	return conf
}
