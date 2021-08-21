package main

import (
	"testing"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/stretchr/testify/assert"
)

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	state, err := checkArgs(nil)
	assert.NoError(err)
	assert.Equal(sensu.CheckStateOK, state)

}
