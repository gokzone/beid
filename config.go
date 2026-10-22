package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Gin struct {
		Mode string `json:"mode"`
	} `json:"gin"`

	Server struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"server"`

	Postgres struct {
		DSN string `json:"dsn"`
	} `json:"postgres"`

	JWT struct {
		Secret string `json:"secret"`
	} `json:"jwt"`
}

func ParseConfigFile(cfgFilePath string) (*Config, error) {
	cfgRawContent, err := os.ReadFile(cfgFilePath)
	if err != nil {
		err = fmt.Errorf("ParseConfigFile \"%s\": %w", cfgFilePath, err)
		return nil, err
	}
	cfg, err := ParseConfigContent(cfgRawContent)
	if err != nil {
		err = fmt.Errorf("ParseConfigFile \"%s\": %w", cfgFilePath, err)
		return nil, err
	}
	return cfg, nil
}

func ParseConfigContent(cfgRawContent []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(cfgRawContent, &cfg)
	if err != nil {
		err = fmt.Errorf("ParseConfigContent: %w", err)
		return nil, err
	}
	return &cfg, nil
}
