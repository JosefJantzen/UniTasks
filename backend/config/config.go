package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var JwtKeyBytes = []byte("SecretYouShouldHide") // Don't edit this value. It's just the fallback value. To change the key edit the config.js
var FrontendUrl = "*"                           // Don't edit this value. It's just the fallback value. To change the url edit the config.js

type Config struct {
	DB           DB
	Port         string
	JwtKey       string
	JwtExpireMin int
	Debug        bool
	FrontendUrl  string
}

func (c *Config) merge(s *Config) {
	c.DB.merge(s.DB)
	if s.Port != "" {
		c.Port = s.Port
	}
	if s.JwtKey != "" {
		c.JwtKey = s.JwtKey
		JwtKeyBytes = []byte(s.JwtKey)
	}
	if s.JwtExpireMin > c.JwtExpireMin {
		c.JwtExpireMin = s.JwtExpireMin
	}
	c.Debug = s.Debug
	if s.FrontendUrl != "" {
		c.FrontendUrl = s.FrontendUrl
		FrontendUrl = s.FrontendUrl
	}
}

type DB struct {
	User     string
	Pwd      string
	Host     string
	Port     string
	Database string
	Initial  string
	TestData string
}

func (c *DB) merge(s DB) {
	if s.User != "" {
		c.User = s.User
	}
	if s.Pwd != "" {
		c.Pwd = s.Pwd
	}
	if s.Host != "" {
		c.Host = s.Host
	}
	if s.Port != "" {
		c.Port = s.Port
	}
	if s.Database != "" {
		c.Database = s.Database
	}
	if s.Initial != "" {
		c.Initial = s.Initial
	}
	if s.TestData != "" {
		c.TestData = s.TestData
	}
}

func (c *Config) GetDBConnectString() string {
	var db = c.DB
	if c.Debug {
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			db.Host,
			db.Port,
			db.Database,
			db.User,
			db.Pwd,
		)
	}
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		db.Host,
		db.Port,
		db.Database,
		db.User,
		db.Pwd,
	)
}

func Read(file string) (*Config, error) {
	buf, err := ioutil.ReadFile("config.sample.json")
	if err != nil {
		return nil, err
	}

	var sampleConfig Config
	if err := json.Unmarshal(buf, &sampleConfig); err != nil {
		return nil, err
	}

	buf, err = ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	config.Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		config.Debug = false
	}

	sampleConfig.merge(&config)

	return &sampleConfig, err
}
