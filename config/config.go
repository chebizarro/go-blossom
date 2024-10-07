package config

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
    "time"
)

type ServerConfig struct {
    Port           int    `yaml:"port"`
    MaxUploadSize  string `yaml:"max_upload_size"`
}

type StorageConfig struct {
    Type     string `yaml:"type"`
    FilePath string `yaml:"file_path"`
}

type LoggingConfig struct {
    Level  string `yaml:"level"`
    Format string `yaml:"format"`
    Output string `yaml:"output"`
}

type SecurityConfig struct {
    AllowedMIMETypes    []string `yaml:"allowed_mime_types"`
    HashBlacklistPath   string   `yaml:"hash_blacklist_path"`
}

type AuthConfig struct {
    NostrRequired   bool   `yaml:"nostr_required"`
    ExpirationTime  string `yaml:"expiration_time"`
}

// Config holds all the configuration for the server
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Storage  StorageConfig  `yaml:"storage"`
    Logging  LoggingConfig  `yaml:"logging"`
    Security SecurityConfig `yaml:"security"`
    Auth     AuthConfig     `yaml:"auth"`
}

// LoadConfig loads configuration from the specified file path
func LoadConfig(filePath string) (*Config, error) {
    configData, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var config Config
    err = yaml.Unmarshal(configData, &config)
    if err != nil {
        return nil, err
    }

    log.Printf("Loaded config: %+v\n", config)
    return &config, nil
}

// GetExpirationDuration parses the expiration time from the config (e.g., "24h")
func (a *AuthConfig) GetExpirationDuration() (time.Duration, error) {
    return time.ParseDuration(a.ExpirationTime)
}
