package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	DB   DB
	Port string
}

type DB struct {
	User     string
	Pwd      string
	Host     string
	Port     string
	Database string
	Initial  string
}

func (db *DB) GetDBConnectString() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		db.Host,
		db.Port,
		db.Database,
		db.User,
		db.Pwd,
	)
}

func Read(file string) (*Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, err
}
