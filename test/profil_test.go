package test

import (
	"testing"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/api"
	"github.com/BeanCodeDe/SpaceLight-Data/test/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	userId := util.CreateUserId()
	token := util.CreateJWTToken(userId)
	profileCreationDto := &api.ProfileCreateDTO{Name: "Bean900"}
	status := util.CreateProfile(userId, profileCreationDto, token)
	assert.Equal(t, 201, status)
}
