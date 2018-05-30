package main

import (
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

// TODO: do we need test for Parse()? If need, how to set CLI flags?
