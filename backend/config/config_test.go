package config

import (
	"fmt"
	"testing"
)

var c = Config{
	DB: DB{
		User:     "abc",
		Pwd:      "1234545",
		Host:     "localhost",
		Port:     "12345",
		Database: "unitasks",
		Initial:  "init.sql",
		TestData: "test.sql",
	},
	Port:         "12346",
	JwtKey:       "jkasd98hreniiad",
	JwtExpireMin: 5,
	Debug:        false,
	FrontendUrl:  "https://test.unitasks.de",
}

func TestMerge(t *testing.T) {
	c1 := Config{
		DB: DB{
			User:    "abc",
			Host:    "localhost",
			Port:    "12345",
			Initial: "init.sql",
		},
		Port:         "12",
		JwtKey:       "jkasd98hreniiad",
		JwtExpireMin: 5,
		Debug:        true,
	}
	c2 := Config{
		DB: DB{
			Pwd:      "1234545",
			Host:     "localhost",
			Database: "unitasks",
			Initial:  "init.sql",
			TestData: "test.sql",
		},
		Port:         "12346",
		JwtExpireMin: 5,
		FrontendUrl:  "https://test.unitasks.de",
	}

	c1.merge(&c2)
	if c1 != c {
		fmt.Println(c1)
		t.Errorf("Config merge not successfull")
	}
}

func TestGetDBConnectString(t *testing.T) {
	if c.GetDBConnectString() != "host=localhost port=12345 dbname=unitasks user=abc password=1234545" {
		t.Errorf("GetDBConnectString error with debug=false")
	}
	c.Debug = true
	if c.GetDBConnectString() != "host=localhost port=12345 dbname=unitasks user=abc password=1234545 sslmode=disable" {
		t.Errorf("GetDBConnectString error with debug=true")
	}
	c.Debug = false
}

/* Doesn't work because of file location of the sample config
func TestRead(t *testing.T) {
	c, err := Read("../config.sample.json")
	if err != nil {
		t.Errorf("Failed to execute Read function")
		return
	}
	buf, err := ioutil.ReadFile("../config.sample.json")
	if err != nil {
		t.Errorf("Failed to read config.sample.json")
		return
	}

	var sampleConfig Config
	if err := json.Unmarshal(buf, &sampleConfig); err != nil {
		t.Errorf("Failed to parse config.sample.json to object")
		return
	}

	if c != &sampleConfig {
		t.Errorf("Both reads are not equivalent")
	}
}
*/
