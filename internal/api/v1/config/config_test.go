package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * TestReadConfig reads the config file thats not exist and return error
 */
func TestReadConfigNotExist(t *testing.T) {
	// read config file
	config, err := ReadConfig("path/to/file.yaml")
	assert.Nil(t, config)
	assert.Error(t, err)
}

/**
 * TestReadConfig reads the config file thats exist but have bad structure and return error
 */
func TestReadConfigExistBadStructure(t *testing.T) {
	// read config file
	config, err := ReadConfig("config.yaml")
	assert.Nil(t, config)
	assert.Error(t, err)
}

/**
 * TestReadConfig reads the config file thats exist and have good structure and return config
 */
func TestReadConfigExistGoodStructure(t *testing.T) {
	// read config file
	config, err := ReadConfig("../../../../config.yaml")
	assert.NotNil(t, config)
	assert.NoError(t, err)
}
