package test

import (
	"testing"

	"github.com/BeanCodeDe/SpaceLight-Data/test/util"
	"github.com/stretchr/testify/assert"
)

func TestGetShips(t *testing.T) {
	_, _, token := util.CreateProfile_Automated(t)
	status, _ := util.GetShips(token)
	assert.Equal(t, 200, status)
}
