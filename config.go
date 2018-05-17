package main

import (
	"io/ioutil"
	"log"

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
		// TODO you should wrap your error. This error won't readable in the future.
		return err
	}
	if s.ConfigFile != "" {
		err = s.LoadOptionsFromFile()
		if err != nil {
			// TODO it is better add space between your error and LoadOptionsFromFile's error.
			// TODO why don't you return this error like return err with this formatting?
			log.Printf("cannot read settings from file:%v", err)
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
