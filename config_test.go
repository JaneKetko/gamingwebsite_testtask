package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_settings_LoadOptionsFromFile(t *testing.T) {
	tt := []struct {
		name             string
		filename         string
		expectError      bool
		expectedSettings settings
	}{
		{
			name:     "Success",
			filename: "testdata/settings.yaml",
			expectedSettings: settings{
				Address:          "127.0.0.1:27017",
				DBName:           "database",
				PlayerCollection: "players",
				ServerAddress:    ":8080",
				ConfigFile:       "testdata/settings.yaml",
			},
			expectError: false,
		},
		{
			name:        "FileError",
			filename:    "testdata/filenotexist.yaml",
			expectError: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := settings{ConfigFile: tc.filename}
			err := s.LoadOptionsFromFile()
			if tc.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedSettings, s)
			}
		})
	}
}

func TestSettings_Parse(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tt := []struct {
		name             string
		args             []string
		expDBName        string
		expAddress       string
		expPlayers       string
		expServerAddress string
		expConfig        string
	}{
		{
			name:             "success",
			args:             []string{"cmd", "--dbname=db", "-a=127.0.0.1", "--players=coll", "--server=:1234"},
			expDBName:        "db",
			expAddress:       "127.0.0.1",
			expPlayers:       "coll",
			expServerAddress: ":1234",
			expConfig:        "",
		},
		{
			name:             "SuccessPartialSettings",
			args:             []string{"cmd", "--dbname=db", "-a=127.0.0.1"},
			expDBName:        "db",
			expAddress:       "127.0.0.1",
			expPlayers:       "players",
			expServerAddress: ":8080",
			expConfig:        "",
		},
		{
			name:             "SuccessLostConfigFile",
			args:             []string{"cmd", "--dbname=db", "-a=127.0.0.1", "--players=coll", "--server=:1234", "--configfile=testdata/filenotexist"},
			expDBName:        "db",
			expAddress:       "127.0.0.1",
			expPlayers:       "coll",
			expServerAddress: ":1234",
			expConfig:        "testdata/filenotexist",
		},
		{
			name:             "SuccessConfigFile",
			args:             []string{"cmd", "--dbname=testDB", "-a=127.0.0.1", "--players=coll", "--server=:1234", "--configfile=testdata/config.yaml"},
			expDBName:        "testDB",
			expAddress:       "10.10.10.10:27017",
			expPlayers:       "players",
			expServerAddress: ":1234",
			expConfig:        "testdata/config.yaml",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := settings{}
			os.Args = tc.args
			err := s.Parse()
			require.NoError(t, err)
			assert.Equal(t, tc.expDBName, s.DBName)
			assert.Equal(t, tc.expAddress, s.Address)
			assert.Equal(t, tc.expPlayers, s.PlayerCollection)
			assert.Equal(t, tc.expServerAddress, s.ServerAddress)
			assert.Equal(t, tc.expConfig, s.ConfigFile)
		})
	}
}
