package stns

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/STNS/STNS/model"
	"github.com/go-yaml/yaml"
)

type decoder interface {
	decode(string, *Config) error
}

type tomlDecoder struct{}
type yamlDecoder struct{}

func (t *tomlDecoder) decode(path string, conf *Config) error {
	_, err := toml.DecodeFile(path, conf)
	return err
}

func (y *yamlDecoder) decode(path string, conf *Config) error {
	_, err := toml.DecodeFile(path, conf)

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, conf)
	return err
}

func decode(path string, conf *Config) error {
	var d decoder
	d = new(tomlDecoder)
	if strings.HasSuffix(path, "yaml") || strings.HasSuffix(path, "yml") {
		d = new(yamlDecoder)
	}
	return d.decode(path, conf)
}

func NewConfig(confPath string) (Config, error) {
	var conf Config
	defaultConfig(&conf)

	if err := decode(confPath, &conf); err != nil {
		return conf, err
	}

	if conf.Include != "" {
		if err := includeConfigFile(&conf, conf.Include); err != nil {
			return Config{}, err
		}
	}

	return conf, nil
}

type Config struct {
	Port      int        `toml:"port"`
	BasicAuth *BasicAuth `toml:"basic_auth" yaml:"basic_auth"`
	TokenAuth *TokenAuth `toml:"token_auth" yaml:"token_auth"`

	UseServerStarter bool
	Users            *model.Users
	Groups           *model.Groups
	Include          string `toml:"include"`
}

type BasicAuth struct {
	User     string
	Password string
}
type TokenAuth struct {
	Tokens []string
}

func defaultConfig(c *Config) {
	c.Port = 1104
}

func includeConfigFile(config *Config, include string) error {
	files, err := filepath.Glob(include)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := decode(file, config); err != nil {
			return fmt.Errorf("while loading included config file %s: %s", file, err)
		}
	}
	return nil
}
