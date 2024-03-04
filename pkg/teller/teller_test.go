package teller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TellerAPI(t *testing.T) {

	tellerApi := Init("fake id", "fake token", "https://api.teller.io", "/path/to/cert", "/path/to/key")

	assert.Equal(t, "fake id", tellerApi.appId)
	assert.Equal(t, "fake token", tellerApi.token)
	assert.Equal(t, "https://api.teller.io", tellerApi.baseUrl)
	assert.Equal(t, "/path/to/cert", tellerApi.certPath)
	assert.Equal(t, "/path/to/key", tellerApi.keyPath)
}
