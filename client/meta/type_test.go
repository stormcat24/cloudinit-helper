package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRegion(t *testing.T) {

	az := AvailabilityZone{
		Name: "ap-northeast-1a",
	}

	assert.Equal(t, "ap-northeast-1", az.GetRegion())
}