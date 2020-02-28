package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {

	uuid, err := GenerateUUID()
	if err != nil {
		log.Fatal(err.Error())
	}

	assert.NotEmpty(t, uuid)
	assert.Equal(t, 36, len(uuid))
}
