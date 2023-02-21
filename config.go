package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gopkg.in/yaml.v3"
)

type StringArray []string

type Config struct {
	LogLevel  string     `yaml:"log-level,omitempty"`
	Port      string     `yaml:"port,omitempty"`
	Junctions []Junction `yaml:"junctions"`
}

type Junction struct {
	Name    string `yaml:"name,omitempty"`
	Apprise string `yaml:"apprise"`
	To      struct {
		Emails     StringArray `yaml:"emails"`
		RequireAll bool        `yaml:"require-all,omitempty"`
	} `yaml:"to,omitempty"`
	From struct {
		Email string `yaml:"email,omitempty"`
		IP    string `yaml:"ip,omitempty"`
	} `yaml:"from,omitempty"`
	Title string `yaml:"title,omitempty"`
	Body  string `yaml:"body,omitempty"`
}

var configPath = "config/config.yaml"
var port = "8025"
var logLevel string
var apprisePath string
var junctions []Junction

func (a *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}

/*
getConf loads the configuration from the Environment Variables and Config File
*/
func getConf() {
	// Setup Logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Check for a custom path
	if path := os.Getenv("CONF_PATH"); path != "" {
		configPath = path
	}

	// Open and read the config file
	var conf Config
	file, err := os.Open(configPath)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error opening config: %s", err))
	}

	b, err := io.ReadAll(file)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error reading config: %s", err))
	}

	// Parse the yaml
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error parsing yaml: %s", err))
	}

	// Get the path to the Apprise executable
	apprisePath, err = exec.LookPath("apprise")
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error getting Apprise Path: %s", err))
	}

	// Load the parsed config
	if conf.Port != "" {
		port = conf.Port
	}
	if conf.LogLevel != "" {
		logLevel = conf.LogLevel
	}

	if logLevel != "" {
		if logLevel == "debug" {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	}

	junctions = conf.Junctions

	log.Print(fmt.Sprintf("Log Level: %s", logLevel))
}
