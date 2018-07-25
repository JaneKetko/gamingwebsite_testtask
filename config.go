package main

import (
	"io/ioutil"
	"log"

	"fmt"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
)

// settings is configure settings
type settings struct {
	// Address is Mongo address.
	Address string `yaml:"dbaddress" short:"a" long:"dbaddress" description:"Mongo db address" required:"true" default:"127.0.0.1:27017"`
	// DBName is name of MongoDB.
	DBName string `yaml:"dbname" long:"dbname" description:"Mongo db name" required:"true" default:"GamingDB"`
	// PlayerCollection is name of players collection.
	PlayerCollection string `yaml:"players" short:"p" long:"players" description:"Player collection" required:"true" default:"players"`
	// TournamentCollection is name of tournament collection.
	TournamentCollection string `yaml:"tournaments" short:"t" long:"tournaments" description:"Tournament collection" required:"true" default:"tournaments"`
	// ServerAddress is address of server.
	ServerAddress string `yaml:"server" short:"s" long:"server" description:"Server address" required:"true" default:":8080"`
	// ConfigFile is file with configs.
	ConfigFile string `short:"f" long:"configfile" description:"File with config"`
}

// Parse parses command line parameters. If there is ConfigFile, then override params by values from file.
func (s *settings) Parse() error {
	parser := flags.NewParser(s, flags.Default|flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}
	if s.ConfigFile != "" {
		err = s.LoadOptionsFromFile()
		if err != nil {
			log.Printf("cannot read settings from file: %v", err)
		}
	}
	return nil
}

// LoadOptionsFromFile tries to read settings from file.
func (s *settings) LoadOptionsFromFile() error {
	data, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, s)
}
