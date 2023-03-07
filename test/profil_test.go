package test

import (
	"testing"

	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/api"
	"github.com/BeanCodeDe/SpaceLight-Data/test/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProfile(t *testing.T) {
	util.CreateProfile_Automated(t)
}

func TestCreateProfile_Retry(t *testing.T) {
	userId, randomUserName, token := util.CreateProfile_Automated(t)

	profileCreationDto := &api.ProfileCreateDTO{Name: randomUserName}

	statusRetry := util.CreateProfile(userId, profileCreationDto, token)
	assert.Equal(t, 201, statusRetry)
}

func TestCreateProfile_Retry_WrongUser(t *testing.T) {
	randomUserName := util.RandomString(8)

	userId, _, token := util.CreateProfile_Automated(t)

	profileCreationDto := &api.ProfileCreateDTO{Name: randomUserName}

	statusRetry := util.CreateProfile(userId, profileCreationDto, token)
	assert.Equal(t, 500, statusRetry)
}

func TestCreateProfile_UserExists(t *testing.T) {
	userId := util.CreateUserId()
	token := util.CreateJWTToken(userId)

	_, randomUserName, _ := util.CreateProfile_Automated(t)

	profileCreationDto := &api.ProfileCreateDTO{Name: randomUserName}

	statusRetry := util.CreateProfile(userId, profileCreationDto, token)
	assert.Equal(t, 201, statusRetry)
}

func TestGetProfile(t *testing.T) {
	userId, randomUserName, token := util.CreateProfile_Automated(t)
	status, profile := util.GetProfile(userId, token)

	assert.Equal(t, 200, status)
	assert.Equal(t, userId, profile.UserId.String())
	assert.Equal(t, randomUserName, profile.Name)
}

func TestGetProfile_WrongUserId(t *testing.T) {
	_, _, token := util.CreateProfile_Automated(t)
	status, _ := util.GetProfile(uuid.NewString(), token)

	assert.Equal(t, 401, status)
}

func TestGetProfile_WrongToken(t *testing.T) {
	_, _, token := util.CreateProfile_Automated(t)
	userId, _, _ := util.CreateProfile_Automated(t)
	status, _ := util.GetProfile(userId, token)

	assert.Equal(t, 401, status)
}
