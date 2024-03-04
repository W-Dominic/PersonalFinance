package teller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TellerAPI(t *testing.T) {
	
    tellerApi := Init("fake id", "fake token")

	assert.Equal(t, "fake id", tellerApi.appId)
	assert.Equal(t, "fake token", tellerApi.token)
}
