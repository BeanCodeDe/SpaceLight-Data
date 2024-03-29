package test

import (
	"testing"

	"github.com/BeanCodeDe/SpaceLight-Data/test/util"
	"github.com/stretchr/testify/assert"
)

func TestGetShips(t *testing.T) {
	_, _, token := util.CreateProfile_Automated(t)
	status, ships := util.GetShips(token)
	assert.Equal(t, 200, status)
	assert.Equal(t, 1, len(*ships))
}

func TestGetShips_WithoutToken(t *testing.T) {
	status, _ := util.GetShips("")
	assert.Equal(t, 401, status)
}
