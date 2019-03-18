package config

import (
	"log"

	"gopkg.in/mgo.v2"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

func Connect() (*mgo.Database, error) {
	var config = Config{}
	config.Read()

	session, err := mgo.Dial(config.Server)
	if err != nil {
		return nil, err
	} else {
		db := session.DB(config.Database)
		return db, nil
	}
}
