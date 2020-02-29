package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDGenerator(t *testing.T) {

	uuid, err := GenerateUUID()
	if err != nil {
		log.Fatal(err.Error())
	}

	assert.NotEmpty(t, uuid)
	assert.Equal(t, 36, len(uuid))
}

func TestPriceFormater(t *testing.T) {

	got := FormatPrice(95.0)
	want := "95.00€"
	assert.Equal(t, got, want)

	got = FormatPrice(95.05)
	want = "95.05€"
	assert.Equal(t, got, want)
}
